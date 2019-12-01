package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sula7/moscow-taxi-parking/storage"
)

type (
	ParkingHandlers interface {
		ping(c echo.Context) error
		parkingInfoByID(c echo.Context) error
	}

	V1 struct {
		ParkingHandlers
		store storage.ParkingStorage
	}
)

func NewAPI(store storage.ParkingStorage) {
	e := echo.New()

	v1 := V1{
		store: store,
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", ping)

	groupRU := e.Group("/api/v1")
	groupRU.GET("/parking/id/:id", v1.parkingInfoByID)
	groupRU.GET("/parking/global-id/:id", v1.parkingInfoByGlobalID)
	groupRU.POST("/parking/mode", v1.parkingInfoByMode)

	e.Logger.Fatal(e.Start(":8080"))
}
