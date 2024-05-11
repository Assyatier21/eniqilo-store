package main

import (
	"github.com/backend-magang/eniqilo-store/config"
	"github.com/backend-magang/eniqilo-store/driver"
	"github.com/backend-magang/eniqilo-store/internal/handler/api"
	"github.com/backend-magang/eniqilo-store/internal/repository/postgres"
	"github.com/backend-magang/eniqilo-store/internal/usecase"
	"github.com/backend-magang/eniqilo-store/middleware"
	"github.com/backend-magang/eniqilo-store/router"
	"github.com/backend-magang/eniqilo-store/utils/pkg"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// @title           Swagger Backend Magang - Project 2
// @version         1.0
// @description     This is a documentation of Backend Magang - Project 2
func main() {
	server := echo.New()

	// Load Config
	cfg := config.Load()
	logger := logrus.New()

	dbClient := driver.InitPostgres(cfg)

	// Set Transaction
	sqlTrx := pkg.NewSqlWithTransactionService(dbClient)
	cfg.SqlTrx = sqlTrx

	postgresRepository := postgres.NewRepository(cfg, dbClient, logger)
	usecase := usecase.NewUsecase(cfg, logger, postgresRepository)
	handler := api.NewHandler(logger, usecase)

	router.InitRouter(server, handler)
	middleware.InitMiddlewares(server)

	server.Start(":8080")
}
