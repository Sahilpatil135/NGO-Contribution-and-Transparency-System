package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	PasswordHash *string   `json:"-" db:"password_hash"` // Hidden from JSON
	Provider     string    `json:"provider" db:"provider"`
	ProviderID   *string   `json:"provider_id" db:"provider_id"`
	AvatarURL    *string   `json:"avatar_url" db:"avatar_url"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	IsVerified   bool      `json:"is_verified" db:"is_verified"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	Role         string    `json:"role" db:"role"`
}

// CreateUserRequest represents the request payload for creating a user
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// LoginRequest represents the request payload for user login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// AuthResponse represents the response after successful authentication
type AuthResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

// UserResponse represents a user response without sensitive data
type UserResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Provider   string    `json:"provider"`
	AvatarURL  *string   `json:"avatar_url"`
	IsActive   bool      `json:"is_active"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Role       string    `json:"role" db:"role"`
}

// ToUserResponse converts a User to UserResponse
func (u *User) ToUserResponse() UserResponse {
	return UserResponse{
		ID:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		Provider:   u.Provider,
		AvatarURL:  u.AvatarURL,
		IsActive:   u.IsActive,
		IsVerified: u.IsVerified,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
		Role:       u.Role,
	}
}
