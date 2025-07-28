package telegramity

import (
	"time"

	"github.com/somosbytes/telegramity/internal/configs"
)

func WithTimeout(timeout time.Duration) configs.ConfigOption {
	return func(c *configs.Config) {
		c.Timeout = timeout
	}
}

func WithMaxRetries(maxRetries int) configs.ConfigOption {
	return func(c *configs.Config) {
		c.MaxRetries = maxRetries
	}
}

func WithRetryDelay(delay time.Duration) configs.ConfigOption {
	return func(c *configs.Config) {
		c.RetryDelay = delay
	}
}

func WithRateLimit(limit int) configs.ConfigOption {
	return func(c *configs.Config) {
		c.RateLimitPerSecond = limit
	}
}

func WithEnvironmentName(env string) configs.ConfigOption {
	return func(c *configs.Config) {
		c.Environment = env
	}
}

func WithAppInfo(name, version string) configs.ConfigOption {
	return func(c *configs.Config) {
		c.AppName = name
		c.AppVersion = version
	}
}

func WithMessageConfig(includeStackTrace, includeTimestamp bool, maxLength int) configs.ConfigOption {
	return func(c *configs.Config) {
		c.IncludeStackTrace = includeStackTrace
		c.IncludeTimestamp = includeTimestamp
		c.MaxMessageLength = maxLength
	}
}
