package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func parkingRU(c echo.Context) error {
	return nil
}

func parkingEN(c echo.Context) error {
	return nil
}
