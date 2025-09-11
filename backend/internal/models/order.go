package models

import (
    "encoding/json"
    "time"
)

type Order struct {
    ID              int             `json:"id" db:"id"`
    OrderNumber     string          `json:"order_number" db:"order_number"`
    UserID          int             `json:"user_id" db:"user_id"`
    Status          string          `json:"status" db:"status"`
    TotalAmount     float64         `json:"total_amount,string" db:"total_amount"`
    ShippingAddress json.RawMessage `json:"shipping_address" db:"shipping_address"`
    BillingAddress  json.RawMessage `json:"billing_address" db:"billing_address"`
    CreatedAt       time.Time       `json:"created_at" db:"created_at"`
    UpdatedAt       time.Time       `json:"updated_at" db:"updated_at"`
    OrderItems      []OrderItem     `json:"order_items,omitempty"`
}

type OrderItem struct {
    ID           int     `json:"id" db:"id"`
    OrderID      int     `json:"order_id" db:"order_id"`
    ProductID    int     `json:"product_id" db:"product_id"`
    ProductName  string  `json:"product_name" db:"product_name"`
    ProductPrice float64 `json:"product_price,string" db:"product_price"`
    Quantity     int     `json:"quantity" db:"quantity"`
    TotalPrice   float64 `json:"total_price,string" db:"total_price"`
}

type Address struct {
    Street   string `json:"street"`
    City     string `json:"city"`
    State    string `json:"state"`
    ZipCode  string `json:"zip_code"`
    Country  string `json:"country"`
}

type CreateOrderRequest struct {
    CartItems       []CartItem `json:"cart_items" validate:"required"`
    ShippingAddress Address    `json:"shipping_address" validate:"required"`
    BillingAddress  Address    `json:"billing_address" validate:"required"`
}

type CartItem struct {
    ProductID int `json:"product_id" validate:"required"`
    Quantity  int `json:"quantity" validate:"required,gt=0"`
}