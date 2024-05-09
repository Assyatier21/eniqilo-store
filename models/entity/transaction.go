package entity

import "time"

type Transaction struct {
	ID             string          `json:"id"`
	CustomerID     string          `json:"customerId"`
	ProductDetails []ProductDetail `json:"productDetails"`
	Paid           float64         `json:"paid"`
	Change         float64         `json:"change"`
}

type ProductDetail struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type TransactionDB struct {
	ID             string    `db:"id"`
	CustomerID     string    `db:"customer_id"`
	ProductDetails string    `db:"product_details"`
	Paid           float64   `db:"paid"`
	Change         float64   `db:"change"`
	CreatedAt      time.Time `db:"created_at"`
}

type GetListTransactionRequest struct {
	CustomerID string `param:"customerId"`
	Limit      string `param:"limit" validate:"omitempty,number"`
	Offset     string `param:"offset" validate:"omitempty,number"`
	CreatedAt  string `param:"createdAt"`
}
