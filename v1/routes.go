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

	g := e.Group("/api/v1")
	g.GET("/parking/:id", parking)

	e.Logger.Fatal(e.Start(":8080"))
}
