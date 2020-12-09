package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowybell/kokoro/handler"
	r "github.com/snowybell/kokoro/repo"
)

func SetupRoutes(app *fiber.App, repo r.Repository) {
	v1 := app.Group("/v1")
	v1.Get("/ping", handler.Hello)

	auth := v1.Group("/auth")
	auth.Post("/login", handler.Login(repo))
}
