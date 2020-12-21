package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowybell/kokoro/handler"
	Auth "github.com/snowybell/kokoro/handler/auth"
	"github.com/snowybell/kokoro/middleware"
	r "github.com/snowybell/kokoro/repo"
	"github.com/snowybell/kokoro/utils"
)

func SetupRoutes(app *fiber.App, jwtConfig *utils.JWTConfig, repo r.Repository) {
	v1 := app.Group("/v1")
	v1.Get("/ping", handler.Hello)
	v1.Get("/me", middleware.Protected(jwtConfig), handler.Me)

	// Authentication
	auth := v1.Group("/auth")
	auth.Post("/login", Auth.Login(jwtConfig, repo))
	auth.Post("/register", Auth.Register(repo))
}
