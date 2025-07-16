package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// CustomLogger creates a custom logger middleware
func CustomLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request (this calls the next middleware/handler)
		err := c.Next()

		// Log request details AFTER the request is processed
		duration := time.Since(start)
		c.Append("X-Response-Time", duration.String())

		// Log the request details
		log.Printf(
			"[%s] %s %s - %d - %v",
			c.Method(),
			c.Path(),
			c.IP(),
			c.Response().StatusCode(),
			duration,
		)

		return err
	}
}

// AuthMiddleware is a placeholder for authentication middleware
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		token := c.Get("Authorization")

		// Check if token exists
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization token required",
			})
		}

		// In a real application, you would validate the token here
		// For example: validate JWT, check API key, etc.

		// If token is valid, continue to the next handler
		// If not valid, return error (don't call c.Next())

		// For demo purposes, let's check if token starts with "Bearer "
		if len(token) < 7 || token[:7] != "Bearer " {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token format",
			})
		}

		// Token looks valid, continue
		return c.Next()
	}
}

// RateLimitMiddleware is a placeholder for rate limiting
func RateLimitMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get client IP
		clientIP := c.IP()

		// In a real application, you would:
		// 1. Check Redis/database for request count
		// 2. Increment counter
		// 3. Check if limit exceeded

		// For demo purposes, let's simulate rate limiting
		// You could store this in memory, Redis, or database
		requestCount := getRequestCount(clientIP)

		if requestCount > 100 { // 100 requests per minute
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Rate limit exceeded",
				"limit": "100 requests per minute",
			})
		}

		// Increment counter
		incrementRequestCount(clientIP)

		// Add rate limit headers
		c.Append("X-RateLimit-Limit", "100")
		c.Append("X-RateLimit-Remaining", string(rune(100-requestCount)))

		return c.Next()
	}
}

// RequestIDMiddleware adds a unique request ID to each request
func RequestIDMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Generate or get request ID
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}

		// Add to response headers
		c.Append("X-Request-ID", requestID)

		// Add to context for handlers to use
		c.Locals("requestID", requestID)

		return c.Next()
	}
}

// ValidationMiddleware validates request body
func ValidationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if request has content
		if c.Method() == "POST" || c.Method() == "PUT" {
			contentType := c.Get("Content-Type")
			if contentType != "application/json" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Content-Type must be application/json",
				})
			}
		}

		return c.Next()
	}
}

// Helper functions (in a real app, these would use Redis/database)
var requestCounts = make(map[string]int)

func getRequestCount(clientIP string) int {
	return requestCounts[clientIP]
}

func incrementRequestCount(clientIP string) {
	requestCounts[clientIP]++
}

func generateRequestID() string {
	return "req-" + time.Now().Format("20060102150405")
}
