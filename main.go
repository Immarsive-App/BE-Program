package main

import (
	"kelompok1/immersive-dash/app/config"
	"kelompok1/immersive-dash/app/database"
	"kelompok1/immersive-dash/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// logger := helpers.NewLogger()
	// logger.SetFormatter(&logrus.TextFormatter{}) // Atur formatter sesuai kebutuhan
	// logger.SetLevel(logrus.InfoLevel)            // Atur tingkat log sesuai kebutuhan

	cfg := config.InitConfig()
	dbMysql := database.InitDBMysql(cfg)
	database.InitialMigration(dbMysql)

	// create a new echo instance
	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.InitRouter(dbMysql, e)
	//start server and port
	e.Logger.Fatal(e.Start(":8000"))
}
