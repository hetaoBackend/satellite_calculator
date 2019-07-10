#!/bin/bash
echo "go build ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build && \
mkdir -p tmp/ && \
cp -a views tmp/views && \
mkdir -p tmp/conf && cp conf/app_docker.conf tmp/conf/app.conf && \
cp -a static tmp/static && \
cp -a swagger tmp/swagger && \
cp -a dockerfiles tmp/dockerfiles && \
cp satellite_calculator tmp/ && \
cd tmp && \

docker build -t hetaobackend7/satellite_calculator -f dockerfiles/dockerfile_restful .
cd ..
rm -rf tmp