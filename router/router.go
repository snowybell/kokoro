package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowybell/kokoro/handler"
	Auth "github.com/snowybell/kokoro/handler/auth"
	r "github.com/snowybell/kokoro/repo"
	"github.com/snowybell/kokoro/utils"
)

func SetupRoutes(app *fiber.App, jwt utils.JWT, repo r.Repository) {
	v1 := app.Group("/v1")
	v1.Get("/ping", handler.Hello)

	// Authentication
	auth := v1.Group("/auth")
	auth.Post("/login", Auth.Login(jwt, repo))
	auth.Post("/register", Auth.Register(repo))
}
