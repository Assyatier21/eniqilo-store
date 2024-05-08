package lib

import "errors"

var (
	ErrorNotFound          = errors.New("data not found")
	ErrorNoRowsAffected    = errors.New("no rows affected")
	ErrorNoRowsResult      = errors.New("no rows in result set")
	ErrConstraintKey       = errors.New(`pq: duplicate key value violates unique constraint "unique_phone_number_role"`)
	ErrInsufficientStock   = errors.New("insufficient stock product")
	ErrInsufficientPayment = errors.New("insufficient payment")
	ErrWrongChange         = errors.New("wrong change")
)
