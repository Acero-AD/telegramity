package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// HealthCheck handles the health check endpoint
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "healthy",
		"message": "Service is running",
	})
}

// DetailedHealthCheck provides more detailed health information
func DetailedHealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":    "healthy",
		"message":   "Service is running",
		"timestamp": time.Now(),
		"version":   "1.0.0",
		"services": fiber.Map{
			"database": "connected",
			"redis":    "connected",
			"api":      "running",
		},
	})
}
