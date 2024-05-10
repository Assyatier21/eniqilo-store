package api

import (
	"github.com/backend-magang/eniqilo-store/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	RegisterStaff(c echo.Context) (err error)
	LoginStaff(c echo.Context) (err error)

	RegisterCustomer(c echo.Context) (err error)
	GetListCustomer(c echo.Context) (err error)

	CreateProduct(c echo.Context) (err error)
	UpdateProduct(c echo.Context) (err error)
	GetListProduct(c echo.Context) (err error)
	GetListProductCustomer(c echo.Context) (err error)
	CheckoutProduct(c echo.Context) (err error)
	DeleteProduct(c echo.Context) (err error)

	GetListTransaction(c echo.Context) (err error)
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
