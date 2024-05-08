package usecase

import (
	"context"
	"net/http"

	"github.com/backend-magang/eniqilo-store/models"
	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/backend-magang/eniqilo-store/utils/constant"
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
