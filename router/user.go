package router

import (
	"github.com/backend-magang/eniqilo-store/internal/handler/api"
	"github.com/labstack/echo/v4"
)

func InitUserRouter(e *echo.Echo, handler api.Handler) {
	v1 := e.Group("/v1")
	staff := v1.Group("/staff")

	staff.POST("/register", handler.RegisterStaff)
	// user.POST("/login", handler.LoginUser)
}
