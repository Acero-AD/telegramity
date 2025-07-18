# Telegramity - Go SDK for Error Reporting to Telegram

Telegramity is a Go SDK that allows you to easily report errors from your applications directly to your Telegram bot. It's designed to be simple to use while providing powerful customization options.

## Project Structure

```
Telegramity/
├── pkg/telegramity/          # Public SDK interface (what users import)
│   ├── client.go            # Main client interface and types
│   ├── config.go            # Configuration types and options
│   └── telegramity.go       # Main entry point (NewClient function)
├── internal/                # Internal implementation details
│   ├── errors/              # Custom error types
│   │   └── errors.go        # SDK error definitions
│   ├── telegram/            # Telegram integration
│   │   └── bot/             # Bot client implementation
│   │       └── client.go    # Telegram bot client
│   ├── formatters/          # Message formatting
│   │   └── formatter.go     # Error message formatter
│   └── [old API files]      # Moved from root (can be removed)
├── cmd/example/             # Example applications
│   └── main.go              # Basic usage example
├── docs/                    # Documentation
├── examples/                # Usage examples
├── tests/                   # Test files
│   ├── unit/                # Unit tests
│   └── integration/         # Integration tests
├── go.mod                   # Go module file
├── go.sum                   # Dependency checksums
└── README.md               # This file
```

## Structure Explanation

### `pkg/telegramity/` - Public API
This is what users will import and use. Contains:
- **client.go**: Main `Client` interface and error types
- **config.go**: Configuration structs and option functions
- **telegramity.go**: Main entry point with `NewClient()` function

### `internal/` - Implementation Details
Internal packages that users don't need to see:
- **errors/**: Custom error types for the SDK
- **telegram/bot/**: Telegram Bot API integration
- **formatters/**: Message formatting logic

### `cmd/example/` - Example Applications
Demonstrates how to use the SDK with real examples.

### `tests/` - Testing
Organized test structure for unit and integration tests.

## Key Design Principles

1. **Interface-Based Design**: Clean separation between interface and implementation
2. **Functional Options Pattern**: For flexible configuration
3. **Error Handling**: Structured error types with proper categorization
4. **Rate Limiting**: Built-in rate limiting for Telegram API compliance
5. **Thread Safety**: Safe for concurrent use

## Next Steps

1. Implement the interfaces and types in `pkg/telegramity/`
2. Create the internal implementation in `internal/`
3. Add comprehensive tests
4. Create usage examples
5. Add documentation

## Development Guidelines

- Keep the public API in `pkg/` simple and stable
- Use `internal/` for implementation details
- Follow Go conventions and idioms
- Write tests for all public APIs
- Document all exported functions and types 