#!/bin/sh

go get -v -u google.golang.org/grpc
go get -v -u github.com/golang/protobuf/protoc-gen-go
go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger