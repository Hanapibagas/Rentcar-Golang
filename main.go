package main

import (
	"BE-REPO-20/app/configs"
	"BE-REPO-20/app/database"
	"BE-REPO-20/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := configs.InitConfig()
	dbsql := database.InitDBMysql(cfg)
	database.InitMigrate(dbsql)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routes.InitRouter(dbsql, e)

	e.Logger.Fatal(e.Start(":8080"))
}
