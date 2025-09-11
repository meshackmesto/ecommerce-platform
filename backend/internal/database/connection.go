package database

import (
    "database/sql"
    "ecommerce-platform/internal/config"
    "fmt"
    "io/ioutil"
    "log"
    "path/filepath"
    "sort"
    "strings"

    _ "github.com/lib/pq"
)

type DB struct {
    *sql.DB
}

// InitDB initializes the database connection using the provided configuration.
func InitDB(cfg *config.Config) (*DB, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

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

// RunMigrations executes all .sql migration files that haven't been run yet.
func RunMigrations(db *DB) error {
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

    var migrationFiles []string
    for _, file := range files {
        if strings.HasSuffix(file.Name(), ".sql") {
            migrationFiles = append(migrationFiles, file.Name())
        }
    }
    sort.Strings(migrationFiles)

    for _, filename := range migrationFiles {
        var count int
        err := db.QueryRow("SELECT COUNT(*) FROM migrations WHERE filename = $1", filename).Scan(&count)
        if err != nil {
            return fmt.Errorf("failed to check migration status: %w", err)
        }

        if count > 0 {
            continue // Skip already executed migration
        }

        migrationPath := filepath.Join(migrationsDir, filename)
        content, err := ioutil.ReadFile(migrationPath)
        if err != nil {
            return fmt.Errorf("failed to read migration file %s: %w", filename, err)
        }

        if _, err := db.Exec(string(content)); err != nil {
            return fmt.Errorf("failed to execute migration %s: %w", filename, err)
        }

        if _, err := db.Exec("INSERT INTO migrations (filename) VALUES ($1)", filename); err != nil {
            return fmt.Errorf("failed to record migration %s: %w", filename, err)
        }

        log.Printf("Successfully executed migration: %s", filename)
    }

    log.Println("Database migrations are up to date.")
    return nil
}