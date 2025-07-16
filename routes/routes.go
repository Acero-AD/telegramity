package routes

import (
	"telegramity/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App) {
	// Initialize handlers
	userHandler := handlers.NewUserHandler()
	messageHandler := handlers.NewMessageHandler()

	// API routes group
	api := app.Group("/api/v1")

	// Health check routes
	api.Get("/health", handlers.HealthCheck)
	api.Get("/health/detailed", handlers.DetailedHealthCheck)

	// User routes
	users := api.Group("/users")
	users.Get("/", userHandler.GetUsers)
	users.Get("/:id", userHandler.GetUser)
	users.Post("/", userHandler.CreateUser)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)

	// Message routes
	messages := api.Group("/messages")
	messages.Get("/", messageHandler.GetMessages)
	messages.Get("/:id", messageHandler.GetMessage)
	messages.Post("/", messageHandler.CreateMessage)
	messages.Get("/user/:user_id", messageHandler.GetMessagesByUser)

	// Home route
	app.Get("/", handlers.Home)

	// 404 handler
	app.Use(handlers.NotFound)
}
