package main

import (
	"log"

	"github.com/snowybell/kokoro/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber app instance
	app := fiber.New()

	// Setup routes and launch app
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3001"))
}
