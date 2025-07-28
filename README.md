# Telegramity

A Go SDK for observability that catches errors and sends them to a Telegram bot.

## 🚀 Quick Start

### 1. Set up your Telegram Bot

1. Create a bot with [@BotFather](https://t.me/botfather) on Telegram
2. Get your bot token
3. Start a chat with your bot
4. Get your chat ID by visiting: `https://api.telegram.org/bot<YOUR_BOT_TOKEN>/getUpdates`

### 2. Install the SDK

```bash
# Option 1: From GitHub (when published)
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

### 4. Alternative: Manual Client Management

```go
// Create a client manually (for advanced use cases)
client, err := telegramity.NewClient("bot_token", 123456789)
if err != nil {
    log.Fatal(err)
}
defer client.Close()

err = client.ReportError(ctx, errors.New("something went wrong"), "database")
```

## 📁 Project Structure

```
Telegramity/
├── pkg/telegramity/          # Public SDK interface
│   ├── telegramity.go        # Main entry point
│   ├── client_impl.go        # Client implementation
│   ├── config.go             # Configuration types
│   ├── errors.go             # Error handling types
│   └── singleton.go          # Global singleton pattern
├── internal/                 # Internal implementation
│   └── telegram/bot/         # Telegram bot client
├── cmd/example/              # Example applications
├── tests/                    # Test files
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

```bash
# Run unit tests
go test ./tests/unit/ -v

# Test the singleton pattern
go run cmd/example/singleton_example.go
```

## 📄 License

MIT License

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📚 Documentation

For more detailed documentation, visit: [https://github.com/somosbytes/telegramity](https://github.com/somosbytes/telegramity) 