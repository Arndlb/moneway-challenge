#!/usr/bin/env bash

# Generate protobuf files
protoc --go_out=plugins=grpc:. ./api/v1/transaction.proto
protoc --go_out=plugins=grpc:. ./api/v1/balance.proto

# Get dependencies
dep ensure

# Build and run scylladb database
sudo docker build -t scylladb . &

# Build Balance and Transaction microservices
go build -o balance ./services/balance/main.go
go build -o transaction ./services/transaction/main.go
