package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "weather_service/protos/go/weather"
	weatherAPI "weather_service/weather_service_server/driver"
)

const (
	port = ":50051"
)

// server is used to implement weather.WeatherServiceServer.
type server struct {
	weatherDriver *weatherAPI.AccuweatherAPI
}

func NewServer() *server {
	weatherDriver := weatherAPI.NewAccuweatherAPIDriver()
	return &server{weatherDriver: weatherDriver}
}

func (s *server) Get(ctx context.Context, req *pb.WeatherRequest) (*pb.WeatherResponse, error) {
	log.Printf("Received: %v", req.City)

	// Get location information from Weather API
	locations, err := s.weatherDriver.GetLocations(req.City)
	if err != nil {
		log.Fatal(err)
	}
	firstLocation := locations[0]

	// Get Current Conditions for first received location
	currentConditions, err := s.weatherDriver.GetCurrentConditions(firstLocation.Key)
	if err != nil {
		log.Fatal(err)
	}
	firstCurrentCondition := currentConditions[0]

	weatherResponse := &pb.WeatherResponse{
		Temperature: firstCurrentCondition.Temperature.Metric.Value,
		City:        firstLocation.LocalizedName,
		GeoPosition: &pb.GeoPosition{
			Latitude:  firstLocation.GeoPosition.Latitude,
			Longitude: firstLocation.GeoPosition.Longitude,
		},
	}

	return weatherResponse, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serverObj := NewServer()

	s := grpc.NewServer()

	pb.RegisterWeatherServiceServer(s, serverObj)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
