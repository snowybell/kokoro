package main

import (
	"log"

	"github.com/snowybell/kokoro/router"

	"github.com/snowybell/kokoro/utils"

	r "github.com/snowybell/kokoro/repo"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Fiber app instance
	app := fiber.New()
	app.Use(cors.New())
	app.Use(requestid.New(), logger.New())

	// Prepare repository
	repo, err := r.NewRepoDefault()
	if err != nil {
		log.Panicf("can not prepare repo, err=%+v", err)
	}

	// Init JSONWebToken config
	jwtConfig, err := utils.NewJWTConfig()
	if err != nil {
		log.Panicf("can not prepare jwt, err=%+v", err)
	}

	// Setup routes and launch app
	router.SetupRoutes(app, jwtConfig, repo)
	log.Fatal(app.Listen(":3000"))
}
