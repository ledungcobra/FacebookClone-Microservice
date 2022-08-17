package base

import (
	"ledungcobra/gateway-go/pkg/config"

	"github.com/gofiber/fiber/v2"
)

type BaseController struct {
	Config *config.Config
}

func (b BaseController) InvalidFormResponse(c *fiber.Ctx, err error) error {
	return c.Status(400).JSON(fiber.Map{
		"message": "Missing form field",
		"errors":  []string{err.Error()},
	})
}

func (b BaseController) SendOk(c *fiber.Ctx, data any, message string) error {
	return c.Status(200).JSON(fiber.Map{
		"message": message,
		"errors":  nil,
		"data":    data,
	})
}

func (b BaseController) SendBadRequest(ctx *fiber.Ctx, message string, errors any) error {
	return ctx.Status(400).JSON(fiber.Map{
		"message": message,
		"errors":  errors,
		"data":    nil,
	})
}

func (b BaseController) SendServerError(ctx *fiber.Ctx, err error) error {
	cfg := config.Cfg
	if cfg.Env == "production" {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Server internal error",
			"errors":  nil,
			"data":    nil,
		})
	}

	return ctx.Status(500).JSON(fiber.Map{
		"message": "Server internal error",
		"errors":  []string{err.Error()},
		"data":    nil,
	})
}
