package main

import (
	"github.com/labstack/echo/v4"
	"github.com/underflow101/src/api"
	"github.com/underflow101/src/db"
)

var (
	GDb *db.Db
)

func main() {
	GDb = db.NewDatabase()

	e := echo.New()

	e.GET("/api/v1/gps", api.GetGps)
	e.POST("/api/v1/gps", api.PostGps)

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
