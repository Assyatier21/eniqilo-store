package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"
	"time"

	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/backend-magang/eniqilo-store/models/lib"
	"github.com/backend-magang/eniqilo-store/utils/helper"
	"github.com/backend-magang/eniqilo-store/utils/pkg"
	"github.com/spf13/cast"
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

func (r *repository) GetActiveProductByID(ctx context.Context, id int) (entity.Product, error) {
	result := entity.Product{}
	query := `SELECT * FROM products WHERE id = $1 AND deleted_at IS NULL`

	err := r.db.QueryRowxContext(ctx, query, id).StructScan(&result)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Errorf("[Repository][Product][GetByID] failed to query, err: %s", err.Error())
		return result, err
	}

	return result, err
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

	productsJSON, _ := json.Marshal(req.ProductsCheckoutRequest)
	transaction := entity.TransactionDB{
		ID:             helper.NewULID(),
		CustomerID:     req.CustomerID,
		ProductDetails: cast.ToString(productsJSON),
		Paid:           req.Paid,
		Change:         req.Change,
	}

	err = r.InsertTransaction(ctx, transaction)

	return err
}

func (r *repository) DeleteProduct(ctx context.Context, id string) (err error) {
	var (
		now    = time.Now()
		result = entity.Product{}
		args   = []interface{}{now, sql.NullTime{Time: now, Valid: true}, id}
	)

	query := `UPDATE products 
		SET updated_at = $1, deleted_at = $2
		WHERE id = $3 AND deleted_at IS NULL RETURNING *`

	err = r.db.QueryRowxContext(ctx, query, args...).StructScan(&result)
	if err != nil {
		r.logger.Errorf("[Repository][Product][DeleteProduct] failed to query, err: %s", err.Error())
		return
	}

	return
}
