package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sula7/moscow-taxi-parking/models"
)

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func (v1 V1) parkingInfoByID(c echo.Context) error {
	parking, err := v1.store.GetParkingById(c.Param("id"))

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusInternalServerError, models.Response{
				Success: false,
				Message: "No info on this ID. Check again, please",
			})
		}
		return c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "OK",
		Parking: parking,
	})
}

func (v1 V1) parkingInfoByGlobalID(c echo.Context) error {
	parking, err := v1.store.GetParkingByGlobalId(c.Param("id"))

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusInternalServerError, models.Response{
				Success: false,
				Message: "No info on this global-ID. Check again, please",
			})
		}
		return c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "OK",
		Parking: parking,
	})
}
