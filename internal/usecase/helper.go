package usecase

import (
	"encoding/json"

	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/backend-magang/eniqilo-store/utils/helper"
)

func (u *usecase) validateProductFilter(req *entity.GetListProductRequest) {
	var (
		validSortByVal = []string{"asc", "desc"}
		booleanVal     = []string{"true", "false"}
		categoriesVal  = []string{"Clothing", "Accessories", "Footwear", "Beverages"}
	)

	if req.IsAvailable != "" && !helper.IsInArray(req.IsAvailable, booleanVal) {
		req.IsAvailable = ""
	}

	if req.InStock != "" && !helper.IsInArray(req.InStock, booleanVal) {
		req.InStock = ""
	}

	if req.Category != "" && !helper.IsInArray(req.Category, categoriesVal) {
		req.Category = ""
	}

	if req.Price != "" && !helper.IsInArray(req.Price, validSortByVal) {
		req.Price = ""
	}

	if req.CreatedAt != "" && !helper.IsInArray(req.CreatedAt, validSortByVal) {
		req.CreatedAt = ""
	}
}

func (u *usecase) validateTransactionHistoryFilter(req *entity.GetListTransactionRequest) {
	validSortByVal := []string{"asc", "desc"}

	if req.CreatedAt != "" && !helper.IsInArray(req.CreatedAt, validSortByVal) {
		req.CreatedAt = "desc"
	}
}

func (u *usecase) buildListTransactionResponse(transactions []entity.TransactionDB) []entity.Transaction {
	result := []entity.Transaction{}

	for _, trx := range transactions {
		products := []entity.ProductDetail{}

		json.Unmarshal([]byte(trx.ProductDetails), &products)
		result = append(result, entity.Transaction{
			ID:             trx.ID,
			CustomerID:     trx.CustomerID,
			ProductDetails: products,
			Paid:           trx.Paid,
			Change:         trx.Change,
		})
	}

	return result
}
