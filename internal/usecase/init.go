package usecase

import (
	"context"

	"github.com/backend-magang/eniqilo-store/config"
	"github.com/backend-magang/eniqilo-store/internal/repository/postgres"
	"github.com/backend-magang/eniqilo-store/models"
	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/sirupsen/logrus"
)

type UsecaseHandler interface {
	GetListProduct(ctx context.Context, req entity.GetListProductRequest) models.StandardResponseReq
}

type usecase struct {
	cfg        config.Config
	logger     *logrus.Logger
	repository postgres.RepositoryHandler
}

func NewUsecase(cfg config.Config, log *logrus.Logger, repository postgres.RepositoryHandler) UsecaseHandler {
	return &usecase{
		cfg:        cfg,
		logger:     log,
		repository: repository,
	}
}
