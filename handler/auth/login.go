package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gofiber/fiber/v2"
	"github.com/snowybell/kokoro/entity"
	r "github.com/snowybell/kokoro/repo"
	"github.com/snowybell/kokoro/utils"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(jwtConfig *utils.JWTConfig, repo r.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input LoginInput
		if err := utils.ShouldBind(ctx, &input); err != nil {
			return ctx.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"error": "bad request"})
		}

		user, err := repo.GetUser(entity.User{
			Username: input.Username,
			Password: input.Password,
		})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.
				Status(fiber.StatusOK).
				JSON(fiber.Map{"error": "username or password is not correct"})
		}

		// Issuing token
		token := jwt.New(jwt.SigningMethodHS256)
		claim := token.Claims.(jwt.MapClaims)
		claim["id"] = user.ID
		claim["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Signing token
		tokenString, err := token.SignedString([]byte(jwtConfig.SecretKey))
		if err != nil {
			return ctx.
				Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": "internal server error"})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"data": fiber.Map{
				"token": tokenString,
			},
		})
	}
}
