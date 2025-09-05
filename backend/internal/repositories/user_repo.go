package repositories

import (
    "database/sql"
    "errors"
    "fmt"

    "ecommerce-platform/internal/database"
    "ecommerce-platform/internal/models"
    "github.com/google/uuid"
)

type UserRepository struct {
    db *database.DB
}

func NewUserRepository(db *database.DB) *UserRepository {
    return &UserRepository{db: db}
}

// GetByEmail finds a user by email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
    query := `
        SELECT id, email, password_hash, first_name, last_name, phone, role, is_active, created_at, updated_at
        FROM users 
        WHERE email = $1 AND is_active = true
    `
    
    var user models.User
    var phone sql.NullString
    
    err := r.db.QueryRow(query, email).Scan(
        &user.ID, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName,
        &phone, &user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("user not found")
        }
        return nil, err
    }
    
    if phone.Valid {
        user.Phone = &phone.String
    }
    
    return &user, nil
}

// GetByID finds a user by ID
func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
    query := `
        SELECT id, email, password_hash, first_name, last_name, phone, role, is_active, created_at, updated_at
        FROM users 
        WHERE id = $1 AND is_active = true
    `
    
    var user models.User
    var phone sql.NullString
    
    err := r.db.QueryRow(query, id).Scan(
        &user.ID, &user.Email, &user.PasswordHash, &user.FirstName, &user.LastName,
        &phone, &user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("user not found")
        }
        return nil, err
    }
    
    if phone.Valid {
        user.Phone = &phone.String
    }
    
    return &user, nil
}

// Create creates a new user
func (r *UserRepository) Create(user *models.User) error {
    query := `
        INSERT INTO users (email, password_hash, first_name, last_name, phone, role)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at, updated_at, is_active
    `
    
    err := r.db.QueryRow(query, 
        user.Email, user.PasswordHash, user.FirstName, user.LastName, user.Phone, user.Role,
    ).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.IsActive)
    
    if err != nil {
        // Check if it's a unique constraint violation (email already exists)
        if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {
            return errors.New("email already exists")
        }
        return fmt.Errorf("failed to create user: %w", err)
    }
    
    return nil
}

// EmailExists checks if an email already exists
func (r *UserRepository) EmailExists(email string) (bool, error) {
    query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
    
    var exists bool
    err := r.db.QueryRow(query, email).Scan(&exists)
    if err != nil {
        return false, err
    }
    
    return exists, nil
}