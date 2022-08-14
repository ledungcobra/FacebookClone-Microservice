package routes

import "github.com/gofiber/fiber/v2"

func SetUpUserRouter(apiRouter fiber.Router) {
	userRouter := apiRouter.Group("/users")

	userRouter.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from user")
	})
}
