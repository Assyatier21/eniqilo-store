package router

import (
	"github.com/backend-magang/eniqilo-store/internal/handler/api"
	_ "github.com/backend-magang/eniqilo-store/middleware"
	"github.com/labstack/echo/v4"
)

func InitRouter(server *echo.Echo, handler api.Handler) {
	InitCatRouter(server, handler)
	InitUserRouter(server, handler)
}
