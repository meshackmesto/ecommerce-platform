// File: backend/scripts/seeder.go (Final Version for Fake Store API)
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"ecommerce-platform/internal/config"
	"ecommerce-platform/internal/database"
	"github.com/joho/godotenv"
)

// This struct matches the Fake Store API's product structure
type APIProduct struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}

func main() {
	// We still load the .env file for the DATABASE credentials
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file from parent directory: %v", err)
	}

	// --- 1. Fetch Data from Fake Store API ---
	log.Println("Starting to fetch products from Fake Store API...")
	
	url := "https://fakestoreapi.com/products"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request to Fake Store API: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Fake Store API returned a non-200 status code: %d", res.StatusCode)
	}

	body, _ := io.ReadAll(res.Body)
	
	var apiProducts []APIProduct
	if err := json.Unmarshal(body, &apiProducts); err != nil {
		log.Fatalf("Failed to parse JSON response: %v.", err)
	}
	
	if len(apiProducts) == 0 {
		log.Fatal("No products were parsed from the Fake Store API response.")
	}

	log.Printf("Successfully parsed %d products from the API.", len(apiProducts))

	// --- 2. Connect to our Database ---
	cfg := config.LoadConfig()
	db, err := database.InitDB(cfg)
	if err != nil { log.Fatalf("Failed to initialize database: %v", err) }
	defer db.Close()

	// --- 3. Clear Existing Products ---
	log.Println("Clearing existing products from the database...")
	if _, err := db.Exec(`TRUNCATE TABLE products RESTART IDENTITY CASCADE;`); err != nil {
		log.Fatalf("Failed to clear products table: %v", err)
	}
	
	// --- 4. Insert New Products into Database ---
	log.Println("Inserting new products...")
	insertedCount := 0
	for _, p := range apiProducts {
		// The API gives one image URL, but our database expects a JSON array of strings.
		// We will create a JSON array containing just the one image.
		imagesJSON, _ := json.Marshal([]string{p.Image})

		// The API doesn't provide a brand, so we will leave it null.
		var brand sql.NullString
		brand.Valid = false

		query := `
			INSERT INTO products (name, description, price, stock_quantity, category, brand, images, is_active, sku)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			ON CONFLICT (name) DO NOTHING;
		`
		sku := fmt.Sprintf("FS-%d", p.ID) // Create a unique SKU
		_, dbErr := db.Exec(query, p.Title, p.Description, p.Price, 100, p.Category, brand, string(imagesJSON), true, sku)
		if dbErr != nil {
			log.Printf("Failed to insert product '%s': %v", p.Title, dbErr)
		} else {
			insertedCount++
		}
	}
	
	log.Printf("Seeding complete. Inserted %d new products into the database.", insertedCount)
}