## Microservices in Go using GRPC 

This repo represents an example of how to interact between golang and python using microservices and gRPC.

The gRPC server is using a http request to get weather data from <b>accuweather.com</b>. The weather is expose through RPC.

There is also a light web server writen in <b>Flask</b> that expose and endpoint. Here, we do gRPC requests to the Weather service to collect weather data of a specific city.
