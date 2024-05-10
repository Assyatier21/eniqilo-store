package router

import (
	"github.com/backend-magang/eniqilo-store/internal/handler/api"
	md "github.com/backend-magang/eniqilo-store/middleware"
	"github.com/labstack/echo/v4"
)

func InitProductRouter(e *echo.Echo, handler api.Handler) {
	v1 := e.Group("/v1")
	product := v1.Group("/product", md.TokenValidationMiddleware())

	product.GET("", handler.GetListProduct)
	product.GET("/checkout/history", handler.GetListTransaction)
	product.POST("", handler.CreateProduct)
	product.POST("/checkout", handler.CheckoutProduct)
	product.PUT("/:id", handler.UpdateProduct)
	product.DELETE("/:id", handler.DeleteProduct)

	// For API No Auth
	productCustomer := v1.Group("/product/customer")
	productCustomer.GET("", handler.GetListProductCustomer)
}
