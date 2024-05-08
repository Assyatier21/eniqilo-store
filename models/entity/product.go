package entity

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          string       `json:"id" db:"id"`
	Name        string       `json:"name" db:"name"`
	SKU         string       `json:"sku" db:"sku"`
	Category    string       `json:"category" db:"category"`
	ImageURL    string       `json:"image_url" db:"image_url"`
	Price       float64      `json:"price" db:"price"`
	Stock       int          `json:"stock" db:"stock"`
	Location    string       `json:"location" db:"location"`
	IsAvailable bool         `json:"is_available" db:"is_available"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt   sql.NullTime `json:"-" db:"deleted_at"`
}

type GetListProductRequest struct {
	ID          string `param:"id"`
	Limit       string `param:"limit" validate:"omitempty,number"`
	Offset      string `param:"offset" validate:"omitempty,number"`
	Name        string `param:"name"`
	IsAvailable string `param:"isAvaliable"`
	Category    string `param:"category"`
	SKU         string `param:"sku"`
	InStock     string `param:"inStock"`
	Price       string `param:"price"`
	CreatedAt   string `param:"createdAt"`
}

type ProductCheckoutRequest struct {
	ProductID string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

type CheckoutProductRequest struct {
	CustomerID              string                   `json:"customerId"`
	ProductsCheckoutRequest []ProductCheckoutRequest `json:"productDetails" validate:"min=1,dive"`
	Paid                    float64                  `json:"paid" validate:"required,min=1"`
	Change                  float64                  `json:"change" validate:"min=0"`
}

type DeleteProductRequest struct {
	ProductID string `param:"id" validate:"required"`
}
