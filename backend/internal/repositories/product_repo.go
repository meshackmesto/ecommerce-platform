package repositories

import (
	"ecommerce-platform/internal/database"
)

type ProductRepository struct {
	db *database.DB
}

func NewProductRepository(db *database.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Implement your Product repository methods here based on the schema...
