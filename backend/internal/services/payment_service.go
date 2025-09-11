package services

import (
	"ecommerce-platform/internal/repositories"
)

type PaymentService struct {
	transactionRepo *repositories.TransactionRepository
	orderRepo       *repositories.OrderRepository
}

func NewPaymentService(transactionRepo *repositories.TransactionRepository, orderRepo *repositories.OrderRepository) *PaymentService {
	return &PaymentService{
		transactionRepo: transactionRepo,
		orderRepo:       orderRepo,
	}
}

// Implement your Payment service methods here...