package controllers

import (
	"ecommerce-platform/internal/middleware"
	"ecommerce-platform/internal/models"
	"ecommerce-platform/internal/services"
	"ecommerce-platform/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Login handles user login
func (ac *AuthController) Login(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	authResp, err := ac.authService.Login(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return utils.SuccessResponse(c, "Login successful", authResp)
}

// Register handles user registration
func (ac *AuthController) Register(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	authResp, err := ac.authService.Register(&req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return utils.SuccessResponse(c, "Registration successful", authResp)
}

// GetProfile returns the current authenticated user's profile
func (ac *AuthController) GetProfile(c *fiber.Ctx) error {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "User not authenticated", err)
	}

	user, err := ac.authService.GetUserByID(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found", err)
	}
    
    // We should not send the password hash to the client
    user.PasswordHash = ""

	return utils.SuccessResponse(c, "User profile retrieved successfully", user)
}

// UpdateProfile is a placeholder for the update logic
func (ac *AuthController) UpdateProfile(c *fiber.Ctx) error {
    // You would add the logic to update a user's profile here.
    // 1. Get user ID from context
    // 2. Parse the request body for new profile data
    // 3. Call a service method to update the user in the database
	return utils.SuccessResponse(c, "Profile update endpoint hit (logic not implemented)", nil)
}


// RefreshToken is a placeholder for refresh token logic
func (ac *AuthController) RefreshToken(c *fiber.Ctx) error {
	// Logic to handle token refreshing would go here
	return utils.SuccessResponse(c, "Refresh token endpoint hit (logic not implemented)", nil)
}