#!/bin/bash

# Ensure we're in the project root
cd "$(dirname "$0")/.."

# Generate gRPC code
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/order.proto

# Generate GraphQL code
go run github.com/99designs/gqlgen generate 