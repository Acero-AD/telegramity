package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Home handles the home endpoint
func Home(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to Telegramity API",
		"status":  "running",
		"version": "1.0.0",
	})
}

// NotFound handles 404 errors
func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error":   "Not Found",
		"message": "The requested resource was not found",
		"path":    c.Path(),
	})
}
