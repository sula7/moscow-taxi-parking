package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/sula7/moscow-taxi-parking/models"
)

var (
	incorrectDataFormat = "Incorrect data format, must be Content-type Application/JSON"
)

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func (v1 V1) parkingInfoByID(c echo.Context) error {
	parking, err := v1.store.GetParkingById(c.Param("id"))

	if err != nil {
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
	page, err := strconv.ParseInt(qParamPage, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, incorrectDataFormat)
	}

	qParamPerPage := c.QueryParam("per_page")
	perPage, err := strconv.ParseInt(qParamPerPage, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, incorrectDataFormat)
	}

	if page == 0 && perPage == 0 {
		page = 1
		perPage = 5
	}

	request := models.Request{}
	err = c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, incorrectDataFormat)
	}

	parking, err := v1.store.GetParkingsByMode(request.Mode, page, perPage)
	if err != nil {
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
