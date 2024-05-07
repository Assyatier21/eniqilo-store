package usecase

import (
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
