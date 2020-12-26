package oauth

import (
	"context"
	"strconv"
	"time"

	"google.golang.org/api/option"

	"github.com/snowybell/kokoro/response"

	"github.com/snowybell/kokoro/utils"
	gOAuth2 "google.golang.org/api/oauth2/v2"

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

type LoginCallBackInput struct {
	State  string `query:"state" validate:"required"`
	Code   string `query:"code" validate:"required"`
	Scopes string `query:"scope" validate:"required"`
}

func GoogleLoginCallback(cfg *oauth2.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input LoginCallBackInput
		if err := utils.ShouldBindQuery(ctx, &input); err != nil {
			return response.
				Error(ctx).
				WithMessage("bad request").End()
		}

		bgCtx := context.Background()
		token, err := cfg.Exchange(bgCtx, input.Code)
		if err != nil {
			return response.
				Error(ctx).
				WithMessage(err.Error()).End()
		}

		gService, err := gOAuth2.NewService(
			bgCtx,
			option.WithScopes(cfg.Scopes...),
			option.WithTokenSource(cfg.TokenSource(bgCtx, token)))
		if err != nil {
			return response.
				Error(ctx).
				WithCode(fiber.StatusInternalServerError).
				WithMessage(err.Error()).End()
		}

		tokenInfo, err := gService.Tokeninfo().Do()
		if err != nil {
			return response.
				Error(ctx).
				WithCode(fiber.StatusInternalServerError).
				WithMessage(err.Error()).End()
		}

		return response.
			Success(ctx).
			WithData(fiber.Map{"info": tokenInfo}).End()
	}
}
