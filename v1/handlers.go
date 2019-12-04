package v1

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/sula7/moscow-taxi-parking/models"
)

var (
	incorrectDataFormat = "Incorrect data format, must be Content-type Application/JSON"
)

// ping is use for test DB connection. Must response pong if DB connection is ok
func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

// parkingInfoByID gets all parking info from DB by ID
func (v1 V1) parkingInfoByID(c echo.Context) error {
	parking, err := v1.store.GetParkingById(c.Param("id"))

	if err != nil {
		logrus.Errorln("couldn't get parking info from DB by ID:", err)
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

// parkingInfoByGlobalID gets all parking info from DB by global_id
func (v1 V1) parkingInfoByGlobalID(c echo.Context) error {
	parking, err := v1.store.GetParkingByGlobalId(c.Param("id"))

	if err != nil {
		logrus.Errorln("couldn't get parking info from DB by global ID:", err)
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

// parkingInfoByMode gets all parking info from DB by global_id
func (v1 V1) parkingInfoByMode(c echo.Context) error {
	qParamPage := c.QueryParam("page")
	page, err := strconv.ParseInt(qParamPage, 10, 64)
	if err != nil {
		logrus.Errorln("couldn't parse param page to int:", err)
		return echo.NewHTTPError(http.StatusBadRequest, incorrectDataFormat)
	}

	qParamPerPage := c.QueryParam("per_page")
	perPage, err := strconv.ParseInt(qParamPerPage, 10, 64)
	if err != nil {
		logrus.Errorln("couldn't parse parse per_page to int:", err)
		return echo.NewHTTPError(http.StatusBadRequest, incorrectDataFormat)
	}

	if page == 0 && perPage == 0 {
		page = 1
		perPage = 5
	}

	request := models.Request{}
	err = c.Bind(&request)
	if err != nil {
		logrus.Errorln("couldn't bind request body to struct:", err)
		return echo.NewHTTPError(http.StatusBadRequest, incorrectDataFormat)
	}

	parking, err := v1.store.GetParkingsByMode(request.Mode, page, perPage)
	if err != nil {
		logrus.Errorln("couldn't get parking info from DB by mode:", err)
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
