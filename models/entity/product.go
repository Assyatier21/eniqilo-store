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
	ImageURL    string       `json:"imageUrl" db:"image_url"`
	Notes       string       `json:"notes" db:"notes"`
	Price       float64      `json:"price" db:"price"`
	Stock       int          `json:"stock" db:"stock"`
	Location    string       `json:"location" db:"location"`
	IsAvailable bool         `json:"isAvailable" db:"is_available"`
	CreatedAt   time.Time    `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time    `json:"-" db:"updated_at"`
	DeletedAt   sql.NullTime `json:"-" db:"deleted_at"`
}

type GetListProductRequest struct {
	ID          string `param:"id"`
	Limit       string `param:"limit" validate:"omitempty,number"`
	Offset      string `param:"offset" validate:"omitempty,number"`
	Name        string `param:"name"`
	IsAvailable string `param:"isAvailable"`
	Category    string `param:"category"`
	SKU         string `param:"sku"`
	InStock     string `param:"inStock"`
	Price       string `param:"price"`
	CreatedAt   string `param:"createdAt"`
}

type CreateProductRequest struct {
	ID          string
	Name        string  `json:"name" validate:"required,min=1,max=30"`
	SKU         string  `json:"sku" validate:"required,min=1,max=30"`
	Category    string  `json:"category" validate:"required,oneof=Clothing Accessories Footwear Beverages"`
	ImageURL    string  `json:"imageUrl" validate:"required,validateImageURL"`
	Notes       string  `json:"notes" validate:"required,min=1,max=200"`
	Price       float64 `json:"price" validate:"required,min=1"`
	Stock       int     `json:"stock" validate:"required,min=0,max=100000"`
	Location    string  `json:"location" validate:"required,min=1,max=200"`
	IsAvailable *bool   `json:"isAvailable"`
}

type CreateProductResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type UpdateProductRequest struct {
	ID          string  `param:"id" validate:"required"`
	Name        string  `json:"name" validate:"required,min=1,max=30"`
	SKU         string  `json:"sku" validate:"required,min=1,max=30"`
	Category    string  `json:"category" validate:"required,oneof=Clothing Accessories Footwear Beverages"`
	Notes       string  `json:"notes" validate:"required,min=1,max=200"`
	ImageURL    string  `json:"imageUrl" validate:"required,validateImageURL"`
	Price       float64 `json:"price" validate:"required,min=1"`
	Stock       int     `json:"stock" validate:"required,min=0,max=100000"`
	Location    string  `json:"location" validate:"required,min=1,max=200"`
	IsAvailable *bool   `json:"isAvailable"`
}

type UpdateProductResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	SKU         string    `json:"sku"`
	Category    string    `json:"category"`
	Notes       string    `json:"notes"`
	ImageURL    string    `json:"imageUrl"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Location    string    `json:"location"`
	IsAvailable *bool     `json:"isAvailable"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
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
