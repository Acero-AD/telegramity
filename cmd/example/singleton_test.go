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
	"github.com/somosbytes/telegramity"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatIDStr := os.Getenv("TELEGRAM_CHAT_ID")
	chatID, _ := strconv.ParseInt(chatIDStr, 10, 64)

	fmt.Println("🚀 Testing Singleton Patterns...")

	// Test 1: Global Singleton Pattern
	fmt.Println("\n📤 Test 1: Global Singleton Pattern")
	testGlobalSingleton(botToken, chatID)

	// Test 2: Manager Pattern
	fmt.Println("\n📤 Test 2: Manager Pattern")
	testManagerPattern(botToken, chatID)

	fmt.Println("\n🎉 All singleton tests completed!")
}

func testGlobalSingleton(botToken string, chatID int64) {
	// Initialize once at app startup
	err := telegramity.InitGlobalClient(
		botToken,
		chatID,
		telegramity.WithEnvironmentName("production"),
		telegramity.WithAppInfo("SingletonApp", "1.0.0"),
	)
	if err != nil {
		log.Printf("Failed to init global client: %v", err)
		return
	}

	// Use anywhere in your code without creating new instances
	client := telegramity.GetGlobalClient()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.ReportError(ctx, errors.New("singleton test error"), telegramity.ErrorTypeInternal)
	if err != nil {
		log.Printf("Failed to report error: %v", err)
	} else {
		fmt.Println("✅ Global singleton error reported!")
	}

	// Clean up at app shutdown
	defer telegramity.CloseGlobalClient()
}

func testManagerPattern(botToken string, chatID int64) {
	// Get the default manager
	manager := telegramity.GetDefaultManager()

	// Initialize once
	err := manager.Init(
		botToken,
		chatID,
		telegramity.WithEnvironmentName("production"),
		telegramity.WithAppInfo("ManagerApp", "1.0.0"),
	)
	if err != nil {
		log.Printf("Failed to init manager: %v", err)
		return
	}

	// Use anywhere in your code
	client := manager.GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.ReportError(ctx, errors.New("manager test error"), telegramity.ErrorTypeInternal)
	if err != nil {
		log.Printf("Failed to report error: %v", err)
	} else {
		fmt.Println("✅ Manager pattern error reported!")
	}

	// Clean up at app shutdown
	defer manager.Close()
}
