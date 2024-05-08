package usecase

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/backend-magang/eniqilo-store/middleware"
	"github.com/backend-magang/eniqilo-store/models"
	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/backend-magang/eniqilo-store/models/lib"
	"github.com/backend-magang/eniqilo-store/utils/constant"
	"github.com/backend-magang/eniqilo-store/utils/helper"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) RegisterStaff(ctx context.Context, req entity.RegisterStaffRequest) models.StandardResponseReq {
	var (
		user    = entity.User{}
		newUser = entity.User{}
		now     = time.Now()
	)

	hashPassword := helper.HashPassword(req.Password, cast.ToInt(u.cfg.BCryptSalt))
	newUser = entity.User{
		ID:          helper.NewULID(),
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Role:        constant.ROLE_STAFF,
		Password:    sql.NullString{String: hashPassword, Valid: true},
		CreatedAt:   now,
	}

	user, err := u.repository.InsertUser(ctx, newUser)
	if err != nil {
		if err.Error() == lib.ErrConstraintKey.Error() {
			return models.StandardResponseReq{Code: http.StatusConflict, Message: constant.PHONE_NUMBER_REGISTERED, Error: nil}
		}
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	newUser.ID = user.ID
	token, _ := middleware.GenerateToken(newUser)

	userJWT := entity.UserJWT{
		ID:          newUser.ID,
		PhoneNumber: newUser.PhoneNumber,
		Name:        newUser.Name,
		Token:       token,
	}

	return models.StandardResponseReq{Code: http.StatusCreated, Message: constant.SUCCESS_REGISTER_USER, Data: userJWT, Error: nil}
}

func (u *usecase) LoginStaff(ctx context.Context, req entity.LoginStaffRequest) models.StandardResponseReq {
	var (
		userJWT = entity.UserJWT{}
		user    = entity.User{}
		token   string
		err     error
	)

	user, err = u.repository.FindUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED, Error: err}
		}
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(req.Password))
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.FAILED_LOGIN}
	}

	token, _ = middleware.GenerateToken(user)
	userJWT = entity.UserJWT{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
		Token:       token,
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_LOGIN, Data: userJWT, Error: nil}
}
