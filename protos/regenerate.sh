#!/usr/bin/env bash

PROTO_DOCKER_PATH=/opt/weather_service/protos


for proto in weather ; do
    # Generate Python pb files
    OUT_PATH=$PROTO_DOCKER_PATH/python PROTO_PATH=$PROTO_DOCKER_PATH PROTO_NAME=$proto docker-compose up generate-python


    # Generate Go pb files
    mkdir -p go/$proto
    OUT_PATH=$PROTO_DOCKER_PATH/go/$proto PROTO_PATH=$PROTO_DOCKER_PATH PROTO_NAME=$proto docker-compose up generate-go
done