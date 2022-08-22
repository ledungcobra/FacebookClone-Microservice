package users

import (
	"fmt"
	"ledungcobra/gateway-go/pkg/controllers/users/response"
	"ledungcobra/gateway-go/pkg/dao"
	"ledungcobra/gateway-go/pkg/htmltemplates"
	"ledungcobra/gateway-go/pkg/middlewares"
	"ledungcobra/gateway-go/pkg/models"
	"ledungcobra/gateway-go/pkg/service"
	"log"
	"os"
	"time"

	"ledungcobra/gateway-go/pkg/common"
	"ledungcobra/gateway-go/pkg/controllers/base"
	"ledungcobra/gateway-go/pkg/controllers/users/request"
	"ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/validators"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userDao interfaces.IUserDAO
	base.BaseController
	notificationService interfaces.INotificationService
}

func NewUserController(userDao interfaces.IUserDAO,
	notificationService interfaces.INotificationService) *UserController {
	this := &UserController{
		userDao:             userDao,
		BaseController:      base.BaseController{},
		notificationService: notificationService,
	}
	return this
}

func (u *UserController) RegisterUserRouter(apiRouter fiber.Router) {
	userRouter := apiRouter.Group("/users")
	userRouter.Post("/register", u.Register)
	userRouter.Post("/login", u.Login)
	userRouter.Post("/activate", u.ActiveAccount)
	userRouter.Get("/auth", middlewares.Protected, u.AuthTest)
	userRouter.Post("/sendVerification", middlewares.Protected, u.ResendVerification)

}

func (u *UserController) Register(ctx *fiber.Ctx) error {
	var err error
	var registerRequest request.RegisterRequest
	if err := ctx.BodyParser(&registerRequest); err != nil {
		log.Println("Error when binding data", err)
		return u.InvalidFormResponse(ctx, err)
	}

	if isValid, errors := validators.Validate(&registerRequest); !isValid {
		return u.SendBadRequestWithError(ctx, "Check invalid form before proceed", errors)
	}

	user := mapRegisterRequestToUser(registerRequest)
	if user.Password, err = common.HashPassword(registerRequest.Password); err != nil {
		log.Println("Hash password error ", err)
		return u.SendServerError(ctx, err)
	}
	userName := registerRequest.FirstName + registerRequest.LastName
	if user.UserName, err = common.GenerateUniqueUserName(u.userDao, userName); err != nil {
		log.Println("Error when generating username ", err)
		return u.SendServerError(ctx, err)
	}

	if err := u.userDao.Save(&user); err != nil {
		if _, e := u.userDao.Find("email=?", user.Email); e == nil {
			return u.SendBadRequestWithError(ctx, "Email is already exist", err)
		}
		return u.SendServerError(ctx, err)
	}

	emailVerificationToken, err := u.sendVerification(ctx, user)
	if err != nil {
		return u.SendServerError(ctx, err)
	}
	return u.SendCreated(ctx, response.RegisterResponse{
		Success:   true,
		UserName:  user.UserName,
		ID:        user.ID,
		Picture:   user.Picture,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Verified:  user.Verified,
		Token:     emailVerificationToken,
	}, "Register user success please active your email to start")
}

func (u *UserController) sendVerification(ctx *fiber.Ctx, user models.User) (string, error) {
	oneMonth := time.Hour * 24 * 30
	emailVerificationToken, err := u.generateToken(ctx, common.JSON{"email": user.Email}, oneMonth)
	user.VerificationToken = emailVerificationToken
	if err := u.userDao.Save(&user); err != nil {
		return "", u.SendServerError(ctx, err)
	}

	if err != nil {
		return "", err
	}
	emailResponse, err := u.notificationService.SendMail(user.Email, "Verification Email",
		htmltemplates.BuildRegistrationTemplate(user.UserName,
			fmt.Sprintf(os.Getenv("GATEWAY_BASE_FRONTEND_URL")+"/v1/user/verification/token=%s&email=%s", emailVerificationToken, user.Email),
		),
	)
	u.handleEmailResponse(err, emailResponse)
	return emailVerificationToken, nil
}

func (u *UserController) Login(ctx *fiber.Ctx) error {
	var loginRequest request.LoginRequest
	var err error
	if err := ctx.BodyParser(&loginRequest); err != nil {
		log.Println("Error when binding data", err)
		return u.InvalidFormResponse(ctx, err)
	}

	if isValid, errors := validators.Validate(&loginRequest); !isValid {
		return u.SendBadRequestWithError(ctx, "Check invalid form before proceed", errors)
	}
	user, err := u.userDao.Find("email=?", loginRequest.Email)
	if err != nil {
		log.Println("Error when finding user ", err)
		if err == dao.ErrRecordNotFound {
			return u.SendNotFound(ctx, "Username or password incorrect")
		}
		return u.SendServerError(ctx, err)
	}

	if !common.ComparePassword(user.Password, loginRequest.Password) {
		return u.SendBadRequest(ctx, "Username or password incorrect")
	}
	token, err := u.generateToken(ctx, common.JSON{"user_id": user.ID}, time.Hour*24*3)
	if err != nil {
		return u.SendServerError(ctx, err)
	}
	return u.SendOK(ctx, response.LoginResponse{
		Token:     token,
		UserName:  user.UserName,
		Email:     user.Email,
		ID:        user.ID,
		Picture:   user.Picture,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, "Login success")
}

func (u *UserController) handleEmailResponse(err error, emailResponse *service.SendMailResponse) {
	if err != nil {
		log.Println("Error when sending email ", err)
	} else {
		if emailResponse.Success {
			log.Println("Send email success")
		} else {
			log.Println("Send email failed")
		}
	}
}

func (u *UserController) ActiveAccount(ctx *fiber.Ctx) error {
	var err error
	var activateAccountRequest request.ActivateAccountRequest
	if err = ctx.BodyParser(&activateAccountRequest); err != nil {
		log.Println("Error when binding data", err)
		return u.SendBadRequestWithError(ctx, "Error binding data", err)
	}
	if isValid, errors := validators.Validate(&activateAccountRequest); !isValid {
		return u.SendBadRequestWithError(ctx, "Check invalid form before proceed", errors)
	}
	var user *models.User
	if user, err = u.userDao.Find("email=?", activateAccountRequest.Email); err != nil {
		if err == dao.ErrRecordNotFound {
			return u.SendNotFound(ctx, "User not found")
		}
	}
	if user.Verified {
		return u.SendBadRequest(ctx, "User is already verified")
	}
	if user.VerificationToken != activateAccountRequest.Token {
		return u.SendBadRequest(ctx, "Invalid token")
	}
	claim, err := common.ExtractFromString(activateAccountRequest.Token)
	if err != nil {
		return u.SendBadRequest(ctx, "Invalid token")
	}
	if err := claim.Valid(); err != nil {
		return u.SendBadRequest(ctx, "Token is expired"+err.Error())
	}
	user.Verified = true
	if err := u.userDao.Save(user); err != nil {
		return u.SendServerError(ctx, err)
	}
	return u.SendOK(ctx, response.ActivateAccountResponse{}, "Account activated")
}

func (u *UserController) generateToken(ctx *fiber.Ctx, data common.JSON, duration time.Duration) (string, error) {
	emailVerificationToken, err := common.GenerateToken(data, duration)
	if err != nil {
		return "", u.SendServerError(ctx, err)
	}
	return emailVerificationToken, nil
}

func (u *UserController) AuthTest(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(uint)
	return u.SendOK(ctx, common.JSON{
		"message": "Auth success",
		"user_id": userId,
	}, "Auth test success")
}

func (u *UserController) ResendVerification(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(uint)
	user, err := u.userDao.Find("id=?", userId)
	if err != nil {
		if err == dao.ErrRecordNotFound {
			return u.SendNotFound(ctx, "Not found user")
		}
	}
	emailVerificationToken, err := u.sendVerification(ctx, *user)
	if err != nil {
		return u.SendServerError(ctx, err)
	}
	return u.SendOK(ctx, common.JSON{
		"token": emailVerificationToken,
		"email": user.Email,
	}, "Send verification success")
}
