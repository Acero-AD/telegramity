# Telegramity

A Go application built with the Fiber framework.

## Features

- Fast HTTP server with Fiber framework
- Environment-based configuration
- Structured project layout
- CORS and logging middleware
- Health check endpoints
- Modular architecture

## Project Structure

```
telegramity/
├── main.go                    # Application entry point
├── go.mod                     # Go module file
├── go.sum                     # Dependency checksums
├── config/
│   └── config.go              # Configuration management
├── handlers/
│   ├── handlers.go            # Common handlers (404, home)
│   ├── health_handlers.go     # Health check endpoints
│   ├── user_handlers.go       # User-related endpoints
│   └── message_handlers.go    # Message-related endpoints
├── services/
│   └── user_service.go        # Business logic layer
├── models/
│   └── models.go              # Data models
├── routes/
│   └── routes.go              # Route definitions
├── middleware/
│   └── middleware.go          # Custom middleware
├── .env                       # Environment variables (create this)
├── .env.example               # Example environment file
└── README.md                 # This file
```

## Setup Instructions

### 1. Prerequisites

- Go 1.19 or higher
- Git

### 2. Clone and Setup

```bash
git clone <your-repo-url>
cd telegramity
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Environment Configuration

Create a `.env` file in the root directory:

```env
# Application Configuration
PORT=3000
APP_NAME=Telegramity
APP_VERSION=1.0.0
APP_ENV=development

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=telegramity
DB_USER=postgres
DB_PASSWORD=

# JWT Configuration (for future use)
JWT_SECRET=your-secret-key-here
JWT_EXPIRY=24h

# Logging
LOG_LEVEL=info
```

### 5. Run the Application

```bash
go run main.go
```

The server will start on `http://localhost:3000`

## API Endpoints

### Health Check
- `GET /` - Welcome message
- `GET /api/v1/health` - Basic health check
- `GET /api/v1/health/detailed` - Detailed health check

### Users
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create new user
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Messages
- `GET /api/v1/messages` - Get all messages
- `GET /api/v1/messages/:id` - Get message by ID
- `POST /api/v1/messages` - Create new message
- `GET /api/v1/messages/user/:user_id` - Get messages by user

## Development

### Adding New Routes

1. Add handler functions in `handlers/handlers.go`
2. Register routes in `routes/routes.go`
3. Import and use the routes in `main.go`

### Adding Middleware

1. Create middleware functions in `middleware/middleware.go`
2. Apply them in `main.go` using `app.Use()`

### Database Integration

The project includes GORM tags in the models for future database integration. To add database support:

1. Install database driver: `go get gorm.io/driver/postgres`
2. Configure database connection in `config/config.go`
3. Initialize database in `main.go`

## Building for Production

```bash
go build -o telegramity main.go
```

## Testing

```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

[Add your license here] 