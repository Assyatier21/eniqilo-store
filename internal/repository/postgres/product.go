package postgres

import (
	"context"

	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/backend-magang/eniqilo-store/models/lib"
	"github.com/backend-magang/eniqilo-store/utils/pkg"
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

func (r *repository) GetActiveProductByIDWithTx(ctx context.Context, id string) (entity.Product, error) {
	var (
		err    error
		result = entity.Product{}

		query = `SELECT * FROM products WHERE id = $1 AND is_available = true FOR UPDATE`
	)

	err = r.db.QueryRowxContext(ctx, query, id).StructScan(&result)
	if err != nil {
		r.logger.Errorf("[Repository][Product][GetByIDWithTx] failed to query, err: %s", err.Error())
		return result, err
	}

	return result, err
}

func (r *repository) CheckoutProducts(ctx context.Context, req entity.CheckoutProductRequest) error {
	var (
		total float64 = 0
		tx, _         = pkg.ExtractTx(ctx)
	)

	for _, v := range req.ProductsCheckoutRequest {
		product, err := r.GetActiveProductByIDWithTx(ctx, v.ProductID)
		if err != nil {
			return err
		}

		if product.Stock <= 0 || product.Stock-v.Quantity < 0 {
			return lib.ErrInsufficientStock
		}

		product.Stock -= v.Quantity
		_, err = tx.ExecContext(ctx, "UPDATE products SET stock = $1 WHERE id = $2", product.Stock, product.ID)
		if err != nil {
			r.logger.Errorf("[Repository][Product][CheckoutProduct] failed to query, err: %s", err.Error())
			return err
		}

		total += product.Price * float64(v.Quantity)
	}

	if req.Paid < total {
		return lib.ErrInsufficientPayment
	}

	if req.Change != (req.Paid - total) {
		return lib.ErrWrongChange
	}

	return nil
}
