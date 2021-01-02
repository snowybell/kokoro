package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewCORSAllowOrigin(originURL string) fiber.Handler {
	return cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     originURL,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	})
}
