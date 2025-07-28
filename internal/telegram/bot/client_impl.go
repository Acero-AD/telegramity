package bot

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/somosbytes/telegramity/internal/configs"
	"github.com/somosbytes/telegramity/internal/errors"
	"github.com/somosbytes/telegramity/internal/formatters"
)

type Client interface {
	ReportError(ctx context.Context, err error, errorType string, opts ...errors.ErrorOption) error
	ReportErrorWithContext(ctx context.Context, err error, errorType string, context map[string]interface{}, opts ...errors.ErrorOption) error
	Close() error
}

type client struct {
	config      *configs.Config
	bot         BotClient
	rateLimiter *time.Ticker
	mu          sync.RWMutex
	closed      bool
}

func NewClient(config *configs.Config, botClient BotClient, rateLimiter *time.Ticker) Client {
	return &client{
		config:      config,
		bot:         botClient,
		rateLimiter: rateLimiter,
	}
}

func (c *client) ReportError(ctx context.Context, err error, errorType string, opts ...errors.ErrorOption) error {
	return c.ReportErrorWithContext(ctx, err, errorType, nil, opts...)
}

func (c *client) ReportErrorWithContext(ctx context.Context, err error, errorType string, context map[string]interface{}, opts ...errors.ErrorOption) error {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return fmt.Errorf("client is closed")
	}
	c.mu.RUnlock()

	if err == nil {
		return fmt.Errorf("error cannot be nil")
	}

	report := errors.CreateErrorReport(err, errorType, opts...)

	if context != nil {
		report.Context = context
	}

	if report.Environment == "" && c.config.Environment != "" {
		report.Environment = c.config.Environment
	}
	if report.AppName == "" && c.config.AppName != "" {
		report.AppName = c.config.AppName
	}

	select {
	case <-c.rateLimiter.C:
	case <-ctx.Done():
		return ctx.Err()
	}

	formatter := formatters.NewErrorFormatter(c.config)
	message, err := formatter.FormatErrorReport(report)
	if err != nil {
		return fmt.Errorf("failed to format error report: %w", err)
	}

	for attempt := 0; attempt <= c.config.MaxRetries; attempt++ {
		err = c.bot.SendMessage(ctx, c.config.ChatID, message)
		if err == nil {
			break
		}

		if attempt == c.config.MaxRetries {
			return fmt.Errorf("failed to send error report after %d attempts: %w", c.config.MaxRetries+1, err)
		}

		select {
		case <-time.After(c.config.RetryDelay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}

func (c *client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil
	}

	c.closed = true
	if c.rateLimiter != nil {
		c.rateLimiter.Stop()
	}
	return nil
}
