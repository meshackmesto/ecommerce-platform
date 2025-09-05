package database

import (
    "database/sql"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "sort"
    "strings"

    _ "github.com/lib/pq"
)

type DB struct {
    *sql.DB
}

func NewConnection() (*DB, error) {
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnv("DB_PORT", "5432")
    dbUser := getEnv("DB_USER", "postgres")
    dbPassword := getEnv("DB_PASSWORD", "password")
    dbName := getEnv("DB_NAME", "ecommerce")

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    log.Println("Successfully connected to PostgreSQL database")
    return &DB{db}, nil
}

func (db *DB) RunMigrations() error {
    // Create migrations table if it doesn't exist
    createMigrationsTable := `
        CREATE TABLE IF NOT EXISTS migrations (
            id SERIAL PRIMARY KEY,
            filename VARCHAR(255) NOT NULL UNIQUE,
            executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `
    
    if _, err := db.Exec(createMigrationsTable); err != nil {
        return fmt.Errorf("failed to create migrations table: %w", err)
    }

    migrationsDir := "./internal/database/migrations"
    files, err := ioutil.ReadDir(migrationsDir)
    if err != nil {
        return fmt.Errorf("failed to read migrations directory: %w", err)
    }

    // Sort migration files
    var migrationFiles []string
    for _, file := range files {
        if strings.HasSuffix(file.Name(), ".sql") {
            migrationFiles = append(migrationFiles, file.Name())
        }
    }
    sort.Strings(migrationFiles)

    for _, filename := range migrationFiles {
        // Check if migration already executed
        var count int
        err := db.QueryRow("SELECT COUNT(*) FROM migrations WHERE filename = $1", filename).Scan(&count)
        if err != nil {
            return fmt.Errorf("failed to check migration status: %w", err)
        }

        if count > 0 {
            log.Printf("Migration %s already executed, skipping", filename)
            continue
        }

        // Execute migration
        migrationPath := filepath.Join(migrationsDir, filename)
        content, err := ioutil.ReadFile(migrationPath)
        if err != nil {
            return fmt.Errorf("failed to read migration file %s: %w", filename, err)
        }

        if _, err := db.Exec(string(content)); err != nil {
            return fmt.Errorf("failed to execute migration %s: %w", filename, err)
        }

        // Record migration as executed
        if _, err := db.Exec("INSERT INTO migrations (filename) VALUES ($1)", filename); err != nil {
            return fmt.Errorf("failed to record migration %s: %w", filename, err)
        }

        log.Printf("Successfully executed migration: %s", filename)
    }

    return nil
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}