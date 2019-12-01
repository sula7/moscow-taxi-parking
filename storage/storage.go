package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/sula7/moscow-taxi-parking/structs"
)

func CreateParking(parkings *structs.Parkings) error {
	dbConn, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/parkings")
	if err != nil {
		fmt.Printf("Failed to connect to database:%v\n", err)
	}

	defer dbConn.Close()

	tx, err := dbConn.Begin()
	if err != nil {
		return err
	}

	for _, v := range parkings.Parking {
		var id int64
		err = dbConn.Get(&id, `SELECT id FROM parkings.moscow WHERE id_ru = ?`, v.ID)

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
				v.ID,
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
