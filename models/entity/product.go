package entity

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          string       `db:"id"`
	Name        string       `db:"name"`
	SKU         string       `db:"sku"`
	Category    string       `db:"category"`
	ImageURL    string       `db:"image_url"`
	Price       float64      `db:"price"`
	Stock       int          `db:"stock"`
	Location    string       `db:"location"`
	IsAvailable bool         `db:"is_available"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}
