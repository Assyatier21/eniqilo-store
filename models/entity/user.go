package entity

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type (
	User struct {
		ID          string    `db:"id"`
		Name        string    `db:"name"`
		PhoneNumber string    `db:"phone_number"`
		Role        string    `db:"role"`
		Password    string    `db:"password"`
		CreatedAt   time.Time `db:"created_at"`
	}

	RegisterStaffRequest struct {
		PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,startswith=+,validatePhoneNumber"`
		Name        string `json:"name" validate:"required,min=5,max=50"`
		Password    string `json:"password" validate:"required,min=5,max=15"`
	}

	LoginUserRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=5,max=15"`
	}

	UserJWT struct {
		ID          string `json:"userId,omitempty"`
		PhoneNumber string `json:"phone_number"`
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
)
