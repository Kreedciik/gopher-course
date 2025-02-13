#!/usr/bin/env bash

for proto_file in ./protos/*.proto; do
    filename=$(basename "$proto_file" .proto)
    mkdir -p "./api-gateway/grpc_gen/$filename"

    protoc -I="./protos" \
        --go_out=paths=source_relative:"./api-gateway/grpc_gen/$filename" \
        --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:"./api-gateway/grpc_gen/$filename" \
        "./protos/$filename.proto"
done

for proto_file in ./protos/*.proto; do
    filename=$(basename "$proto_file" .proto)
    mkdir -p "./auth-service/grpc_gen/$filename"
    
    protoc -I="./protos" \
        --go_out=paths=source_relative:"./auth-service/grpc_gen/$filename" \
        --go-grpc_out=paths=source_relative,require_unimplemented_servers=false:"./auth-service/grpc_gen/$filename" \
        "./protos/$filename.proto"
done