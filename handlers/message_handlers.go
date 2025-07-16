package handlers

import (
	"strconv"

	"telegramity/models"

	"github.com/gofiber/fiber/v2"
)

// MessageHandler handles message-related HTTP requests
type MessageHandler struct {
	// In a real application, you would inject services here
	// messageService *services.MessageService
	// db             *gorm.DB
}

// NewMessageHandler creates a new MessageHandler instance
func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

// GetMessages returns all messages
func (h *MessageHandler) GetMessages(c *fiber.Ctx) error {
	// In a real application, you would fetch from database
	messages := []models.Message{
		{
			ID:      1,
			Content: "Hello, world!",
			UserID:  1,
			User: models.User{
				ID:       1,
				Username: "john_doe",
				Email:    "john@example.com",
			},
		},
		{
			ID:      2,
			Content: "How are you?",
			UserID:  2,
			User: models.User{
				ID:       2,
				Username: "jane_smith",
				Email:    "jane@example.com",
			},
		},
	}

	return c.JSON(models.Response{
		Success: true,
		Data:    messages,
		Message: "Messages retrieved successfully",
	})
}

// GetMessage returns a specific message by ID
func (h *MessageHandler) GetMessage(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Error:   "Invalid message ID",
		})
	}

	// In a real application, you would fetch from database
	message := models.Message{
		ID:      uint(id),
		Content: "Hello, world!",
		UserID:  1,
		User: models.User{
			ID:       1,
			Username: "john_doe",
			Email:    "john@example.com",
		},
	}

	return c.JSON(models.Response{
		Success: true,
		Data:    message,
		Message: "Message retrieved successfully",
	})
}

// CreateMessage creates a new message
func (h *MessageHandler) CreateMessage(c *fiber.Ctx) error {
	var message models.Message

	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	// In a real application, you would validate and save to database
	message.ID = 3 // Simulate generated ID

	return c.Status(fiber.StatusCreated).JSON(models.Response{
		Success: true,
		Data:    message,
		Message: "Message created successfully",
	})
}

// GetMessagesByUser returns all messages for a specific user
func (h *MessageHandler) GetMessagesByUser(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("user_id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Error:   "Invalid user ID",
		})
	}

	// In a real application, you would fetch from database
	messages := []models.Message{
		{
			ID:      1,
			Content: "Hello, world!",
			UserID:  uint(userID),
		},
	}

	return c.JSON(models.Response{
		Success: true,
		Data:    messages,
		Message: "User messages retrieved successfully",
	})
}
