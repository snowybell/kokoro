package router

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/gofiber/fiber/v2"
	"github.com/snowybell/kokoro/handler"
	OAuth "github.com/snowybell/kokoro/handler/oauth"
	mw "github.com/snowybell/kokoro/middleware"
	r "github.com/snowybell/kokoro/repo"
	"github.com/snowybell/kokoro/utils"
	"go.uber.org/fx"
	"golang.org/x/oauth2"
)

func Register(app *fiber.App, repo r.Repository,
	jwtConfig *utils.JWTConfig, gOAuthConfig *oauth2.Config, feConfig *utils.FrontendConfig) {
	// Setup mandatory middlewares
	app.Use(logger.New(), requestid.New(),
		utils.NewCORSAllowOrigin(feConfig.URL))

	// Group version /v1
	v1 := app.Group("/v1")
	v1.Get("/me", mw.Protected(jwtConfig, repo), handler.Me)

	// OAuth
	oauth := v1.Group("/oauth")
	oauth.Get("/google", OAuth.GoogleLoginRedirect(gOAuthConfig))
	oauth.Post("/google/callback", OAuth.GoogleLoginCallback(gOAuthConfig, jwtConfig, repo))
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
			utils.NewFrontendConfig,
			utils.NewGoogleOAuthConfig,
		),
		fx.Invoke(Register),
		fx.NopLogger,
	)

	if err := fxApp.Err(); err != nil {
		log.Fatalf("can not setup routes, err=%s", err.Error())
	}
}
