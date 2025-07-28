package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/somosbytes/telegramity/pkg/telegramity"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get configuration from environment
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("Please set TELEGRAM_BOT_TOKEN environment variable")
	}

	chatIDStr := os.Getenv("TELEGRAM_CHAT_ID")
	if chatIDStr == "" {
		log.Fatal("Please set TELEGRAM_CHAT_ID environment variable")
	}

	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Invalid chat ID: %v", err)
	}

	fmt.Println("🚀 Testing new error handling with global singleton...")

	// Initialize global client once
	err = telegramity.InitGlobalClient(
		botToken,
		chatID,
		telegramity.WithEnvironmentName("production"),
		telegramity.WithAppInfo("TestApp", "1.0.0"),
		telegramity.WithTimeout(10*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to initialize global client: %v", err)
	}
	defer func() {
		if err := telegramity.CloseGlobalClient(); err != nil {
			log.Printf("Failed to close global client: %v", err)
		}
	}()

	fmt.Println("✅ Global client initialized successfully!")

	// Get the global client
	client := telegramity.GetGlobalClient()

	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test 1: Simple error with error type
	fmt.Println("\n📤 Test 1: Simple error with error type")
	err = client.ReportError(ctx, errors.New("database connection failed"), telegramity.ErrorTypeDatabase)
	if err != nil {
		log.Printf("Failed to report error: %v", err)
	} else {
		fmt.Println("✅ Database error reported successfully!")
	}

	// Test 2: Error with context
	fmt.Println("\n📤 Test 2: Error with context")
	err = client.ReportErrorWithContext(
		ctx,
		errors.New("user authentication failed"),
		telegramity.ErrorTypeAuth,
		map[string]interface{}{
			"user_id": "user123",
			"ip":      "192.168.1.100",
			"action":  "login",
		},
	)
	if err != nil {
		log.Printf("Failed to report error: %v", err)
	} else {
		fmt.Println("✅ Auth error with context reported successfully!")
	}

	// Test 3: Network error
	fmt.Println("\n📤 Test 3: Network error")
	err = client.ReportError(ctx, errors.New("API request timeout"), telegramity.ErrorTypeNetwork)
	if err != nil {
		log.Printf("Failed to report error: %v", err)
	} else {
		fmt.Println("✅ Network error reported successfully!")
	}

	fmt.Println("\n🎉 All tests completed! Check your Telegram chat for the error reports.")
}
