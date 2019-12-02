package storage

import (
	"strconv"

	"github.com/go-redis/redis/v7"

	"github.com/sula7/moscow-taxi-parking/models"
)

type (
	ParkingStorage interface {
		Close()

		CreateParkings(parkings *models.Parkings) error
		GetParkingById(parkingID string) (parking map[string]string, err error)
		GetParkingByGlobalId(globalID string) (parking models.Parking, err error)
		GetParkingsByMode(mode string, page, maxParkingsPerPage int) (parking []models.Parking, err error)
	}

	Storage struct {
		ParkingStorage
		db *redis.Client
	}
)

func New(connString string) *Storage {
	dbConn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &Storage{db: dbConn}
}

func (s *Storage) Close() {
	_ = s.db.Close()
}

func (s *Storage) CreateParkings(parkings *models.Parkings) error {
	pipe := s.db.TxPipeline()

	for _, parking := range parkings.Parking {
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "ID", parking.Id)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "global_id", parking.GlobalID)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "system_object_id", parking.SystemObjectID)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "name", parking.Name)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "adm_area", parking.AdmArea)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "district", parking.District)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "address", parking.Address)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "lon", parking.LongitudeWGS84)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "lat", parking.LatitudeWGS84)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "car_capacity", parking.CarCapacity)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "mode", parking.Mode)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "id_en", parking.IDEn)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "name_en", parking.NameEn)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "adm_area_en", parking.AdmAreaEn)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "district_en", parking.DistrictEn)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "address_en", parking.AddressEn)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "lon_en", parking.LongitudeWGS84En)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "lat_en", parking.LatitudeWGS84En)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "car_capacity_en", parking.CarCapacityEn)
		pipe.HSet(strconv.FormatInt(parking.Id, 10), "mode_en", parking.ModeEn)
	}

	_, err := pipe.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetParkingById(parkingID string) (parking map[string]string, err error) {
	return s.db.HGetAll(parkingID).Result()
}

func (s *Storage) GetParkingByGlobalId(globalID string) (parking models.Parking, err error) {
	return parking, err
}

func (s *Storage) GetParkingsByMode(mode string, page, maxParkingsPerPage int) (parking []models.Parking, err error) {
	return parking, err
}
