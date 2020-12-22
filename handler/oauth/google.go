package oauth

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

func GoogleLoginRedirect(cfg *oauth2.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		state := strconv.FormatInt(time.Now().Unix(), 10)
		redirectURL := cfg.AuthCodeURL(state)
		return ctx.Redirect(redirectURL, fiber.StatusMovedPermanently)
	}
}

func GoogleLoginCallback(cfg *oauth2.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}
