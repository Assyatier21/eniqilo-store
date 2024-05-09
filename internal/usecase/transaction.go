package usecase

import (
	"context"
	"net/http"

	"github.com/backend-magang/eniqilo-store/models"
	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/backend-magang/eniqilo-store/utils/constant"
)

func (u *usecase) GetHistoryTransaction(ctx context.Context, req entity.GetListTransactionRequest) models.StandardResponseReq {
	var (
		err          error
		transactions = []entity.TransactionDB{}
	)

	u.validateTransactionHistoryFilter(&req)

	transactions, err = u.repository.GetListTransaction(ctx, req)
	if err != nil {
		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_PRODUCTS, Error: err}
	}

	trxResponse := u.buildListTransactionResponse(transactions)

	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS, Data: trxResponse}
}
