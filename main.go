package main

import (
	"log"

	"telegramity/config"
	"telegramity/middleware"
	"telegramity/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Load configuration
	config := config.LoadConfig()

	// Create new Fiber app
	app := fiber.New(fiber.Config{
		AppName: config.App.Name,
	})

	// Global middleware (applied to ALL routes)
	app.Use(logger.New())                     // Built-in logger
	app.Use(cors.New())                       // Built-in CORS
	app.Use(middleware.RequestIDMiddleware()) // Custom request ID

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	log.Printf("Server starting on port %s", config.Port)
	log.Fatal(app.Listen(":" + config.Port))
}
