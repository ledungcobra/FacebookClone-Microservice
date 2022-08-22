package middlewares

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"ledungcobra/gateway-go/pkg/common"
)

func Protected(ctx *fiber.Ctx) error {
	claim, err := common.ExtractClaim(ctx.Request())
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(common.JSON{
			"message": "Unauthorized",
		})
	}
	if err := claim.Valid(); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(common.JSON{
			"message": "Token is invalid",
		})
	}
	var jsonObject = common.JSON{}
	if err := json.Unmarshal([]byte(claim.Subject), &jsonObject); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.JSON{
			"message": "Cannot extract subject",
		})
	}
	userId := uint(jsonObject["user_id"].(float64))
	ctx.Locals("user_id", userId)
	return ctx.Next()
}
