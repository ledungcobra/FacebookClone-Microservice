package middlewares

import "github.com/gofiber/fiber/v2"

func notFoundMiddleware(c *fiber.Ctx) error {
	return c.Status(404).JSON(fiber.Map{
		"message": "Not found",
	})
}
