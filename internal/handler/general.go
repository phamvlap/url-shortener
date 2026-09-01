package handler

import (
	"github.com/gofiber/fiber/v3"
)

func CheckHealthHandler(c fiber.Ctx) error {
	data := map[string]string{
		"health": "ok",
	}
	return c.Status(fiber.StatusOK).JSON(data)
}
