package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snowybell/kokoro/handler"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Get("/ping", handler.Hello)

	auth := v1.Group("/auth")
	auth.Post("/login", handler.Login)
}
