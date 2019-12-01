package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewAPI() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", ping)

	groupRU := e.Group("/api/v1/ru")
	groupRU.GET("/parking/:id", parkingRU)

	groupEN := e.Group("/api/v1/en")
	groupEN.GET("/parking/:id", parkingEN)

	e.Logger.Fatal(e.Start(":8080"))
}
