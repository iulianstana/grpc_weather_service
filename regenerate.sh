#!/usr/bin/env bash

if [ "python" == $1 ]; then
    if [ "code" == $2 ]; then
        docker build --no-cache -f docker_containers/python/dockerfile_webserver -t weather_webserver .
    elif [ "base_image" == $2 ]; then
        docker build --no-cache -f docker_containers/python/base_docker_file -t python_weather_base_image docker_containers/python
    else
        echo "Usage: ./regenerate <python|go> <code|base_image>"
    fi

elif [ "go" == $1 ]; then
    if [ "code" == $2 ]; then
        docker build --no-cache -f docker_containers/go/dockerfile_weather_service -t weather_service .
    elif [ "base_image" == $2 ]; then
        docker build --no-cache -f docker_containers/go/base_docker_file -t go_weather_base_image .
    else
        echo "Usage: ./regenerate <python|go> <code|base_image>"
    fi
else
    echo "Usage: ./regenerate <python|go> <code|base_image>"
fi
