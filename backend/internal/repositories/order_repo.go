package repositories

import (
	"ecommerce-platform/internal/database"
)

type OrderRepository struct {
	db *database.DB
}

func NewOrderRepository(db *database.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// Implement your Order repository methods here...