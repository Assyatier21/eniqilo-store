package api

import (
	"github.com/backend-magang/eniqilo-store/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	GetListProduct(c echo.Context) (err error)
}

type handler struct {
	logger  *logrus.Logger
	usecase usecase.UsecaseHandler
}

func NewHandler(log *logrus.Logger, usecase usecase.UsecaseHandler) Handler {
	return &handler{
		logger:  log,
		usecase: usecase,
	}
}
