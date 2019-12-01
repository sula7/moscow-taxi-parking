package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func parking(c echo.Context) error {
	return nil
}
