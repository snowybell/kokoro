package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/snowybell/kokoro/handler"
	Auth "github.com/snowybell/kokoro/handler/auth"
	OAuth "github.com/snowybell/kokoro/handler/oauth"
	"github.com/snowybell/kokoro/middleware"
	r "github.com/snowybell/kokoro/repo"
	"github.com/snowybell/kokoro/utils"
	"go.uber.org/fx"
	"golang.org/x/oauth2"
)

func Register(app *fiber.App, repo r.Repository,
	jwtConfig *utils.JWTConfig, gOAuthConfig *oauth2.Config) {
	v1 := app.Group("/v1")
	v1.Get("/me", middleware.Protected(jwtConfig), handler.Me)

	// Authentication
	auth := v1.Group("/auth")
	auth.Post("/login", Auth.Login(jwtConfig, repo))
	auth.Post("/register", Auth.Register(repo))

	// OAuth
	oauth := v1.Group("/oauth")
	oauth.Get("/google", OAuth.GoogleLoginRedirect(gOAuthConfig))
	oauth.Get("/google/callback", OAuth.GoogleLoginCallback(gOAuthConfig))
}

func SetupRoutes(app *fiber.App) {
	fxApp := fx.New(
		fx.Provide(
			// inject app instance
			func() *fiber.App {
				return app
			},
			r.NewRepoDefault,
			utils.NewJWTConfig,
			utils.NewGoogleOAuthConfig,
		),
		fx.Invoke(Register),
		fx.NopLogger,
	)

	if err := fxApp.Err(); err != nil {
		log.Fatalf("can not setup routes, err=%s", err.Error())
	}
}
