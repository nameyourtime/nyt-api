#!/bin/bash

go build -o app ./cmd/web/*
./app -port=":6000"