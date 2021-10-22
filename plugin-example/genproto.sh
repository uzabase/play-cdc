#!/bin/sh

pushd ../../gauge-proto

PATH=$PATH:$GOPATH/bin protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative spec.proto messages.proto services.proto

popd
