package telegramity

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/somosbytes/telegramity/internal/telegram/bot"
)

// Client is the main interface for the Telegramity SDK
type Client interface {
	// ReportError sends an error to the configured Telegram bot
	ReportError(ctx context.Context, err error, errorType string, opts ...ErrorOption) error

	// ReportErrorWithContext sends an error with additional context
	ReportErrorWithContext(ctx context.Context, err error, errorType string, context map[string]interface{}, opts ...ErrorOption) error

	// Close gracefully shuts down the client
	Close() error
}

// client implements the telegramity.Client interface
type client struct {
	config      *Config
	bot         bot.BotClient
	rateLimiter *time.Ticker
	mu          sync.RWMutex
	closed      bool
}

// ReportError implements telegramity.Client.ReportError
func (c *client) ReportError(ctx context.Context, err error, errorType string, opts ...ErrorOption) error {
	return c.ReportErrorWithContext(ctx, err, errorType, nil, opts...)
}

// ReportErrorWithContext implements telegramity.Client.ReportErrorWithContext
func (c *client) ReportErrorWithContext(ctx context.Context, err error, errorType string, context map[string]interface{}, opts ...ErrorOption) error {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return fmt.Errorf("client is closed")
	}
	c.mu.RUnlock()

	if err == nil {
		return fmt.Errorf("error cannot be nil")
	}

	// Create error report using our new function
	report := createErrorReport(err, errorType, opts...)

	// Add context if provided
	if context != nil {
		report.Context = context
	}

	// Set environment and app info from config if not already set
	if report.Environment == "" && c.config.Environment != "" {
		report.Environment = c.config.Environment
	}
	if report.AppName == "" && c.config.AppName != "" {
		report.AppName = c.config.AppName
	}

	// Wait for rate limiter
	select {
	case <-c.rateLimiter.C:
	case <-ctx.Done():
		return ctx.Err()
	}

	// Format and send message
	message, err := c.formatErrorReport(report)
	if err != nil {
		return fmt.Errorf("failed to format error report: %w", err)
	}

	// Send with retries
	for attempt := 0; attempt <= c.config.MaxRetries; attempt++ {
		err = c.bot.SendMessage(ctx, c.config.ChatID, message)
		if err == nil {
			break
		}

		if attempt == c.config.MaxRetries {
			return fmt.Errorf("failed to send error report after %d attempts: %w", c.config.MaxRetries+1, err)
		}

		// Wait before retry
		select {
		case <-time.After(c.config.RetryDelay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}

// Close implements telegramity.Client.Close
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

// formatErrorReport formats an error report into a Telegram message
func (c *client) formatErrorReport(report *ErrorReport) (string, error) {
	message := "🚨 <b>Error Report</b>\n\n"

	// Add timestamp if configured
	if c.config.IncludeTimestamp {
		message += fmt.Sprintf("⏰ <b>Time:</b> %s\n", report.Timestamp.Format("2006-01-02 15:04:05"))
	}

	// Add error type
	message += fmt.Sprintf("🔍 <b>Type:</b> %s\n", report.ErrorType)

	// Add error message
	message += fmt.Sprintf("❌ <b>Error:</b> %s\n", report.Error.Error())

	// Add severity
	if report.Severity != "" {
		message += fmt.Sprintf("⚠️ <b>Severity:</b> %s\n", report.Severity)
	}

	// Add user info if available
	if report.UserID != "" {
		message += fmt.Sprintf("👤 <b>User:</b> %s\n", report.UserID)
	}

	// Add environment if available
	if report.Environment != "" {
		message += fmt.Sprintf("🌍 <b>Environment:</b> %s\n", report.Environment)
	}

	// Add app info if available
	if report.AppName != "" {
		message += fmt.Sprintf("📱 <b>App:</b> %s\n", report.AppName)
	}

	// Add context if available
	if len(report.Context) > 0 {
		message += fmt.Sprintf("📋 <b>Context:</b> %+v\n", report.Context)
	}

	// Add stack trace if configured
	if c.config.IncludeStackTrace && report.StackTrace != "" {
		// Format stack trace for better readability
		stackTrace := c.formatStackTrace(report.StackTrace)
		message += fmt.Sprintf("\n🔍 <b>Stack Trace:</b>\n<pre><code>%s</code></pre>", stackTrace)
	}

	return message, nil
}

// formatStackTrace formats a stack trace for better readability in Telegram
func (c *client) formatStackTrace(stackTrace string) string {
	// Split into lines for processing
	lines := strings.Split(stackTrace, "\n")

	// Filter and format lines
	var formattedLines []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Highlight important parts (file paths, line numbers)
		if strings.Contains(line, ".go:") {
			// This is a file:line reference
			formattedLines = append(formattedLines, fmt.Sprintf("📍 %s", line))
		} else if strings.Contains(line, "github.com/") || strings.Contains(line, "main.") {
			// This is a function call
			formattedLines = append(formattedLines, fmt.Sprintf("🔗 %s", line))
		} else {
			// Regular line
			formattedLines = append(formattedLines, line)
		}
	}

	// Limit the number of lines to avoid huge messages
	maxLines := 20
	if len(formattedLines) > maxLines {
		formattedLines = formattedLines[:maxLines]
		formattedLines = append(formattedLines, "...")
	}

	return strings.Join(formattedLines, "\n")
}
