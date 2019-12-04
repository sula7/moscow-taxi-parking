package v1

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"

	"github.com/sula7/moscow-taxi-parking/config"
	"github.com/sula7/moscow-taxi-parking/storage"
)

func setupTest(t *testing.T) *V1 {
	cfg, err := config.Get()
	require.NoError(t, err)
	store, err := storage.New(cfg.DSN, cfg.DBPwd)
	require.NoError(t, err)
	v1 := &V1{store: store}
	return v1
}

func TestParkingInfoByID(t *testing.T) {
	v1 := setupTest(t)
	defer v1.store.Close()

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath("api/v1/parking/id/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("161")

	require.NoError(t, v1.parkingInfoByID(ctx))
	require.Equal(t, http.StatusOK, rec.Code)
	require.NotEmpty(t, rec.Body.Bytes())
}

func TestParkingInfoByGlobalID(t *testing.T) {
	v1 := setupTest(t)
	defer v1.store.Close()

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath("api/v1/parking/global-id/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1704691")

	require.NoError(t, v1.parkingInfoByGlobalID(ctx))
	require.Equal(t, http.StatusOK, rec.Code)
	require.NotEmpty(t, rec.Body.Bytes())
}

func TestParkingInfoByMode(t *testing.T) {
	v1 := setupTest(t)
	defer v1.store.Close()

	e := echo.New()

	q := make(url.Values)
	q.Set("page", "1")
	q.Set("per_page", "10")

	req := httptest.NewRequest(http.MethodPost, "/?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.SetPath("api/v1/parking/mode")

	require.NoError(t, v1.parkingInfoByMode(ctx))
	require.Equal(t, http.StatusOK, rec.Code)
	require.NotEmpty(t, rec.Body.Bytes())
}
