package auth

import (
	"errors"

	"github.com/snowybell/kokoro/response"

	"github.com/gofiber/fiber/v2"
	"github.com/snowybell/kokoro/entity"
	"github.com/snowybell/kokoro/repo"
	"github.com/snowybell/kokoro/utils"
	"gorm.io/gorm"
)

type RegisterInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Register(repo repo.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input RegisterInput
		if err := utils.ShouldBind(ctx, &input); err != nil {
			return response.
				Error(ctx).
				WithMessage("bad request").
				End()
		}

		_, err := repo.GetUserByUsername(input.Username)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return response.
				Error(ctx).
				WithMessage("username has already been taken").
				End()
		}

		user, err := repo.SaveUser(entity.User{
			Name:     input.Name,
			Email:    input.Email,
			Username: input.Username,
			Password: input.Password,
		})
		if err != nil || user == nil {
			return response.
				Error(ctx).
				WithCode(fiber.StatusInternalServerError).
				WithMessage("internal server error").
				End()
		}

		return response.
			Success(ctx).
			WithData(fiber.Map{"createdAt": user.CreatedAt}).
			End()
	}
}
