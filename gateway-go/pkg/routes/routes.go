package routes

import (
	"github.com/gofiber/fiber/v2"
	"ledungcobra/gateway-go/pkg/controllers/users"
	"ledungcobra/gateway-go/pkg/dao"
	"ledungcobra/gateway-go/pkg/database"
	"ledungcobra/gateway-go/pkg/interfaces"
)

func SetUpRoutes(db *database.SQLConnector, app *fiber.App, notificationService interfaces.INotificationService) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	userDao := dao.NewUserDao(db.GetDatabase())
	userController := users.NewUserController(userDao, notificationService)
	userController.RegisterUserRouter(v1)
}
