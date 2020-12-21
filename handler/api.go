package handler

import (
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/snowybell/kokoro/entity"
	"gorm.io/gorm"
)

func Hello(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"hello": "world"})
}

func Me(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*jwt.Token)
	claim := token.Claims.(jwt.MapClaims)

	id := uint(claim["id"].(float64))
	user := entity.User{Model: gorm.Model{ID: id}}

	return ctx.
		Status(fiber.StatusOK).
		JSON(fiber.Map{"me": true, "id": user.ID})
}
