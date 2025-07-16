package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"` // "-" means don't include in JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Message represents a message in the system
type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"not null"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Pagination represents pagination metadata
type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}
