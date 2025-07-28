package unit

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/somosbytes/telegramity/pkg/telegramity"
)

// TestSingletonIntegration tests the actual singleton behavior with real bot tokens
// This test will be skipped if no real bot token is available
func TestSingletonIntegration(t *testing.T) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatID == "" {
		t.Skip("Skipping integration test: TELEGRAM_BOT_TOKEN or TELEGRAM_CHAT_ID not set")
	}

	t.Run("successful_initialization", func(t *testing.T) {
		telegramity.CloseGlobalClient()

		err := telegramity.InitGlobalClient(botToken, 123456789) // Using fake chat ID for safety
		if err != nil {
			t.Fatalf("Failed to initialize client: %v", err)
		}

		client := telegramity.GetGlobalClient()
		if client == nil {
			t.Errorf("Expected client but got nil")
		}

		err = telegramity.CloseGlobalClient()
		if err != nil {
			t.Errorf("Failed to close client: %v", err)
		}
	})

	t.Run("multiple_initialization_calls", func(t *testing.T) {
		telegramity.CloseGlobalClient()

		err1 := telegramity.InitGlobalClient(botToken, 123456789)
		if err1 != nil {
			t.Fatalf("First initialization failed: %v", err1)
		}

		client1 := telegramity.GetGlobalClient()

		err2 := telegramity.InitGlobalClient("different_token", 987654321)
		if err2 != nil {
			t.Errorf("Second initialization should not fail: %v", err2)
		}

		client2 := telegramity.GetGlobalClient()

		if client1 != client2 {
			t.Errorf("Expected same client instance, got different instances")
		}

		telegramity.CloseGlobalClient()
	})

	t.Run("initialization_after_close", func(t *testing.T) {
		telegramity.CloseGlobalClient()

		err1 := telegramity.InitGlobalClient(botToken, 123456789)
		if err1 != nil {
			t.Fatalf("First initialization failed: %v", err1)
		}

		client1 := telegramity.GetGlobalClient()

		err := telegramity.CloseGlobalClient()
		if err != nil {
			t.Errorf("Failed to close client: %v", err)
		}

		err2 := telegramity.InitGlobalClient(botToken, 111222333)
		if err2 != nil {
			t.Errorf("Re-initialization failed: %v", err2)
		}

		client2 := telegramity.GetGlobalClient()

		if client1 == client2 {
			t.Errorf("Expected different client instances after re-initialization")
		}

		telegramity.CloseGlobalClient()
	})

	t.Run("configuration_options", func(t *testing.T) {
		telegramity.CloseGlobalClient()

		err := telegramity.InitGlobalClient(botToken, 123456789,
			telegramity.WithTimeout(5),
			telegramity.WithEnvironmentName("test_env"),
			telegramity.WithAppInfo("TestApp", "1.0.0"),
			telegramity.WithRateLimit(2),
		)
		if err != nil {
			t.Errorf("Failed to initialize with options: %v", err)
		}

		client := telegramity.GetGlobalClient()
		if client == nil {
			t.Errorf("Expected client but got nil")
		}

		telegramity.CloseGlobalClient()
	})
}

// TestErrorReporting tests actual error reporting functionality
func TestErrorReporting(t *testing.T) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatID == "" {
		t.Skip("Skipping integration test: TELEGRAM_BOT_TOKEN or TELEGRAM_CHAT_ID not set")
	}

	t.Run("report_error", func(t *testing.T) {
		telegramity.CloseGlobalClient()

		err := telegramity.InitGlobalClient(botToken, 123456789)
		if err != nil {
			t.Fatalf("Failed to initialize client: %v", err)
		}

		client := telegramity.GetGlobalClient()

		reportErr := client.ReportError(context.Background(), errors.New("test error"), "test_type")
		if reportErr != nil {
			t.Errorf("Failed to report error: %v", reportErr)
		}

		telegramity.CloseGlobalClient()
	})
}
