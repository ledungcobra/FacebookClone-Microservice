package users

import (
	"ledungcobra/gateway-go/pkg/common"
	"ledungcobra/gateway-go/pkg/config"
	"ledungcobra/gateway-go/pkg/controllers/users/request"
	"ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/models"
	"ledungcobra/gateway-go/pkg/validators"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userDao interfaces.IUserDAO
	config  *config.Config
}

func NewUserController(userDao interfaces.IUserDAO, config *config.Config) *UserController {
	this := &UserController{userDao, config}
	return this
}

func (userRoute *UserController) RegisterUserRouter(apiRouter fiber.Router) {
	userRouter := apiRouter.Group("/users")
	userRouter.Post("/register", userRoute.Register)
}

func (userRoute *UserController) Register(ctx *fiber.Ctx) error {
	var request request.RegisterRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Println("Error when binding data", err)
		return common.InvalidFormResponse(ctx)
	}
	if isValid, errors := validators.Validate(&request); !isValid {
		return ctx.Status(400).JSON(
			fiber.Map{
				"message": "Form input error",
				"errors":  errors,
			})
	}
	mappedUser := models.User{
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Email:      request.Email,
		BirthDay:   request.BirthDay,
		BirthYear:  request.BirthYear,
		BirthMonth: request.BirthMonth,
		Gender:     request.Gender,
	}
	var err error
	mappedUser.Password, err = common.HashPassword(request.Password)
	if err != nil {
		log.Println("Hash password error ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Password hashing error",
		})
	}
	if err := userRoute.userDao.SaveUser(&mappedUser); err != nil {
		log.Println("Cannot save user ", err)
		
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Cannot save user",
			"error":   err.Error(),
		})
	}
	return ctx.JSON(mappedUser)
}
