package driver

import (
	"fmt"
	"log"
	"weather_service/weather_service_server/driver/models"
)

const (
	accuweatherURL = "http://dataservice.accuweather.com/"
)

type AccuweatherAPI struct {
	apiKey string
}

func NewAccuweatherAPIDriver() *AccuweatherAPI {
	return &AccuweatherAPI{apiKey: "wOWjXa1okNfSUaFGmkNNLEEHcLQAYnWA"}
}

func (api *AccuweatherAPI) GetLocations(query string) ([]models.Locations, error) {
	var locations []models.Locations

	citiesLocationEndpoint := "locations/v1/cities/search"

	// Compose locations url
	locationsUrl := fmt.Sprintf("%s%s?apikey=%s&q=%s", accuweatherURL, citiesLocationEndpoint, api.apiKey, query)

	fmt.Println(locationsUrl)
	err := extractContent(locationsUrl, &locations)
	if err != nil {
		log.Printf("failed to serve: %v", err)
		return []models.Locations{}, err
	}
	return locations, nil
}

func (api *AccuweatherAPI) GetCurrentConditions(cityKey string) ([]models.CurrentConditions, error) {
	var currentConditions []models.CurrentConditions

	citiesLocationEndpoint := "/currentconditions/v1/"

	// Compose currentConditions url
	currentConditionsUrl := fmt.Sprintf("%s%s%s?apikey=%s&&metric=True", accuweatherURL, citiesLocationEndpoint, cityKey, api.apiKey)

	err := extractContent(currentConditionsUrl, &currentConditions)
	if err != nil {
		log.Printf("failed to serve: %v", err)
		return []models.CurrentConditions{}, err
	}
	return currentConditions, nil
}

// API example
//func main() {
//	queryLocation := "Bucuresti"
//
//	weatherAPI := NewAccuweatherAPIDriver()
//
//	// Get locations
//	locations, err := weatherAPI.GetLocations(queryLocation)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(locations)
//	cityKey := locations[0].Key
//
//	currentConditions, err := weatherAPI.GetCurrentConditions(cityKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(currentConditions)
//}
