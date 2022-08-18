package users

import (
	"fmt"
	"ledungcobra/gateway-go/pkg/controllers/users/response"
	"ledungcobra/gateway-go/pkg/htmltemplates"
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
}

func (u *UserController) Register(ctx *fiber.Ctx) error {
	var err error
	var registerRequest request.RegisterRequest
	if err := ctx.BodyParser(&registerRequest); err != nil {
		log.Println("Error when binding data", err)
		return u.InvalidFormResponse(ctx, err)
	}

	if isValid, errors := validators.Validate(&registerRequest); !isValid {
		return u.SendBadRequest(ctx, "Check invalid form before proceed", errors)
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
			return u.SendBadRequest(ctx, "Email is already exist", err)
		}
		return u.SendServerError(ctx, err)
	}
	emailVerificationToken, err := common.GenerateToken(common.JSON{
		"user_id": user.ID,
	}, time.Hour)
	if err != nil {
		return u.SendServerError(ctx, err)
	}

	emailResponse, err := u.notificationService.SendMail(user.Email, "Verification Email",
		htmltemplates.BuildRegistrationTemplate(user.UserName,
			fmt.Sprintf(os.Getenv("GATEWAY_BASE_FRONTEND_URL")+"/v1/user/verification/token=%s&email=%s", emailVerificationToken, user.Email),
		),
	)
	if err != nil {
		return u.SendServerError(ctx, err)
	}
	if emailResponse.Success {
		log.Println("Send email success")
	} else {
		log.Println("Send email failed")
	}
	return u.SendCreated(ctx, response.RegisterResponse{
		Success:  true,
		UserName: user.UserName,
		UserID:   user.ID,
		Token:    emailVerificationToken,
	}, "Create user success")
}
