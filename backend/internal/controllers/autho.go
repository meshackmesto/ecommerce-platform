package controllers

import (
    "ecommerce-platform/internal/models"
    "ecommerce-platform/internal/services"
    "ecommerce-platform/internal/utils"
    "ecommerce-platform/internal/middleware"

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

// Me returns the current authenticated user
func (ac *AuthController) Me(c *fiber.Ctx) error {
    userID, err := middleware.GetUserIDFromContext(c)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusUnauthorized, "User not authenticated", err)
    }

    user, err := ac.authService.GetUserByID(userID)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found", err)
    }

    return utils.SuccessResponse(c, "User retrieved successfully", user)
}