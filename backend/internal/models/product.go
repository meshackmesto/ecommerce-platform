package models

import (
    "encoding/json"
    "time"
)

type Product struct {
    ID              int             `json:"id" db:"id"`
    Name            string          `json:"name" db:"name"`
    Description     *string         `json:"description" db:"description"`
    Price           float64         `json:"price,string" db:"price"`
    StockQuantity   int             `json:"stock_quantity" db:"stock_quantity"`
    Category        *string         `json:"category" db:"category"`
    Brand           *string         `json:"brand" db:"brand"`
    SKU             *string         `json:"sku" db:"sku"`
    Images          json.RawMessage `json:"images" db:"images"` // ["url1", "url2"]
    IsActive        bool            `json:"is_active" db:"is_active"`
    IsFeatured      bool            `json:"is_featured" db:"is_featured"`
    CreatedBy       *int            `json:"created_by" db:"created_by"`
    CreatedAt       time.Time       `json:"created_at" db:"created_at"`
    UpdatedAt       time.Time       `json:"updated_at" db:"updated_at"`
}

type CreateProductRequest struct {
    Name          string  `json:"name" validate:"required"`
    Description   string  `json:"description"`
    Price         float64 `json:"price" validate:"required,gt=0"`
    StockQuantity int     `json:"stock_quantity" validate:"gte=0"`
    Category      string  `json:"category"`
    Brand         string  `json:"brand"`
    SKU           string  `json:"sku"`
}