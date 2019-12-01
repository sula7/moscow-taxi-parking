package models

type (
	Response struct {
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Parking Parking `json:"parking"`
	}

	Parkings struct {
		Parking []Parking
	}

	Parking struct {
		IdRU             int64  `json:"ID,omitempty"db:"id_ru"`
		GlobalID         int64  `json:"global_id,omitempty"db:"global_id"`
		SystemObjectID   string `json:"system_object_id,omitempty"db:"system_object_id"`
		Name             string `json:"Name,omitempty"db:"name"`
		AdmArea          string `json:"AdmArea,omitempty"db:"adm_area"`
		District         string `json:"District,omitempty"db:"district"`
		Address          string `json:"Address,omitempty"db:"address"`
		LongitudeWGS84   string `json:"Longitude_WGS84,omitempty"db:"lon"`
		LatitudeWGS84    string `json:"Latitude_WGS84,omitempty"db:"lat"`
		CarCapacity      int64  `json:"CarCapacity,omitempty"db:"car_capacity"`
		Mode             string `json:"Mode,omitempty"db:"mode"`
		IDEn             int64  `json:"ID_en,omitempty"db:"id_en"`
		NameEn           string `json:"Name_en,omitempty"db:"name_en"`
		AdmAreaEn        string `json:"AdmArea_en,omitempty"db:"adm_area_en"`
		DistrictEn       string `json:"District_en,omitempty"db:"district_en"`
		AddressEn        string `json:"Address_en,omitempty"db:"address_en"`
		LongitudeWGS84En string `json:"Longitude_WGS84_en,omitempty"db:"lon_en"`
		LatitudeWGS84En  string `json:"Latitude_WGS84_en,omitempty"db:"lat_en"`
		CarCapacityEn    int64  `json:"CarCapacity_en,omitempty"db:"car_capacity_en"`
		ModeEn           string `json:"Mode_en,omitempty"db:"mode_en"`
	}
)
