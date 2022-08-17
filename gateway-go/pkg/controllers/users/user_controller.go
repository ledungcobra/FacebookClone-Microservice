package users

import (
	"log"
	"time"

	"ledungcobra/gateway-go/pkg/common"
	"ledungcobra/gateway-go/pkg/config"
	"ledungcobra/gateway-go/pkg/controllers/base"
	"ledungcobra/gateway-go/pkg/controllers/users/request"
	"ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/models"
	"ledungcobra/gateway-go/pkg/validators"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userDao interfaces.IUserDAO
	base.BaseController
}

func NewUserController(userDao interfaces.IUserDAO, config *config.Config) *UserController {
	this := &UserController{userDao, base.BaseController{Config: config}}
	return this
}

func (userRoute *UserController) RegisterUserRouter(apiRouter fiber.Router) {
	userRouter := apiRouter.Group("/users")
	userRouter.Post("/register", userRoute.Register)
}

func (u *UserController) Register(ctx *fiber.Ctx) error {
	var err error
	var request request.RegisterRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Println("Error when binding data", err)
		return u.InvalidFormResponse(ctx, err)
	}

	if isValid, errors := validators.Validate(&request); !isValid {
		return u.SendBadRequest(ctx, "Check invalid form before proceed", errors)
	}

	user := registerRequestToUser(request)
	if user.Password, err = common.HashPassword(request.Password); err != nil {
		log.Println("Hash password error ", err)
		return u.SendServerError(ctx, err)
	}
	userName := request.FirstName + request.LastName
	if user.UserName, err = common.GenerateUniqueUserName(u.userDao, userName); err != nil {
		log.Println("Error when generating username ", err)
		return u.SendServerError(ctx, err)
	}

	if err := u.userDao.SaveUser(&user); err != nil {
		return u.SendServerError(ctx, err)
	}
	emailVerificationToken, err := common.GenerateToken(common.JSON{
		"user_id": user.ID,
	}, time.Hour)
	if err != nil {
		return u.SendServerError(ctx, err)
	}
	log.Print("Email verification token ", emailVerificationToken)
	return u.SendOk(ctx, user, "Create user success")
}

func registerRequestToUser(request request.RegisterRequest) models.User {
	return models.User{
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Email:      request.Email,
		BirthDay:   request.BirthDay,
		BirthYear:  request.BirthYear,
		BirthMonth: request.BirthMonth,
		Gender:     request.Gender,
	}
}
