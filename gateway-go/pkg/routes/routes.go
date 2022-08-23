package routes

import (
	"github.com/gofiber/fiber/v2"
	"ledungcobra/gateway-go/pkg/cloudinary"
	"ledungcobra/gateway-go/pkg/controllers/posts"
	"ledungcobra/gateway-go/pkg/controllers/users"
	"ledungcobra/gateway-go/pkg/dao"
	"ledungcobra/gateway-go/pkg/database"
	"ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/models"
	"ledungcobra/gateway-go/pkg/service"
)

func SetUpRoutes(connector *database.SQLConnector, app *fiber.App, notificationService interfaces.INotificationService) {
	db := connector.GetDatabase()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	userDao := dao.NewUserDao(db)
	codeDao := dao.NewCommonDao[models.Code](db)
	postDao := dao.NewCommonDao[models.Post](db)

	postService := service.NewPostService(postDao, userDao, db)
	userService := service.NewUserService(userDao, codeDao)

	users.NewUserController(notificationService, userService).
		RegisterUserRouter(v1)
	posts.NewPostsController(postService, userService, cloudinary.NewCloudinaryService()).RegisterRoutes(v1)
}
