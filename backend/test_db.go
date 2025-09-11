//go:build ignore

package main

import (
	"ecommerce-platform/internal/config"
	"ecommerce-platform/internal/database"
	"log"
)

func main() {
	// Load configuration (which loads .env)
	cfg := config.LoadConfig()

	// Connect to database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	log.Println("Database setup completed successfully!")
}
