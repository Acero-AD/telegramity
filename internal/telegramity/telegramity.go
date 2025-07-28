package telegramity

import (
	"fmt"
	"time"

	"github.com/somosbytes/telegramity/internal/configs"
	"github.com/somosbytes/telegramity/internal/telegram/bot"
)

func NewClient(botToken string, chatID int64, options ...configs.ConfigOption) (bot.Client, error) {
	// Create default configuration
	config := configs.DefaultConfig()
	config.BotToken = botToken
	config.ChatID = chatID

	// Apply configuration options
	for _, option := range options {
		option(&config)
	}

	// Validate required fields
	if config.BotToken == "" {
		return nil, fmt.Errorf("bot token is required")
	}
	if config.ChatID == 0 {
		return nil, fmt.Errorf("chat ID is required")
	}

	// Create the internal client implementation
	return newClient(&config)
}

// newClient creates the internal client implementation
func newClient(config *configs.Config) (bot.Client, error) {
	// Create the bot client
	botClient, err := bot.NewBotClient(config.BotToken, config.Timeout)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot client: %w", err)
	}

	// Create rate limiter
	rateLimiter := time.NewTicker(time.Duration(1000/config.RateLimitPerSecond) * time.Millisecond)

	// Create the main client implementation
	client := bot.NewClient(config, botClient, rateLimiter)

	return client, nil
}
