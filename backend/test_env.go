//go:build ignore

package main

import (
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file:", err)
    }

    fmt.Println("Environment variables loaded:")
    fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
    fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
    fmt.Println("DB_USER:", os.Getenv("DB_USER"))
    fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
    fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
}
