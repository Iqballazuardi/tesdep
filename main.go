package main

import (
	"deploy/config"
	"deploy/features/user/handler"
	"deploy/features/user/repository"
	"deploy/features/user/usecase"
	"deploy/routes"
	"deploy/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	dbConn := database.InitDB(*cfg)
	database.Migrate(dbConn)

	repo := repository.New(dbConn)
	srv := usecase.New(repo)
	hdl := handler.New(srv)

	routes.InitRoute(e, hdl)

	if err := e.Start(":8000"); err != nil {
		e.Logger.Fatal("cannot start server ", err.Error())
	}
}
