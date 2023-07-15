package main

import (
	"accounting-project/initializers"
	"accounting-project/pkg/postgres"
	"accounting-project/pkg/redis"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	postgres.ConnectToDb()
	initializers.SyncDatabase()
	redis.ConnectToRedis()
}
func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"*"},
	}))
	initializers.RouteInit(e.Group(""))

	e.Static("/uploads", "./uploads")
	e.Logger.Fatal(e.Start("0.0.0.0:3535"))
}
