package postgres

import (
	"context"

	"github.com/backend-magang/eniqilo-store/config"
	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type RepositoryHandler interface {
	FindUserByPhoneNumber(ctx context.Context, phoneNumber string) (result entity.User, err error)
	InsertUser(ctx context.Context, req entity.User) (result entity.User, err error)

	GetListProduct(ctx context.Context, req entity.GetListProductRequest) ([]entity.Product, error)
}

type repository struct {
	cfg    config.Config
	db     *sqlx.DB
	logger *logrus.Logger
}

func NewRepository(cfg config.Config, db *sqlx.DB, log *logrus.Logger) RepositoryHandler {
	return &repository{
		cfg:    cfg,
		db:     db,
		logger: log,
	}
}
