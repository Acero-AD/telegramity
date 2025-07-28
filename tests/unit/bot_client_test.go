package unit

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/somosbytes/telegramity/internal/telegram/bot"
)

// MockBotClient is a mock implementation of the BotClient interface
type MockBotClient struct {
	shouldFail  bool
	lastMessage string
	lastChatID  int64
}

// SendMessage is a mock implementation of the SendMessage method
func (m *MockBotClient) SendMessage(ctx context.Context, chatID int64, message string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if message == "" {
		return errors.New("message cannot be empty")
	}
	if chatID == 0 {
		return errors.New("chat ID cannot be zero")
	}

	m.lastMessage = message
	m.lastChatID = chatID

	if m.shouldFail {
		return errors.New("mock send message failed")
	}

	return nil
}

// TestConnection is a mock implementation of the TestConnection method
func (m *MockBotClient) TestConnection(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if m.shouldFail {
		return errors.New("mock connection test failed")
	}

	return nil
}

// TestNewBotClient tests the bot client constructor
func TestNewBotClient(t *testing.T) {
	tests := []struct {
		name        string
		token       string
		timeout     time.Duration
		expectError bool
	}{
		{
			name:        "empty token",
			token:       "",
			timeout:     10 * time.Second,
			expectError: true,
		},
		{
			name:        "zero timeout",
			token:       "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11",
			timeout:     0,
			expectError: true,
		},
		{
			name:        "negative timeout",
			token:       "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11",
			timeout:     -1 * time.Second,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := bot.NewBotClient(tt.token, tt.timeout)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Check that we got a client when we didn't expect an error
			if client == nil {
				t.Errorf("Expected client but got nil")
			}
		})
	}
}

func TestBotClientInterface(t *testing.T) {
	client := &MockBotClient{}

	var _ bot.BotClient = client

	t.Log("BotClient interface is implemented correctly")
}

func TestSendMessageValidation(t *testing.T) {
	client := &MockBotClient{}

	ctx := context.Background()

	tests := []struct {
		name        string
		chatID      int64
		message     string
		expectError bool
	}{
		{
			name:        "valid message",
			chatID:      123456789,
			message:     "Hello, world!",
			expectError: false,
		},
		{
			name:        "empty message",
			chatID:      123456789,
			message:     "",
			expectError: true,
		},
		{
			name:        "zero chat ID",
			chatID:      0,
			message:     "Hello, world!",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := client.SendMessage(ctx, tt.chatID, tt.message)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if client.lastMessage != tt.message {
				t.Errorf("Expected message %q, got %q", tt.message, client.lastMessage)
			}
			if client.lastChatID != tt.chatID {
				t.Errorf("Expected chat ID %d, got %d", tt.chatID, client.lastChatID)
			}
		})
	}
}

func TestContextCancellation(t *testing.T) {
	client := &MockBotClient{}

	// Create a context that's already cancelled
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	// Test SendMessage with cancelled context
	err := client.SendMessage(ctx, 123456789, "test message")
	if err != context.Canceled {
		t.Errorf("Expected context.Canceled error, got %v", err)
	}

	// Test TestConnection with cancelled context
	err = client.TestConnection(ctx)
	if err != context.Canceled {
		t.Errorf("Expected context.Canceled error, got %v", err)
	}
}

// TestMockFailure tests that our mock can simulate failures
func TestMockFailure(t *testing.T) {
	// Create a mock that should fail
	client := &MockBotClient{shouldFail: true}
	ctx := context.Background()

	// Test SendMessage failure
	err := client.SendMessage(ctx, 123456789, "test message")
	if err == nil {
		t.Error("Expected error but got none")
	}

	// Test TestConnection failure
	err = client.TestConnection(ctx)
	if err == nil {
		t.Error("Expected error but got none")
	}
}
