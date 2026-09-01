package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/phamvlap/url-shortener/internal/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handler.CheckHealthHandler)
	app.Get("/:shortID", handler.RedirectURLHandler)
	app.Post("/api/v1/url/shorten", handler.ShortenURLHandler)
}
