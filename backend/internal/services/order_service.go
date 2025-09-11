package services

import (
	"ecommerce-platform/internal/repositories"
)

type OrderService struct {
	orderRepo   *repositories.OrderRepository
	productRepo *repositories.ProductRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository, productRepo *repositories.ProductRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo, productRepo: productRepo}
}

// Implement your Order service methods here...