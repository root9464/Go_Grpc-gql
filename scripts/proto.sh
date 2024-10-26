#!/bin/bash

# Определите переменную для выходной директории
OUT_DIR="../shared/proto/out"

protoc -I../shared/proto \
    --go_out="$OUT_DIR" --go_opt=paths=source_relative \
    --go-grpc_out="$OUT_DIR" --go-grpc_opt=paths=source_relative \
    --graphql_out="$OUT_DIR" --graphql_opt=paths=source_relative \
    auth.proto