package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/phamvlap/url-shortener/internal/config"
	"github.com/phamvlap/url-shortener/internal/errors"
	"github.com/phamvlap/url-shortener/internal/router"
)

func main() {
	// Load configuration
	config := config.GetConfig()

	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: errors.ErrorHandler,
	})

	// Set up middleware
	app.Use(logger.New())

	// Set up routes
	router.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":" + config.App.Port))
}
