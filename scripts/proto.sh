#!/bin/bash

# Определите переменную для выходной директории
OUT_DIR="../shared/proto/out2"

protoc -I../shared/proto \
    --go_out="$OUT_DIR" --go_opt=paths=source_relative \
    --go-grpc_out="$OUT_DIR" --go-grpc_opt=paths=source_relative \
    --graphql_out="$OUT_DIR" --graphql_opt=paths=source_relative \
    auth.proto


# protoc -I. --go_out=./out --go_opt=paths=source_relative --go-grpc_out=./out --go-grpc_opt=paths=source_relative --graphql_out=./out --graphql_opt=paths=source_relative auth.proto   