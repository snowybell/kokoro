package auth

import (
	"errors"
	"time"

	"github.com/snowybell/kokoro/response"

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
			return response.
				Error(ctx).
				WithMessage("bad request").End()
		}

		user, err := repo.GetUser(entity.User{
			Username: input.Username,
			Password: input.Password,
		})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.
				Error(ctx).
				WithMessage("username or password is not correct").End()
		}

		// Issuing token
		token := jwt.New(jwt.SigningMethodHS256)
		claim := token.Claims.(jwt.MapClaims)
		claim["id"] = user.ID
		claim["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Signing token
		tokenString, err := token.SignedString(jwtConfig.SecretKey)
		if err != nil {
			return response.
				Error(ctx).
				WithCode(fiber.StatusInternalServerError).
				WithMessage("internal server error").End()
		}

		return response.
			Success(ctx).
			WithData(fiber.Map{"token": tokenString}).End()
	}
}
