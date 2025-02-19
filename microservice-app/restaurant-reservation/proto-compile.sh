#!/usr/bin/env bash

for proto_file in ./protos/*.proto; do
    filename=$(basename "$proto_file" .proto)

    # API Gateway
    mkdir -p "./api-gateway/grpc_gen/$filename"
    protoc -I="./protos" \
        --go_out=paths=source_relative:"./api-gateway/grpc_gen/$filename" \
        --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:"./api-gateway/grpc_gen/$filename" \
        "./protos/$filename.proto"

    # Auth Service
    mkdir -p "./auth-service/grpc_gen/$filename"
    protoc -I="./protos" \
        --go_out=paths=source_relative:"./auth-service/grpc_gen/$filename" \
        --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:"./auth-service/grpc_gen/$filename" \
        "./protos/$filename.proto"

    # Reservation Service
    mkdir -p "./reservation-service/grpc_gen/$filename"
    protoc -I="./protos" \
        --go_out=paths=source_relative:"./reservation-service/grpc_gen/$filename" \
        --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:"./reservation-service/grpc_gen/$filename" \
        "./protos/$filename.proto"
done
