#!/bin/bash

rm -f app
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/web/*
docker build -t nyt-api:latest .
#docker push ...
rm -f app