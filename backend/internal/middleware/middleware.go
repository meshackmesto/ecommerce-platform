package middleware

import (
	"ecommerce-platform/internal/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Logger returns a configured logger middleware for Fiber.
func Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path} - ${latency}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Africa/Nairobi",
	})
}

// CORS returns a configured CORS middleware.
func CORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Your Vue.js frontend address
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH, OPTIONS",
	})
}

// ErrorHandler is a custom global error handler for the application.
// It ensures that all errors are returned in our standard JSON format.
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Default to 500 Internal Server Error
	code := fiber.StatusInternalServerError
	message := "An unexpected error occurred. Please try again later."

	// Check if it's a fiber.Error type to get status code and message
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	// Log the error for debugging purposes
	log.Printf("Error: %v - Path: %s", err, c.Path())

	// Return a structured JSON error response
	return utils.ErrorResponse(c, code, message, err)
}
