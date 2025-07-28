package telegramity

import (
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Severity represents the severity level of an error
type Severity string

const (
	SeverityLow      Severity = "low"      // Minor issues, informational
	SeverityMedium   Severity = "medium"   // Moderate issues, warnings
	SeverityHigh     Severity = "high"     // Important issues, requires attention
	SeverityCritical Severity = "critical" // Critical issues, immediate action required
)

// ErrorOption allows customizing error reporting behavior
type ErrorOption func(*ErrorReport)

// Common error types that users can use
const (
	ErrorTypeValidation = "validation"
	ErrorTypeNetwork    = "network"
	ErrorTypeDatabase   = "database"
	ErrorTypeAuth       = "auth"
	ErrorTypePayment    = "payment"
	ErrorTypeInternal   = "internal"
	ErrorTypeRateLimit  = "rate_limit"
	ErrorTypeTimeout    = "timeout"
)

// ErrorReport represents an error report to be sent
type ErrorReport struct {
	// Core Error Info (Required)
	Error      error  // The actual error
	ErrorType  string // User-defined error type
	StackTrace string // Auto-extracted unless user provides

	// Context (Optional)
	UserID      string // Affected user (optional)
	Environment string // dev, staging, prod (optional)
	AppName     string // Application name (optional)

	// Operational (Required)
	Severity  Severity  // low, medium, high, critical
	Timestamp time.Time // When it happened

	// Custom Data (Optional)
	Context map[string]interface{} // Additional metadata
}

// extractStackTrace extracts a readable stack trace from an error
// If the error has a stack trace (from pkg/errors), it uses that
// Otherwise, it gets the current stack trace
func extractStackTrace(err error) string {
	// First, try to get stack trace from pkg/errors
	if stackTracer, ok := err.(interface{ StackTrace() errors.StackTrace }); ok {
		// The error has a stack trace, format it nicely
		stack := stackTracer.StackTrace()
		return fmt.Sprintf("%+v", stack)
	}

	// Fallback: get current stack trace
	stack := debug.Stack()
	lines := strings.Split(string(stack), "\n")

	// Skip the first few lines (they're usually just debug.Stack() calls)
	// Start from where the actual error occurred
	startIndex := 0
	for i, line := range lines {
		if strings.Contains(line, "telegramity") && !strings.Contains(line, "extractStackTrace") {
			startIndex = i
			break
		}
	}

	if startIndex > 0 && startIndex < len(lines) {
		return strings.Join(lines[startIndex:], "\n")
	}

	// If we can't find a good starting point, return the full stack
	return string(stack)
}

// createErrorReport creates an ErrorReport with automatic stack trace extraction
func createErrorReport(err error, errorType string, opts ...ErrorOption) *ErrorReport {
	report := &ErrorReport{
		Error:      err,
		ErrorType:  errorType,
		StackTrace: extractStackTrace(err), // Auto-extract stack trace
		Severity:   SeverityMedium,         // Default severity
		Timestamp:  time.Now(),
		Context:    make(map[string]interface{}),
	}

	// Apply options
	for _, opt := range opts {
		opt(report)
	}

	return report
}
