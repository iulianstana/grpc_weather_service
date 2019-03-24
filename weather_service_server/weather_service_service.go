package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "weather_service/protos/go/weather"
)

const (
	port = ":50051"
)

// server is used to implement weather.WeatherServiceServer.
type server struct{}

func (s *server) Get(ctx context.Context, in *pb.WeatherRequest) (*pb.WeatherResponse, error) {
	log.Printf("Received: %v", in.City)
	return &pb.WeatherResponse{Temperature: 30.5}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
