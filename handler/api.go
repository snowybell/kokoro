package handler

import (
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/snowybell/kokoro/entity"
	"github.com/snowybell/kokoro/response"
	"gorm.io/gorm"
)

func Me(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*jwt.Token)
	claim := token.Claims.(jwt.MapClaims)

	id := uint(claim["id"].(float64))
	user := entity.User{Model: gorm.Model{ID: id}}

	return response.Success(ctx).
		WithData(fiber.Map{"id": user.ID}).End()
}
