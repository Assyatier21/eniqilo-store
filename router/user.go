package router

import (
	"github.com/backend-magang/eniqilo-store/internal/handler/api"
	"github.com/backend-magang/eniqilo-store/middleware"
	"github.com/labstack/echo/v4"
)

func InitUserRouter(e *echo.Echo, handler api.Handler) {
	v1 := e.Group("/v1")
	staff := v1.Group("/staff")
	customer := v1.Group("/customer", middleware.TokenValidationMiddleware())

	staff.POST("/register", handler.RegisterStaff)
	staff.POST("/login", handler.LoginStaff)

	customer.POST("/register", handler.RegisterCustomer)
}
