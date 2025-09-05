package utils

import (
    "github.com/gofiber/fiber/v2"
)

type Response struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

// SuccessResponse returns a success response
func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
    return c.JSON(Response{
        Success: true,
        Message: message,
        Data:    data,
    })
}

// ErrorResponse returns an error response
func ErrorResponse(c *fiber.Ctx, statusCode int, message string, err error) error {
    response := Response{
        Success: false,
        Message: message,
    }
    
    if err != nil {
        response.Error = err.Error()
    }
    
    return c.Status(statusCode).JSON(response)
}

// ValidationErrorResponse returns a validation error response
func ValidationErrorResponse(c *fiber.Ctx, message string) error {
    return c.Status(fiber.StatusBadRequest).JSON(Response{
        Success: false,
        Message: message,
    })
}