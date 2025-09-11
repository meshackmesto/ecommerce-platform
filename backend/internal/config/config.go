package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
	JWTSecret  string
	
	// M-Pesa Configuration
	MpesaConsumerKey    string
	MpesaConsumerSecret string
	MpesaEnvironment    string // sandbox or production
	
	// File upload configuration
	MaxFileSize   int64
	UploadDir     string
	AllowedTypes  []string
}

func LoadConfig() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "ecommerce"),
		Port:       getEnv("PORT", "8080"),
		JWTSecret:  getEnv("JWT_SECRET", "your-super-secret-jwt-key"),
		
		// M-Pesa settings
		MpesaConsumerKey:    getEnv("MPESA_CONSUMER_KEY", "test_consumer_key"),
		MpesaConsumerSecret: getEnv("MPESA_CONSUMER_SECRET", "test_consumer_secret"),
		MpesaEnvironment:    getEnv("MPESA_ENVIRONMENT", "sandbox"),
		
		// File upload settings
		MaxFileSize:  10 * 1024 * 1024, // 10MB
		UploadDir:    getEnv("UPLOAD_DIR", "./uploads"),
		AllowedTypes: []string{"image/jpeg", "image/png", "image/gif", "image/webp"},
	}

	// Create upload directory if it doesn't exist
	if _, err := os.Stat(config.UploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(config.UploadDir, 0755); err != nil {
			log.Fatalf("Failed to create upload directory: %v", err)
		}
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) DatabaseURL() string {
	return "postgres://" + c.DBUser + ":" + c.DBPassword + "@" + c.DBHost + ":" + c.DBPort + "/" + c.DBName + "?sslmode=disable"
}