package formatters

import (
	"fmt"
	"strings"

	"github.com/somosbytes/telegramity/internal/configs"
	"github.com/somosbytes/telegramity/internal/errors"
)

type ErrorFormatter struct {
	config *configs.Config
}

func NewErrorFormatter(config *configs.Config) *ErrorFormatter {
	return &ErrorFormatter{
		config: config,
	}
}

func (f *ErrorFormatter) FormatErrorReport(report *errors.ErrorReport) (string, error) {
	message := "🚨 <b>Error Report</b>\n\n"

	if f.config.IncludeTimestamp {
		message += fmt.Sprintf("⏰ <b>Time:</b> %s\n", report.Timestamp.Format("2006-01-02 15:04:05"))
	}

	message += fmt.Sprintf("🔍 <b>Type:</b> %s\n", report.ErrorType)

	message += fmt.Sprintf("❌ <b>Error:</b> %s\n", report.Error.Error())

	if report.Severity != "" {
		message += fmt.Sprintf("⚠️ <b>Severity:</b> %s\n", report.Severity)
	}

	if report.UserID != "" {
		message += fmt.Sprintf("👤 <b>User:</b> %s\n", report.UserID)
	}

	if report.Environment != "" {
		message += fmt.Sprintf("🌍 <b>Environment:</b> %s\n", report.Environment)
	}

	if report.AppName != "" {
		message += fmt.Sprintf("📱 <b>App:</b> %s\n", report.AppName)
	}

	if len(report.Context) > 0 {
		message += fmt.Sprintf("📋 <b>Context:</b> %+v\n", report.Context)
	}

	if f.config.IncludeStackTrace && report.StackTrace != "" {
		stackTrace := f.formatStackTrace(report.StackTrace)
		message += fmt.Sprintf("\n🔍 <b>Stack Trace:</b>\n<pre><code>%s</code></pre>", stackTrace)
	}

	return message, nil
}

func (f *ErrorFormatter) formatStackTrace(stackTrace string) string {
	lines := strings.Split(stackTrace, "\n")

	var formattedLines []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.Contains(line, ".go:") {
			formattedLines = append(formattedLines, fmt.Sprintf("📍 %s", line))
		} else if strings.Contains(line, "github.com/") || strings.Contains(line, "main.") {
			formattedLines = append(formattedLines, fmt.Sprintf("🔗 %s", line))
		} else {
			formattedLines = append(formattedLines, line)
		}
	}

	maxLines := 20
	if len(formattedLines) > maxLines {
		formattedLines = formattedLines[:maxLines]
		formattedLines = append(formattedLines, "...")
	}

	return strings.Join(formattedLines, "\n")
}
