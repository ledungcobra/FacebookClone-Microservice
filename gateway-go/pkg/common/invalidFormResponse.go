package common

import "github.com/gofiber/fiber/v2"

func InvalidFormResponse(c *fiber.Ctx) error {
	return c.Status(400).JSON(fiber.Map{
		"Message": "Form invalid",
	})
}
