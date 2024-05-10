package usecase

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/backend-magang/eniqilo-store/models"
	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/backend-magang/eniqilo-store/models/lib"
	"github.com/backend-magang/eniqilo-store/utils/constant"
	"github.com/backend-magang/eniqilo-store/utils/helper"
	"github.com/backend-magang/eniqilo-store/utils/pkg"
	"github.com/spf13/cast"
)

func (u *usecase) GetListProduct(ctx context.Context, req entity.GetListProductRequest) models.StandardResponseReq {
	var (
		products = []entity.Product{}
	)

	u.validateProductFilter(&req)

	products, err := u.repository.GetListProduct(ctx, req)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_PRODUCTS, Error: err}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS, Data: products}
}

func (u *usecase) CreateProduct(ctx context.Context, req entity.CreateProductRequest) models.StandardResponseReq {
	now := time.Now()

	newProduct := entity.Product{
		ID:          helper.NewULID(),
		Name:        req.Name,
		SKU:         req.SKU,
		Category:    req.Category,
		ImageURL:    req.ImageURL,
		Price:       req.Price,
		Stock:       req.Stock,
		Location:    req.Location,
		IsAvailable: true,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	product, err := u.repository.InsertProduct(ctx, newProduct)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	resp := entity.CreateProductResponse{
		ID:        cast.ToString(product.ID),
		CreatedAt: product.CreatedAt,
	}

	return models.StandardResponseReq{Code: http.StatusCreated, Message: constant.SUCCESS_ADD_PRODUCT, Data: resp, Error: nil}
}

func (u *usecase) CheckoutProduct(ctx context.Context, req entity.CheckoutProductRequest) models.StandardResponseReq {
	_, err := u.repository.FindUserByID(ctx, req.CustomerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED_GET_USER, Error: err}
		}
	}

	if err = pkg.WithTransaction(ctx, u.cfg.SqlTrx, func(ctx context.Context) (err error) {
		err = u.repository.CheckoutProducts(ctx, req)
		return err
	}); err != nil {
		code := http.StatusInternalServerError

		if errors.Is(err, sql.ErrNoRows) {
			code = http.StatusNotFound
		}

		if errors.Is(err, lib.ErrInsufficientPayment) || errors.Is(err, lib.ErrWrongChange) || errors.Is(err, lib.ErrInsufficientStock) {
			code = http.StatusBadRequest
		}

		return models.StandardResponseReq{Code: code, Message: err.Error()}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS, Data: nil}
}

func (u *usecase) DeleteProduct(ctx context.Context, req entity.DeleteProductRequest) models.StandardResponseReq {
	err := u.repository.DeleteProduct(ctx, req.ProductID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED_PRODUCT_NOT_FOUND, Error: err}
		}
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS}
}
