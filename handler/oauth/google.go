package oauth

import (
	"context"
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/dgrijalva/jwt-go"
	"github.com/snowybell/kokoro/entity"

	"github.com/snowybell/kokoro/repo"

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
		redirectURL := cfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
		return ctx.Redirect(redirectURL, fiber.StatusMovedPermanently)
	}
}

type LoginCallBackInput struct {
	State  string `query:"state" validate:"required"`
	Code   string `query:"code" validate:"required"`
	Scopes string `query:"scope" validate:"required"`
}

func GoogleLoginCallback(cfg *oauth2.Config, jwtConfig *utils.JWTConfig, repo repo.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input LoginCallBackInput
		if err := utils.ShouldBindQuery(ctx, &input); err != nil {
			return response.
				Error(ctx).
				WithMessage("bad request").End()
		}

		bgCtx := context.Background()
		oAuthToken, err := cfg.Exchange(bgCtx, input.Code)
		if err != nil {
			return response.
				Error(ctx).
				WithMessage(err.Error()).End()
		}

		gService, err := gOAuth2.NewService(
			bgCtx,
			option.WithScopes(cfg.Scopes...),
			option.WithTokenSource(cfg.TokenSource(bgCtx, oAuthToken)))
		if err != nil {
			return response.
				Error(ctx).
				WithCode(fiber.StatusInternalServerError).
				WithMessage(err.Error()).End()
		}

		gUser, err := gService.Userinfo.V2.Me.Get().Do()
		if err != nil {
			return response.
				Error(ctx).
				WithCode(fiber.StatusInternalServerError).
				WithMessage(err.Error()).End()
		}

		user, err := repo.GetUserByEmail(gUser.Email)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			token, err := repo.SaveToken(entity.Token{
				Expiry:       oAuthToken.Expiry,
				TokenType:    oAuthToken.TokenType,
				AccessToken:  oAuthToken.AccessToken,
				RefreshToken: oAuthToken.RefreshToken,
			})
			if err != nil {
				return response.
					Error(ctx).
					WithCode(fiber.StatusInternalServerError).
					WithMessage("save access token fail").End()
			}

			user, err = repo.SaveUser(entity.User{
				Name:    gUser.Name,
				Email:   gUser.Email,
				TokenID: token.ID,
			})
			if err != nil {
				return response.
					Error(ctx).
					WithCode(fiber.StatusInternalServerError).
					WithMessage("create user fail").End()
			}
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
				WithMessage("can not sign a token").End()
		}

		return response.
			Success(ctx).
			WithData(fiber.Map{"token": tokenString}).End()
	}
}
