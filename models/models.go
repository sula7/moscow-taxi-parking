package models

type (
	Response struct {
		Success bool              `json:"success"`
		Message string            `json:"message"`
		Parking map[string]string `json:"parking"`
	}

	ListResponse struct {
		Success  bool                `json:"success"`
		Message  string              `json:"message"`
		Parkings []map[string]string `json:"parking"`
	}

	Request struct {
		Mode string `json:"mode"`
	}

	Parkings struct {
		Parking []Parking
	}

	Parking struct {
		Id               int64  `json:"ID,omitempty"form:"id"`
		GlobalID         int64  `json:"global_id,omitempty"form:"global_id"`
		SystemObjectID   string `json:"system_object_id,omitempty"form:"system_object_id"`
		Name             string `json:"Name,omitempty"form:"name"`
		AdmArea          string `json:"AdmArea,omitempty"form:"adm_area"`
		District         string `json:"District,omitempty"form:"district"`
		Address          string `json:"Address,omitempty"form:"address"`
		LongitudeWGS84   string `json:"Longitude_WGS84,omitempty"form:"lon"`
		LatitudeWGS84    string `json:"Latitude_WGS84,omitempty"form:"lat"`
		CarCapacity      int64  `json:"CarCapacity,omitempty"form:"car_capacity"`
		Mode             string `json:"Mode,omitempty"form:"mode"`
		IDEn             int64  `json:"ID_en,omitempty"form:"id_en"`
		NameEn           string `json:"Name_en,omitempty"form:"name_en"`
		AdmAreaEn        string `json:"AdmArea_en,omitempty"form:"adm_area_en"`
		DistrictEn       string `json:"District_en,omitempty"form:"district_en"`
		AddressEn        string `json:"Address_en,omitempty"form:"address_en"`
		LongitudeWGS84En string `json:"Longitude_WGS84_en,omitempty"form:"lon_en"`
		LatitudeWGS84En  string `json:"Latitude_WGS84_en,omitempty"form:"lat_en"`
		CarCapacityEn    int64  `json:"CarCapacity_en,omitempty"form:"car_capacity_en"`
		ModeEn           string `json:"Mode_en,omitempty"form:"mode_en"`
	}
)
