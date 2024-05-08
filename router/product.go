package router

import (
	"github.com/backend-magang/eniqilo-store/internal/handler/api"
	"github.com/labstack/echo/v4"
)

func InitProductRouter(e *echo.Echo, handler api.Handler) {
	v1 := e.Group("/v1")
	product := v1.Group("/product")

	product.GET("", handler.GetListProduct)
	product.POST("/checkout", handler.CheckoutProduct)

}
