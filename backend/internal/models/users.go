package models

import (
    "time"
    "github.com/google/uuid"
)

type UserRole string

const (
    RoleAdmin    UserRole = "admin"
    RoleCustomer UserRole = "customer"
)

type User struct {
    ID           uuid.UUID `json:"id" db:"id"`
    Email        string    `json:"email" db:"email"`
    PasswordHash string    `json:"-" db:"password_hash"` // "-" means exclude from JSON
    FirstName    string    `json:"first_name" db:"first_name"`
    LastName     string    `json:"last_name" db:"last_name"`
    Phone        *string   `json:"phone" db:"phone"`
    Role         UserRole  `json:"role" db:"role"`
    IsActive     bool      `json:"is_active" db:"is_active"`
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
    UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// LoginRequest represents the login request payload
type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}

// RegisterRequest represents the registration request payload
type RegisterRequest struct {
    Email     string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required,min=6"`
    FirstName string `json:"first_name" validate:"required,min=2"`
    LastName  string `json:"last_name" validate:"required,min=2"`
    Phone     string `json:"phone"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
    User  User   `json:"user"`
    Token string `json:"token"`
}

// GetFullName returns the user's full name
func (u *User) GetFullName() string {
    return u.FirstName + " " + u.LastName
}

// IsAdmin checks if the user is an admin
func (u *User) IsAdmin() bool {
    return u.Role == RoleAdmin
}