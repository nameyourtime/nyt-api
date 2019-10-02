#!/bin/bash

rm -f app
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/web/*
docker build -t docker.pkg.github.com/nameyourtime/nyt-api/nyt-api:latest .
docker push docker.pkg.github.com/nameyourtime/nyt-api/nyt-api:latest
rm -f app

# remove previous image(s)
docker rmi $(docker images --filter dangling=true -q)