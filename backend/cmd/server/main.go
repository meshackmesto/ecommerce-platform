package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"ecommerce-platform/internal/config"
	"ecommerce-platform/internal/controllers"
	"ecommerce-platform/internal/database"
	"ecommerce-platform/internal/middleware"
	"ecommerce-platform/internal/repositories"
	"ecommerce-platform/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Run migrations
	 if err := database.RunMigrations(db); err != nil { // This now correctly calls the package-level function
        log.Fatalf("Failed to run migrations: %v", err)
    }
	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	productService := services.NewProductService(productRepo)
	orderService := services.NewOrderService(orderRepo, productRepo)
	paymentService := services.NewPaymentService(transactionRepo, orderRepo)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService)
	orderController := controllers.NewOrderController(orderService)
	paymentController := controllers.NewPaymentController(paymentService)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	// Global middleware
	app.Use(recover.New())
	app.Use(middleware.Logger())
	app.Use(middleware.CORS())

	// WebSocket upgrade middleware
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// WebSocket handler for payment status updates
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// Handle WebSocket connections for real-time updates
		paymentController.HandleWebSocket(c)
	}))

	// API routes
	api := app.Group("/api/v1")

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)
	auth.Post("/refresh", authController.RefreshToken)

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.JWTAuth(cfg.JWTSecret))

	// User routes
	user := protected.Group("/user")
	user.Get("/profile", authController.GetProfile)
	user.Put("/profile", authController.UpdateProfile)

	// Product routes
	products := api.Group("/products")
	products.Get("/", productController.GetProducts) // Public
	products.Get("/:id", productController.GetProduct) // Public
	
	// Protected product routes (admin only)
	adminProducts := protected.Group("/admin/products")
	adminProducts.Use(middleware.RequireRole("admin"))
	adminProducts.Post("/", productController.CreateProduct)
	adminProducts.Put("/:id", productController.UpdateProduct)
	adminProducts.Delete("/:id", productController.DeleteProduct)

	// Order routes (protected)
	orders := protected.Group("/orders")
	orders.Get("/", orderController.GetUserOrders)
	orders.Post("/", orderController.CreateOrder)
	orders.Get("/:id", orderController.GetOrder)
	orders.Put("/:id/cancel", orderController.CancelOrder)

	// Admin order routes
	adminOrders := protected.Group("/admin/orders")
	adminOrders.Use(middleware.RequireRole("admin"))
	adminOrders.Get("/", orderController.GetAllOrders)
	adminOrders.Put("/:id/status", orderController.UpdateOrderStatus)

	// Payment routes (protected)
	payments := protected.Group("/payments")
	payments.Post("/mpesa/initiate", paymentController.InitiateMpesaPayment)
	payments.Post("/mpesa/callback", paymentController.MpesaCallback) // This might need to be public for M-Pesa callbacks
	payments.Get("/:id/status", paymentController.GetPaymentStatus)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "E-commerce API is running",
		})
	})

	// Handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}