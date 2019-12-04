package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sula7/moscow-taxi-parking/config"
	"github.com/sula7/moscow-taxi-parking/models"
)

func setupTest(t *testing.T) *Storage {
	cfg, err := config.Get()
	require.NoError(t, err)
	store, err := New(cfg.DSN, cfg.DBPwd)
	require.NoError(t, err)
	return store
}

func TestCreateParking(t *testing.T) {
	store := setupTest(t)
	defer store.Close()
	modelParking := models.Parking{
		Id:             111,
		GlobalID:       222,
		SystemObjectID: "333",
		Name:           "foo",
		IDEn:           111,
		NameEn:         "goo",
	}

	assert.NoError(t, store.CreateParking(modelParking))

	parking, err := store.GetParkingById("111")
	assert.NoError(t, err)
	assert.Equal(t, modelParking.Name, parking["name"])

	assert.NoError(t, store.RemoveParkingByID("111"))
}

func TestCreateParkings(t *testing.T) {
	store := setupTest(t)
	defer store.Close()
	modelParkings := models.Parkings{
		Parking: []models.Parking{
			{
				Id:             111,
				GlobalID:       222,
				SystemObjectID: "333",
				Name:           "foo",
				IDEn:           111,
				NameEn:         "goo",
			},
			{
				Id:             999,
				GlobalID:       888,
				SystemObjectID: "777",
				Name:           "koo",
				IDEn:           999,
				NameEn:         "noo",
			},
		},
	}

	assert.NoError(t, store.CreateParkings(modelParkings))

	parking, err := store.GetParkingById("111")
	assert.NoError(t, err)
	assert.Equal(t, modelParkings.Parking[0].Name, parking["name"])

	parking, err = store.GetParkingById("999")
	assert.NoError(t, err)
	assert.Equal(t, modelParkings.Parking[1].Name, parking["name"])

	assert.NoError(t, store.RemoveParkingByID("111"))
	assert.NoError(t, store.RemoveParkingByID("999"))
}

func TestGetParkingByGlobalId(t *testing.T) {
	store := setupTest(t)
	defer store.Close()
	modelParking := models.Parking{
		Id:             111,
		GlobalID:       222,
		SystemObjectID: "333",
		Name:           "foo",
		IDEn:           111,
		NameEn:         "goo",
	}

	assert.NoError(t, store.CreateParking(modelParking))

	parking, err := store.GetParkingByGlobalId("222")
	assert.NoError(t, err)

	assert.Equal(t, modelParking.Name, parking["name"])

	assert.NoError(t, store.RemoveParkingByID("111"))
}

func TestGetParkingsByMode(t *testing.T) {
	store := setupTest(t)
	defer store.Close()
	modelParking := models.Parking{
		Id:             111,
		GlobalID:       222,
		SystemObjectID: "333",
		Name:           "foo",
		IDEn:           111,
		NameEn:         "goo",
		Mode:           "zoo",
	}

	assert.NoError(t, store.CreateParking(modelParking))

	parking, err := store.GetParkingsByMode("zoo", 1, 5)
	assert.NoError(t, err)

	for _, v := range parking {
		assert.Equal(t, modelParking.Name, v["name"])
	}

	assert.NoError(t, store.RemoveParkingByID("111"))
}
