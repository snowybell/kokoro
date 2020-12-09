package auth

import (
	"errors"

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
			return ctx.
				Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"error": "bad request"})
		}

		_, err := repo.GetUserByUsername(input.Username)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.
				Status(fiber.StatusOK).
				JSON(fiber.Map{"error": "username has already been taken"})
		}

		user, err := repo.SaveUser(entity.User{
			Name:     input.Name,
			Email:    input.Email,
			Username: input.Username,
			Password: input.Password,
		})
		if err != nil || user == nil {
			return ctx.
				Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": "internal server error"})
		}

		return ctx.
			Status(fiber.StatusOK).
			JSON(fiber.Map{
				"success":   true,
				"createdAt": user.CreatedAt,
			})
	}
}
