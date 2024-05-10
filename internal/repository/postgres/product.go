package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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

func (r *repository) InsertProduct(ctx context.Context, req entity.Product) (result entity.Product, err error) {
	query := `INSERT INTO products (id, name, sku, category, image_url, notes, price, stock, location, is_available, created_at, updated_at, deleted_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) 
		RETURNING *`

	err = r.db.QueryRowxContext(ctx,
		query,
		req.ID,
		req.Name,
		req.SKU,
		req.Category,
		req.ImageURL,
		req.Notes,
		req.Price,
		req.Stock,
		req.Location,
		req.IsAvailable,
		req.CreatedAt,
		req.UpdatedAt,
		req.DeletedAt,
	).StructScan(&result)

	if err != nil {
		r.logger.Errorf("[Repository][Product][InsertProduct] failed to insert new product, err: %s", err.Error())
		return
	}

	return
}

func (r *repository) GetActiveProductsByIDsWithTx(ctx context.Context, ids []interface{}) ([]entity.Product, error) {
	var (
		err     error
		results []entity.Product
	)

	stringIDs := make([]string, len(ids))
	for i, id := range ids {
		stringIDs[i] = cast.ToString(id)
	}

	inClause := "'" + strings.Join(stringIDs, "','") + "'"
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (%s) FOR UPDATE", inClause)
	query = r.db.Rebind(query)

	err = r.db.SelectContext(ctx, &results, query)
	if err != nil {
		r.logger.Errorf("[Repository][Product][GetActiveProductsByIDsWithTx] failed to query, err: %s", err.Error())
		return nil, err
	}

	idMap := make(map[string]struct{}, len(results))
	for _, product := range results {
		idMap[product.ID] = struct{}{}
		if !product.IsAvailable {
			return nil, lib.ErrorProductNotAvailable
		}
	}

	for _, id := range stringIDs {
		if _, ok := idMap[id]; !ok {
			return nil, sql.ErrNoRows
		}
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

	if *req.Change != (req.Paid - total) {
		return lib.ErrWrongChange
	}

	productsJSON, _ := json.Marshal(req.ProductsCheckoutRequest)
	transaction := entity.TransactionDB{
		ID:             helper.NewULID(),
		CustomerID:     req.CustomerID,
		ProductDetails: cast.ToString(productsJSON),
		Paid:           req.Paid,
		Change:         *req.Change,
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
