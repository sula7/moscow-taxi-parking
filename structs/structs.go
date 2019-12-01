package structs

type (
	Parkings struct {
		Parking []Parking
	}

	Parking struct {
		GlobalID         int64  `json:"global_id,omitempty"`
		SystemObjectID   string `json:"system_object_id,omitempty"`
		ID               int64  `json:"ID,omitempty"`
		Name             string `json:"Name,omitempty"`
		AdmArea          string `json:"AdmArea,omitempty"`
		District         string `json:"District,omitempty"`
		Address          string `json:"Address,omitempty"`
		LongitudeWGS84   string `json:"Longitude_WGS84,omitempty"`
		LatitudeWGS84    string `json:"Latitude_WGS84,omitempty"`
		CarCapacity      int64  `json:"CarCapacity,omitempty"`
		Mode             string `json:"Mode,omitempty"`
		IDEn             int64  `json:"ID_en,omitempty"`
		NameEn           string `json:"Name_en,omitempty"`
		AdmAreaEn        string `json:"AdmArea_en,omitempty"`
		DistrictEn       string `json:"District_en,omitempty"`
		AddressEn        string `json:"Address_en,omitempty"`
		LongitudeWGS84En string `json:"Longitude_WGS84_en,omitempty"`
		LatitudeWGS84En  string `json:"Latitude_WGS84_en,omitempty"`
		CarCapacityEn    int64  `json:"CarCapacity_en,omitempty"`
		ModeEn           string `json:"Mode_en,omitempty"`
	}
)
