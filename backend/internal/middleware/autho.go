package middleware

import (
    "strings"

    "ecommerce-platform/internal/models"
    "ecommerce-platform/internal/utils"
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

// AuthMiddleware validates JWT token and sets user context
func AuthMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Get token from Authorization header
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Authorization header required", nil)
        }

        // Check Bearer format
        tokenParts := strings.Split(authHeader, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid authorization format", nil)
        }

        token := tokenParts[1]

        // Validate token
        claims, err := utils.ValidateToken(token)
        if err != nil {
            return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid or expired token", err)
        }

        // Set user context
        c.Locals("user_id", claims.UserID)
        c.Locals("user_email", claims.Email)
        c.Locals("user_role", claims.Role)

        return c.Next()
    }
}

// AdminMiddleware ensures the user has admin role
func AdminMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        userRole, ok := c.Locals("user_role").(string)
        if !ok || userRole != string(models.RoleAdmin) {
            return utils.ErrorResponse(c, fiber.StatusForbidden, "Admin access required", nil)
        }

        return c.Next()
    }
}

// GetUserIDFromContext extracts user ID from context
func GetUserIDFromContext(c *fiber.Ctx) (uuid.UUID, error) {
    userID, ok := c.Locals("user_id").(uuid.UUID)
    if !ok {
        return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
    }
    return userID, nil
}

// GetUserRoleFromContext extracts user role from context
func GetUserRoleFromContext(c *fiber.Ctx) string {
    userRole, _ := c.Locals("user_role").(string)
    return userRole
}