package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sula7/moscow-taxi-parking/storage"
)

type (
	V1 struct {
		store *storage.Storage
	}
)

func NewAPI(store *storage.Storage, bindPort string) {
	e := echo.New()

	v1 := &V1{
		store: store,
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", ping)

	g := e.Group("/api/v1")
	g.GET("/parking/id/:id", v1.parkingInfoByID)
	g.GET("/parking/global-id/:id", v1.parkingInfoByGlobalID)
	g.POST("/parking/mode", v1.parkingInfoByMode)

	e.Logger.Fatal(e.Start(bindPort))
}
