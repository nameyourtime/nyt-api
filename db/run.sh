#!/bin/bash

docker run --name nameyourtime-db \
  -e POSTGRES_DB=nameyourtime \
  -e POSTGRES_USER=nameyourtime \
  -e POSTGRES_PASSWORD=nameyourtime \
  -p 5432:5432 \
  -v /tmp/nameyourtime-db:/var/lib/postgresql/data \
  nameyourtime-db:latest