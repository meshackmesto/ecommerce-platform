package repositories

import (
	"ecommerce-platform/internal/database"
)

type TransactionRepository struct {
	db *database.DB
}

func NewTransactionRepository(db *database.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Implement your Transaction repository methods here...