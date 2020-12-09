package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowybell/kokoro/utils"
)

type LoginInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(ctx *fiber.Ctx) error {
	var input LoginInput
	if err := utils.ShouldBind(ctx, &input); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "username and password is required"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}
