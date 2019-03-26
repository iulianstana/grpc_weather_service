package main

import (
	"context"
	"fmt"
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

func (s *server) Get(ctx context.Context, in *pb.WeatherRequest) (*pb.WeatherResponse, error) {
	log.Printf("Received: %v", in.City)

	locations, err := s.weatherDriver.GetLocations(in.City)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(locations)
	firstLocation := locations[0]

	currentConditions, err := s.weatherDriver.GetCurrentConditions(firstLocation.Key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentConditions)
	firstCurrentCondition := currentConditions[0]

	weatherResponse := &pb.WeatherResponse{
		Temperature: firstCurrentCondition.Temperature.Metric.Value,
		City:        firstLocation.LocalizedName,
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
