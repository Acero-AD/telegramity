package unit

import (
	"testing"

	"github.com/somosbytes/telegramity/internal/configs"
	internalerrors "github.com/somosbytes/telegramity/internal/errors"
	"github.com/somosbytes/telegramity/pkg/telegramity"
)

func resetSingletonForTesting() {
	telegramity.CloseGlobalClient()
}

func TestInitGlobalClient(t *testing.T) {
	tests := []struct {
		name      string
		botToken  string
		chatID    int64
		options   []configs.ConfigOption
		expectErr bool
	}{
		{
			name:      "empty_bot_token",
			botToken:  "",
			chatID:    123456789,
			expectErr: true,
		},
		{
			name:      "zero_chat_id",
			botToken:  "test_token",
			chatID:    0,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetSingletonForTesting()

			err := telegramity.InitGlobalClient(tt.botToken, tt.chatID, tt.options...)

			if tt.expectErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}

func TestGetGlobalClient(t *testing.T) {
	t.Run("client_not_initialized", func(t *testing.T) {
		resetSingletonForTesting()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic but got none")
			}
		}()

		telegramity.GetGlobalClient()
	})
}

func TestCloseGlobalClient(t *testing.T) {
	t.Run("close_not_initialized", func(t *testing.T) {
		resetSingletonForTesting()

		err := telegramity.CloseGlobalClient()
		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}
	})
}

func TestSingletonBehavior(t *testing.T) {
	t.Run("validation_only", func(t *testing.T) {
		resetSingletonForTesting()

		err := telegramity.InitGlobalClient("", 123456789)
		if err == nil {
			t.Errorf("Expected error for empty bot token")
		}

		err = telegramity.InitGlobalClient("test_token", 0)
		if err == nil {
			t.Errorf("Expected error for zero chat ID")
		}
	})
}

func TestConfigurationOptions(t *testing.T) {
	t.Run("validation_with_options", func(t *testing.T) {
		resetSingletonForTesting()

		err := telegramity.InitGlobalClient("", 123456789,
			telegramity.WithTimeout(5),
		)
		if err == nil {
			t.Errorf("Expected error for empty bot token even with options")
		}
	})

	t.Run("validation_with_multiple_options", func(t *testing.T) {
		resetSingletonForTesting()

		err := telegramity.InitGlobalClient("", 0,
			telegramity.WithEnvironmentName("test_env"),
			telegramity.WithAppInfo("TestApp", "1.0.0"),
			telegramity.WithRateLimit(2),
		)
		if err == nil {
			t.Errorf("Expected error for invalid parameters even with multiple options")
		}
	})
}

func TestErrorTypes(t *testing.T) {
	errorTypes := []string{
		telegramity.ErrorTypeValidation,
		telegramity.ErrorTypeNetwork,
		telegramity.ErrorTypeDatabase,
		telegramity.ErrorTypeAuth,
		telegramity.ErrorTypePayment,
		telegramity.ErrorTypeInternal,
		telegramity.ErrorTypeRateLimit,
		telegramity.ErrorTypeTimeout,
	}

	for _, errorType := range errorTypes {
		if errorType == "" {
			t.Errorf("Error type should not be empty")
		}
	}
}

func TestSeverityLevels(t *testing.T) {
	severityLevels := []internalerrors.Severity{
		telegramity.SeverityLow,
		telegramity.SeverityMedium,
		telegramity.SeverityHigh,
		telegramity.SeverityCritical,
	}

	for _, severity := range severityLevels {
		if string(severity) == "" {
			t.Errorf("Severity level should not be empty")
		}
	}
}
