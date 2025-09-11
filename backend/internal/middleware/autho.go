package middleware

import (
	"ecommerce-platform/internal/models"
	"ecommerce-platform/internal/utils"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// JWTAuth must now accept the secret from main.go
func JWTAuth(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Missing authorization header", nil)
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid authorization header format", nil)
		}

		tokenString := parts[1]
		// Pass the secret to ValidateToken
		claims, err := utils.ValidateToken(tokenString, secret)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid or expired token", err)
		}

		// Store claims in context for later use
		c.Locals("user_id", claims.UserID)
		c.Locals("user_role", claims.Role)
		c.Locals("user_email", claims.Email)

		return c.Next()
	}
}

// GetUserIDFromContext retrieves the user ID from the Fiber context.
func GetUserIDFromContext(c *fiber.Ctx) (uuid.UUID, error) {
	id, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return uuid.Nil, fmt.Errorf("user_id not found in context or is of wrong type")
	}
	return id, nil
}

// GetRoleFromContext retrieves the user role from the Fiber context.
func GetRoleFromContext(c *fiber.Ctx) (models.UserRole, error) {
	role, ok := c.Locals("user_role").(models.UserRole)
	if !ok {
		return "", fmt.Errorf("user_role not found in context or is of wrong type")
	}
	return role, nil
}

// RequireRole is a middleware that checks if the user has the required role.
func RequireRole(role models.UserRole) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole, err := GetRoleFromContext(c)
		if err != nil || userRole != role {
			return utils.ErrorResponse(c, fiber.StatusForbidden, "Forbidden: Insufficient permissions.", nil)
		}
		return c.Next()
	}
}

// ... (rest of the file is the same)
