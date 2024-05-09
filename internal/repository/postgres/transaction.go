package postgres

import (
	"context"
	"time"

	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/backend-magang/eniqilo-store/utils/pkg"
)

func (r *repository) InsertTransaction(ctx context.Context, req entity.TransactionDB) error {
	var (
		err   error
		tx, _ = pkg.ExtractTx(ctx)
		now   = time.Now()
		args  = []interface{}{req.ID, req.CustomerID, req.ProductDetails, req.Paid, req.Change, now}
	)
	query := `INSERT INTO transactions (id, customer_id, product_details, paid, change, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING *`

	if tx != nil {
		_, err = tx.ExecContext(ctx, query, args...)
	} else {
		_, err = r.db.ExecContext(ctx, query, args...)
	}
	if err != nil {
		r.logger.Errorf("[Repository][Transaction][Insert] failed to query, err: %s", err.Error())
		return err
	}

	return err
}

func (r *repository) GetListTransaction(ctx context.Context, req entity.GetListTransactionRequest) ([]entity.TransactionDB, error) {
	result := []entity.TransactionDB{}

	query, args := buildQueryGetListTransactions(req)
	query = r.db.Rebind(query)

	err := r.db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		r.logger.Errorf("[Repository][Transaction][GetList] failed to query, err: %s", err.Error())
		return result, err
	}

	return result, err
}
