package main

import (
    "log"
    "ecommerce-platform/internal/database"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Connect to database
    db, err := database.NewConnection()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    // Run migrations
    if err := db.RunMigrations(); err != nil {
        log.Fatal("Failed to run migrations:", err)
    }

    log.Println("Database setup completed successfully!")
}
