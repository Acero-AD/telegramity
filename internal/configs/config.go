package configs

import (
	"time"
)

// Config holds the configuration for the Telegramity client
type Config struct {
	// Telegram Bot Configuration
	BotToken string // Your bot token from @BotFather
	ChatID   int64  // Chat ID where to send error messages

	// Client Configuration
	Timeout    time.Duration // How long to wait for API calls
	MaxRetries int           // Maximum number of retry attempts
	RetryDelay time.Duration // Delay between retries

	// Rate Limiting
	RateLimitPerSecond int // Messages per second limit

	// Message Configuration
	MaxMessageLength  int  // Maximum message length (Telegram limit: 4096)
	IncludeStackTrace bool // Whether to include stack traces
	IncludeTimestamp  bool // Whether to include timestamps

	// Environment
	Environment string // Environment name (dev, staging, prod)
	AppName     string // Application name
	AppVersion  string // Application version
}

// DefaultConfig returns a default configuration
func DefaultConfig() Config {
	return Config{
		Timeout:            30 * time.Second,
		MaxRetries:         3,
		RetryDelay:         1 * time.Second,
		RateLimitPerSecond: 1,    // 1 message per second by default
		MaxMessageLength:   4096, // Telegram message limit
		IncludeStackTrace:  true,
		IncludeTimestamp:   true,
		Environment:        "development",
		AppName:            "unknown",
		AppVersion:         "1.0.0",
	}
}

// ConfigOption allows customizing the client configuration
type ConfigOption func(*Config)
