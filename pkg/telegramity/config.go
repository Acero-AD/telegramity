package telegramity

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

// Configuration option functions

// WithTimeout sets the timeout for API calls
func WithTimeout(timeout time.Duration) ConfigOption {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithMaxRetries sets the maximum number of retries
func WithMaxRetries(maxRetries int) ConfigOption {
	return func(c *Config) {
		c.MaxRetries = maxRetries
	}
}

// WithRetryDelay sets the delay between retries
func WithRetryDelay(delay time.Duration) ConfigOption {
	return func(c *Config) {
		c.RetryDelay = delay
	}
}

// WithRateLimit sets the rate limit per second
func WithRateLimit(limit int) ConfigOption {
	return func(c *Config) {
		c.RateLimitPerSecond = limit
	}
}

// WithEnvironmentName sets the environment name
func WithEnvironmentName(env string) ConfigOption {
	return func(c *Config) {
		c.Environment = env
	}
}

// WithAppInfo sets the application name and version
func WithAppInfo(name, version string) ConfigOption {
	return func(c *Config) {
		c.AppName = name
		c.AppVersion = version
	}
}

// WithMessageConfig configures message formatting options
func WithMessageConfig(includeStackTrace, includeTimestamp bool, maxLength int) ConfigOption {
	return func(c *Config) {
		c.IncludeStackTrace = includeStackTrace
		c.IncludeTimestamp = includeTimestamp
		c.MaxMessageLength = maxLength
	}
}
