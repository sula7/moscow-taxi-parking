package v1

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/sula7/moscow-taxi-parking/models"
)

var (
	incorrectDataFormat = "incorrect data format"
	noInfoOnThisID      = "No info on this ID. Check again, please"
)

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func (v1 V1) parkingInfoByID(c echo.Context) error {
	parking, err := v1.store.GetParkingById(c.Param("id"))

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusInternalServerError, models.Response{
				Success: false,
				Message: noInfoOnThisID,
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
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusInternalServerError, models.Response{
				Success: false,
				Message: noInfoOnThisID,
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

func (v1 V1) parkingInfoByMode(c echo.Context) error {
	qParamPage := c.QueryParam("page")
	page, err := strconv.Atoi(qParamPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, incorrectDataFormat)
	}

	qParamPerPage := c.QueryParam("per_page")
	perPage, err := strconv.Atoi(qParamPerPage)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, incorrectDataFormat)
	}

	request := models.Request{}
	err = c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, incorrectDataFormat)
	}

	parking, err := v1.store.GetParkingsByMode(request.Mode, page, perPage)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusInternalServerError, models.Response{
				Success: false,
				Message: noInfoOnThisID,
			})
		}
		return c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.ListResponse{
		Success:  true,
		Message:  "OK",
		Parkings: parking,
	})
}
