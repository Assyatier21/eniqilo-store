package postgres

import (
	"context"
	"strings"
	"time"

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

func (r *repository) GetActiveProductsByIDsWithTx(ctx context.Context, ids []interface{}) ([]entity.Product, error) {
	var (
		err          error
		results      []entity.Product
		placeholders []string
	)

	for range ids {
		placeholders = append(placeholders, "?")
	}

	query := `SELECT * FROM products WHERE id IN (` + strings.Join(placeholders, ",") + `) AND is_available = true FOR UPDATE`
	query = r.db.Rebind(query)

	err = r.db.SelectContext(ctx, &results, query, ids...)
	if err != nil {
		r.logger.Errorf("[Repository][Product][GetActiveProductsByIDsWithTx] failed to query, err: %s", err.Error())
		return nil, err
	}

	return results, nil
}

func (r *repository) CheckoutProducts(ctx context.Context, req entity.CheckoutProductRequest) error {
	var (
		total      float64 = 0
		productIDs []interface{}
		now        = time.Now()
		tx, _      = pkg.ExtractTx(ctx)
	)

	for _, v := range req.ProductsCheckoutRequest {
		productIDs = append(productIDs, v.ProductID)
	}

	activeProducts, err := r.GetActiveProductsByIDsWithTx(ctx, productIDs)
	if err != nil {
		return err
	}

	productMap := make(map[string]entity.Product)
	for _, p := range activeProducts {
		productMap[p.ID] = p
	}

	for _, v := range req.ProductsCheckoutRequest {
		product := productMap[v.ProductID]

		if product.Stock <= 0 || product.Stock-v.Quantity < 0 {
			return lib.ErrInsufficientStock
		}

		product.Stock -= v.Quantity
		_, err = tx.ExecContext(ctx, "UPDATE products SET stock = $1, updated_at = $2 WHERE id = $3", product.Stock, now, product.ID)
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
