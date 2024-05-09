package entity

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type (
	User struct {
		ID          string         `json:"userId,omitempty" db:"id"`
		Name        string         `json:"name,omitempty" db:"name"`
		PhoneNumber string         `json:"phoneNumber,omitempty" db:"phone_number"`
		Role        string         `json:"role,omitempty" db:"role"`
		Password    sql.NullString `json:"-" db:"password"`
		CreatedAt   time.Time      `json:"createdAt,omitempty" db:"created_at"`
	}

	RegisterStaffRequest struct {
		PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,startswith=+,validatePhoneNumber"`
		Name        string `json:"name" validate:"required,min=5,max=50"`
		Password    string `json:"password" validate:"required,min=5,max=15"`
	}

	LoginStaffRequest struct {
		PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,startswith=+,validatePhoneNumber"`
		Password    string `json:"password" validate:"required,min=5,max=15"`
	}

	UserJWT struct {
		ID          string `json:"userId,omitempty"`
		PhoneNumber string `json:"phoneNumber"`
		Name        string `json:"name"`
		Token       string `json:"accessToken"`
	}

	UserClaims struct {
		ID        int                  `json:"id"`
		Name      string               `json:"name"`
		Email     string               `json:"email"`
		ExpiredAt time.Time            `json:"expired_at"`
		Claims    jwt.RegisteredClaims `json:"claims"`
	}

	UserClaimsResponse struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		ExpiredAt time.Time `json:"expired_at"`
	}

	RegisterCustomerRequest struct {
		PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,startswith=+,validatePhoneNumber"`
		Name        string `json:"name" validate:"required,min=5,max=50"`
	}

	RegisterCustomerResponse struct {
		ID          string `json:"userId,omitempty"`
		PhoneNumber string `json:"phoneNumber"`
		Name        string `json:"name"`
	}

	GetListUserRequest struct {
		PhoneNumber string `param:"phoneNumber"`
		Name        string `param:"name"`
		Role        string `param:"role"`
		Limit       string `param:"limit" validate:"omitempty,number"`
		Offset      string `param:"offset" validate:"omitempty,number"`
	}

	GetListCustomerResponse struct {
		ID          string `json:"userId" db:"id"`
		Name        string `json:"name" db:"name"`
		PhoneNumber string `json:"phoneNumber" db:"phone_number"`
	}
)
