package postgres

import (
	"context"

	"github.com/backend-magang/eniqilo-store/models/entity"
)

func (r *repository) GetListProduct(ctx context.Context, req entity.GetListProductRequest) ([]entity.Product, error) {
	result := []entity.Product{}

	query, args := buildQueryGetListProducts(req)
	query = r.db.Rebind(query)

	err := r.db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		r.logger.Errorf("[Repository][Product][GetList] failed to query, err: %s", err.Error())
		return result, err
	}

	return result, err
}
