package models

type Locations struct {
	Key           string            `json:"Key"`
	LocalizedName string            `json:"LocalizedName"`
	GeoPosition   GeoPositionStruct `json:"GeoPosition"`
}

type GeoPositionStruct struct {
	Latitude  float32 `json:"Latitude"`
	Longitude float32 `json:"Longitude"`
}
