package handler

import (
	"github.com/gofiber/fiber/v2"
	mw "github.com/snowybell/kokoro/middleware"
	"github.com/snowybell/kokoro/response"
)

func Me(ctx *fiber.Ctx) error {
	passport := ctx.Locals(mw.CtxPassportKey)
	return response.
		Success(ctx).
		WithData(fiber.Map{"passport": passport}).End()
}
