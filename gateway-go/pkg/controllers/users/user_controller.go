package users

import (
	"ledungcobra/gateway-go/pkg/controllers/users/response"
	"ledungcobra/gateway-go/pkg/dao"
	"ledungcobra/gateway-go/pkg/htmltemplates"
	"ledungcobra/gateway-go/pkg/middlewares"
	"ledungcobra/gateway-go/pkg/service"
	"log"
	"time"

	"ledungcobra/gateway-go/pkg/common"
	"ledungcobra/gateway-go/pkg/controllers/base"
	"ledungcobra/gateway-go/pkg/controllers/users/request"
	"ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/validators"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	base.BaseController
	notificationService interfaces.INotificationService
	service             *service.UserService
}

func NewUserController(notificationService interfaces.INotificationService, userService *service.UserService) *UserController {
	this := &UserController{
		BaseController:      base.BaseController{},
		notificationService: notificationService,
		service:             userService,
	}
	return this
}

func (u *UserController) RegisterUserRouter(apiRouter fiber.Router) {
	userRouter := apiRouter.Group("/users")
	userRouter.Post("/register", u.Register)
	userRouter.Post("/login", u.Login)
	userRouter.Post("/activate", u.ActiveAccount)
	userRouter.Post("/sendVerification", middlewares.Protected, u.ResendVerification)
	userRouter.Get("/", u.FindAccount)
	userRouter.Post("/resetPassword", u.SendResetPassword)
	userRouter.Post("/verifyCode", u.VerifyCode)
	userRouter.Post("/changePassword", u.ChangePassword)
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

	user, err := u.service.Register(registerRequest)
	switch err {
	case service.ErrHashingPassword, service.ErrGenerateUniqueName:
		return u.SendServerError(ctx, err)
	case service.ErrDuplicateEmail:
		return u.SendBadRequest(ctx, err.Error())
	}
	emailVerificationToken, err := u.sendVerification(ctx, *user)
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
	user, err := u.service.FindByEmail(loginRequest.Email)
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
	err = u.service.Verify(activateAccountRequest.Email, activateAccountRequest.Token)
	if err != nil {
		switch err {
		case service.ErrUserAlreadyVerified, service.ErrInvalidToken:
			return u.SendBadRequest(ctx, err.Error())
		case service.ErrRecordNotfound:
			return u.SendNotFound(ctx, err.Error())
		default:
			return u.SendServerError(ctx, err)
		}
	}
	return u.SendOK(ctx, response.ActivateAccountResponse{}, "Account activated")
}

func (u *UserController) ResendVerification(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(uint)
	user, err := u.service.FindByID(userId)
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

func (u *UserController) FindAccount(ctx *fiber.Ctx) error {
	email := ctx.Query("email")
	user, err := u.service.FindByEmail(email)
	if err != nil {
		if err == dao.ErrRecordNotFound {
			return u.SendNotFound(ctx, "Not found email "+email)
		}
		return u.SendServerError(ctx, err)
	}
	return u.SendOK(ctx, common.JSON{
		"email":   user.Email,
		"picture": user.Picture,
	}, "Found user for email")
}

func (u *UserController) SendResetPassword(ctx *fiber.Ctx) error {
	var resetPasswordRequest request.ResetPasswordRequest
	if err := ctx.BodyParser(&resetPasswordRequest); err != nil {
		return u.InvalidFormResponse(ctx, err)
	}
	code, err := u.service.ResetPassword(resetPasswordRequest.Email)
	if err != nil {
		switch err {
		case dao.ErrRecordNotFound:
			return u.SendNotFound(ctx, "Not found user for email "+resetPasswordRequest.Email)
		default:
			return u.SendServerError(ctx, err)
		}
	}
	u.handleSendEmailResponse(u.notificationService.SendMail(resetPasswordRequest.Email, "Reset password ",
		htmltemplates.BuildResetPasswordTemplate(resetPasswordRequest.Email, code),
	))
	return u.SendOK(ctx, common.JSON{}, "Sent email to "+resetPasswordRequest.Email+" please check your mail to complete the process")
}

func (u *UserController) VerifyCode(ctx *fiber.Ctx) error {
	var codeVerificationRequest request.CodeVerificationRequest
	if err := ctx.BodyParser(&codeVerificationRequest); err != nil {
		return u.InvalidFormResponse(ctx, err)
	}
	if ok, errs := validators.Validate(codeVerificationRequest); !ok {
		return u.SendBadRequestWithError(ctx, "Form validation failed", errs)
	}
	if isValid := u.service.VerifyCode(codeVerificationRequest.Code, codeVerificationRequest.Email); !isValid {
		return u.SendBadRequest(ctx, "Invalid code")
	}
	user, _ := u.service.FindByEmail(codeVerificationRequest.Email)
	token, _ := u.generateToken(ctx, common.JSON{"user_id": user.ID}, time.Hour)
	return u.SendOK(ctx, common.JSON{
		"token": token,
	}, "Verify code successfully")
}

func (u *UserController) ChangePassword(ctx *fiber.Ctx) error {
	var changePasswordRequest request.ChangePasswordRequest
	if err := ctx.BodyParser(&changePasswordRequest); err != nil {
		return u.InvalidFormResponse(ctx, err)
	}
	if ok, errs := validators.Validate(changePasswordRequest); !ok {
		return u.SendBadRequestWithError(ctx, "Check your request body before proceed", errs)
	}
	err := u.service.ChangePassword(changePasswordRequest)
	if err != nil {
		switch err {
		case service.ErrHashingPassword:
			return u.SendServerError(ctx, err)
		case service.ErrInvalidToken:
			return u.SendBadRequest(ctx, "Invalid token")
		case service.ErrPasswordNotMatch:
			return u.SendBadRequest(ctx, "Password do not match")
		default:
			log.Println(err)
			return u.SendBadRequest(ctx, "An error occur when proceed change password")
		}

	}
	return u.SendOK(ctx, nil, "Change password success")
}
