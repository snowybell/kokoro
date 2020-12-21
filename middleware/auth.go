package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/snowybell/kokoro/response"
	"github.com/snowybell/kokoro/utils"
)

func Protected(jwtConfig *utils.JWTConfig) func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtConfig.SecretKey,
		ErrorHandler: jwtError,
	})
}

func jwtError(ctx *fiber.Ctx, err error) error {
	return response.Error(ctx).WithMessage(err.Error()).End()
}
