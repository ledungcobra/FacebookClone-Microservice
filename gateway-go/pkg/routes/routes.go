package routes

import (
	"ledungcobra/gateway-go/pkg/config"
	"ledungcobra/gateway-go/pkg/controllers/users"
	"ledungcobra/gateway-go/pkg/dao"
	"ledungcobra/gateway-go/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(db *database.SQLDBManager, app *fiber.App, config *config.Config) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	userDao := dao.NewUserDao(db.GetDatabase())
	userController := users.NewUserController(userDao, config)
	userController.RegisterUserRouter(v1)
}
