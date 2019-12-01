package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"

	"github.com/sula7/moscow-taxi-parking/models"
)

type (
	ParkingStorage interface {
		Close()

		CreateParkings(parkings *models.Parkings) error
		GetParkingByIdRU(parkingID string) (parking models.Parking, err error)
	}

	Storage struct {
		ParkingStorage
		db *sqlx.DB
	}
)

func New(connString string) (*Storage, error) {
	dbConn, err := sqlx.Connect("mysql", connString)
	return &Storage{db: dbConn}, err
}

func (s *Storage) Close() {
	_ = s.db.Close()
}

func Migrate(connString string) error {
	m, err := migrate.New("file://migrations/", connString)
	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil {
		if err.Error() == "no change" {
			return nil
		}

		return err
	}

	version, _, err := m.Version()
	if err != nil {
		return err
	}

	fmt.Println("Migration success. Current version: ", version)
	return nil
}

func (s *Storage) CreateParkings(parkings *models.Parkings) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	for _, v := range parkings.Parking {
		var id int64
		err = s.db.Get(&id, `SELECT id FROM parkings.moscow WHERE id_ru = ?`, v.IdRU)

		if err == sql.ErrNoRows {
			_, err = tx.Exec(`INSERT INTO parkings.moscow
    (global_id,
     system_object_id,
     id_ru,
     name,
     adm_area,
     district,
     address,
     lon,
     lat,
     car_capacity,
     mode,
     id_en,
     name_en,
     adm_area_en,
     district_en,
     address_en,
     lon_en,
     lat_en,
     car_capacity_en,
     mode_en)
     VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				v.GlobalID,
				v.SystemObjectID,
				v.IdRU,
				v.Name,
				v.AdmArea,
				v.District,
				v.Address,
				v.LongitudeWGS84,
				v.LatitudeWGS84,
				v.CarCapacity,
				v.Mode,
				v.IDEn,
				v.NameEn,
				v.AdmAreaEn,
				v.DistrictEn,
				v.AddressEn,
				v.LongitudeWGS84En,
				v.LatitudeWGS84En,
				v.CarCapacityEn,
				v.ModeEn)
			if err != nil {
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetParkingByIdRU(parkingID string) (parking models.Parking, err error) {
	err = s.db.Get(&parking, `SELECT
       id_ru,
       global_id,
       system_object_id,
       name, adm_area,
       district,
       address,
       lon,
       lat,
       car_capacity,
       mode,
       id_en,
       name_en,
       adm_area_en,
       district_en,
       address_en,
       lon_en,
       lat_en,
       car_capacity_en,
       mode_en FROM parkings.moscow WHERE id_ru = ?`, parkingID)
	return parking, err
}
