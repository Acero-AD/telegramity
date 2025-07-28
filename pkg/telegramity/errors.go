package telegramity

import "github.com/somosbytes/telegramity/internal/errors"

const (
	SeverityLow      errors.Severity = errors.SeverityLow      // Minor issues, informational
	SeverityMedium   errors.Severity = errors.SeverityMedium   // Moderate issues, warnings
	SeverityHigh     errors.Severity = errors.SeverityHigh     // Important issues, requires attention
	SeverityCritical errors.Severity = errors.SeverityCritical // Critical issues, immediate action required
)

const (
	ErrorTypeValidation = errors.ErrorTypeValidation
	ErrorTypeNetwork    = errors.ErrorTypeNetwork
	ErrorTypeDatabase   = errors.ErrorTypeDatabase
	ErrorTypeAuth       = errors.ErrorTypeAuth
	ErrorTypePayment    = errors.ErrorTypePayment
	ErrorTypeInternal   = errors.ErrorTypeInternal
	ErrorTypeRateLimit  = errors.ErrorTypeRateLimit
	ErrorTypeTimeout    = errors.ErrorTypeTimeout
)
