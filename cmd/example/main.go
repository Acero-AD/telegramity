package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/somosbytes/Telegramity/internal/telegram/bot"
)

func main() {
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

	fmt.Println("🤖 Creating Telegram bot client...")

	client, err := bot.NewBotClient(botToken, 10*time.Second)
	if err != nil {
		log.Fatalf("Failed to create bot client: %v", err)
	}

	fmt.Println("✅ Bot client created successfully!")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("🔍 Testing bot connection...")
	err = client.TestConnection(ctx)
	if err != nil {
		log.Fatalf("Failed to test connection: %v", err)
	}
	fmt.Println("✅ Bot connection test successful!")

	fmt.Println("📤 Sending test message...")
	testMessage := fmt.Sprintf("🚀 Hello from Telegramity SDK!\n\nTime: %s\nThis is a test message from your Go SDK.", time.Now().Format("2006-01-02 15:04:05"))

	err = client.SendMessage(ctx, chatID, testMessage)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println("✅ Test message sent successfully!")
	fmt.Println("📱 Check your Telegram chat to see the message!")
}
