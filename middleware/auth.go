package middleware

import (
	"net/http"

	"golang.org/x/oauth2"

	"github.com/snowybell/kokoro/entity"
	"github.com/snowybell/kokoro/repo"

	jwt "github.com/form3tech-oss/jwt-go"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/snowybell/kokoro/response"
	"github.com/snowybell/kokoro/utils"
)

func Protected(jwtConfig *utils.JWTConfig, repo repo.Repository) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtConfig.SecretKey,
		ErrorHandler:   jwtError,
		SuccessHandler: IssuePassport(repo),
	})
}

func IssuePassport(repo repo.Repository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Locals(CtxUserKey).(*jwt.Token)
		claim := token.Claims.(jwt.MapClaims)

		// Using uint(...(float64)) so TRICKY!
		// NOTE: Find another solution
		id := uint(claim["id"].(float64))
		user, err := repo.GetUserByID(id)

		// Any errors when parsing credential
		if err != nil {
			return response.
				Error(ctx).
				WithCode(fiber.StatusBadRequest).
				WithMessage("invalid credentials").End()
		}

		uToken, err := repo.GetTokenByID(user.TokenID)
		if err != nil {
			return response.
				Error(ctx).
				WithCode(fiber.StatusInternalServerError).
				WithMessage("can not find token attached to user").End()
		}

		// Issue passport
		passport := &entity.Passport{
			User: user,
			OAuthToken: &oauth2.Token{
				Expiry:       uToken.Expiry,
				TokenType:    uToken.TokenType,
				AccessToken:  uToken.AccessToken,
				RefreshToken: uToken.RefreshToken,
			},
		}

		ctx.Locals(CtxPassportKey, passport)
		return ctx.Next()
	}
}

func jwtError(ctx *fiber.Ctx, err error) error {
	return response.
		Error(ctx).
		WithCode(http.StatusForbidden).
		WithMessage(err.Error()).End()
}
