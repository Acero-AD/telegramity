# Telegramity

A Go SDK for observability that catches errors and sends them to a Telegram bot with rich stack traces and configurable options.

## 🚀 Quick Start

### 1. Set up your Telegram Bot

1. Create a bot with [@BotFather](https://t.me/botfather) on Telegram
2. Get your bot token
3. Start a chat with your bot
4. Get your chat ID by visiting: `https://api.telegram.org/bot<YOUR_BOT_TOKEN>/getUpdates`

### 2. Install the SDK

```bash
go get github.com/somosbytes/telegramity
```

### 3. Use the Global Singleton (Recommended)

```go
package main

import (
    "context"
    "errors"
    "log"
    "os"

    "github.com/somosbytes/telegramity/pkg/telegramity"
)

func main() {
    // Initialize once at app startup
    err := telegramity.InitGlobalClient(
        os.Getenv("TELEGRAM_BOT_TOKEN"),
        123456789, // Your chat ID
        telegramity.WithEnvironmentName("production"),
        telegramity.WithAppInfo("MyApp", "1.0.0"),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer telegramity.CloseGlobalClient()

    // Use anywhere in your code
    client := telegramity.GetGlobalClient()
    
    ctx := context.Background()
    err = client.ReportError(ctx, errors.New("database connection failed"), telegramity.ErrorTypeDatabase)
    if err != nil {
        log.Printf("Failed to report error: %v", err)
    }
}
```

## 📁 Project Structure

```
Telegramity/
├── pkg/telegramity/          # Public SDK interface
│   ├── config.go             # Configuration options
│   ├── errors.go             # Error types and constants
│   └── singleton.go          # Global singleton pattern
├── internal/                 # Internal implementation
│   ├── configs/              # Configuration management
│   ├── errors/               # Error handling internals
│   ├── formatters/           # Message formatting
│   ├── telegram/bot/         # Telegram bot client
│   └── telegramity/          # Core client factory
├── cmd/example/              # Example applications
├── tests/unit/               # Unit tests
└── config.env.example        # Environment template
```

## 🎯 Features

- **Global Singleton Pattern**: Easy to use anywhere in your app
- **Automatic Stack Traces**: Uses `github.com/pkg/errors` for readable traces
- **Rate Limiting**: Configurable message rate limits
- **Retry Logic**: Automatic retries with exponential backoff
- **Context Support**: Full context.Context integration
- **Thread Safe**: Safe for concurrent use
- **Configurable**: Rich configuration options
- **Comprehensive Testing**: 100% test coverage for core functionality

## 🔧 Configuration

```go
telegramity.InitGlobalClient(
    "bot_token",
    123456789,
    telegramity.WithEnvironmentName("production"),
    telegramity.WithAppInfo("MyApp", "1.0.0"),
    telegramity.WithTimeout(30*time.Second),
    telegramity.WithRateLimit(2), // 2 messages per second
    telegramity.WithMaxRetries(3),
)
```

### Configuration Options

| Option | Description | Default |
|--------|-------------|---------|
| `WithEnvironmentName()` | Set environment (production, staging, etc.) | `""` |
| `WithAppInfo()` | Set application name and version | `""` |
| `WithTimeout()` | Configure HTTP timeout | `30s` |
| `WithRateLimit()` | Set messages per second limit | `1` |
| `WithMaxRetries()` | Configure retry attempts | `3` |

## 📝 Error Types

Predefined error types for common scenarios:

```go
telegramity.ErrorTypeDatabase    // Database errors
telegramity.ErrorTypeNetwork     // Network/API errors
telegramity.ErrorTypeAuth        // Authentication errors
telegramity.ErrorTypeValidation  // Validation errors
telegramity.ErrorTypePayment     // Payment processing errors
telegramity.ErrorTypeInternal    // Internal server errors
telegramity.ErrorTypeRateLimit   // Rate limiting errors
telegramity.ErrorTypeTimeout     // Timeout errors
```

## 🧪 Testing

### Unit Tests
```bash
# Run all unit tests
go test ./tests/unit/ -v

# Run specific test categories
go test ./tests/unit/ -v -run "Test.*Singleton"
go test ./tests/unit/ -v -run "Test.*BotClient"
```

### Integration Tests
```bash
# Set up environment variables for integration tests
export TELEGRAM_BOT_TOKEN="your_bot_token"
export TELEGRAM_CHAT_ID="your_chat_id"

# Run integration tests
go test ./tests/unit/ -v -run "Test.*Integration"
```

### Test Coverage
- **Singleton Pattern**: Comprehensive tests for global client behavior
- **Bot Client**: Mock-based tests for Telegram API integration
- **Configuration**: Tests for all configuration options
- **Error Handling**: Tests for validation and error scenarios

## 🛡️ Security & Reliability

- **Input Validation**: Comprehensive validation of all inputs
- **Error Handling**: Graceful handling of network failures
- **Rate Limiting**: Prevents spam and respects API limits
- **Retry Logic**: Automatic retries for transient failures
- **Context Support**: Proper cancellation and timeout handling
- **Thread Safety**: Safe for concurrent use across your application

## 📚 Examples

### Basic Error Reporting
```go
client := telegramity.GetGlobalClient()
err = client.ReportError(ctx, errors.New("something went wrong"), telegramity.ErrorTypeDatabase)
```

### Error with Context
```go
ctx := context.Background()
context := map[string]interface{}{
    "user_id": 12345,
    "action": "database_query",
    "query": "SELECT * FROM users",
}

err = client.ReportErrorWithContext(ctx, errors.New("query failed"), telegramity.ErrorTypeDatabase, context)
```

### Custom Error Types
```go
err = client.ReportError(ctx, errors.New("payment failed"), "payment_processing")
```

## 🔄 Recent Updates

### v1.0.0 (Latest)
- ✅ **Clean Architecture**: Separated public API from internal implementation
- ✅ **Comprehensive Testing**: Added 100% test coverage for core functionality
- ✅ **Mock-Based Testing**: Removed dependency on real bot tokens for unit tests
- ✅ **Better Error Handling**: Improved validation and error scenarios
- ✅ **CI/CD Integration**: GitHub Actions for automated testing and security scanning
- ✅ **Documentation**: Enhanced README and release notes

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📚 Documentation

For more detailed documentation, visit: [https://github.com/somosbytes/telegramity](https://github.com/somosbytes/telegramity)

---

**Happy Error Reporting! 🚀** 