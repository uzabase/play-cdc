#!/bin/sh

OUTDIR=$(pwd)/gauge_messages

pushd gauge-proto

PATH=$PATH:$GOPATH/bin protoc --go_out=$OUTDIR --go_opt=paths=source_relative --go-grpc_out=$OUTDIR --go-grpc_opt=paths=source_relative spec.proto messages.proto services.proto

popd
