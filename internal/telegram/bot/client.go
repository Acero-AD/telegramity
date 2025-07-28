package bot

import (
	"context"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// BotClient defines the interface for Telegram bot operations
type BotClient interface {
	SendMessage(ctx context.Context, chatID int64, message string) error

	TestConnection(ctx context.Context) error
}

type botClient struct {
	bot     *tgbotapi.BotAPI
	timeout time.Duration
}

// NewBotClient creates a new Telegram bot client
func NewBotClient(token string, timeout time.Duration) (BotClient, error) {
	if token == "" {
		return nil, fmt.Errorf("bot token cannot be empty")
	}
	if timeout <= 0 {
		return nil, fmt.Errorf("timeout must be positive")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot API client: %w", err)
	}

	return &botClient{
		bot:     bot,
		timeout: timeout,
	}, nil
}

func (c *botClient) SendMessage(ctx context.Context, chatID int64, message string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// Validate inputs
	if message == "" {
		return fmt.Errorf("message cannot be empty")
	}
	if chatID == 0 {
		return fmt.Errorf("chat ID cannot be zero")
	}

	msg := tgbotapi.NewMessage(chatID, message)

	msg.ParseMode = "HTML"

	_, err := c.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (c *botClient) TestConnection(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := c.bot.GetMe()
	if err != nil {
		return fmt.Errorf("failed to test bot connection: %w", err)
	}

	return nil
}
