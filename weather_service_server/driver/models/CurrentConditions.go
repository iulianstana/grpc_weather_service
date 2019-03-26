package models

type CurrentConditions struct {
	LocalTime   string            `json:"LocalObservationDateTime"`
	Temperature TemperatureStruct `json:"Temperature"`
}

type TemperatureStruct struct {
	Metric MetricStruct `json:"Metric"`
}

type MetricStruct struct {
	Value float32 `json:"Value"`
	Unit  string  `json:"Unit"`
}
