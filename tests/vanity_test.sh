#!/bin/bash

cd "$(dirname "$0")"

rm -rf bin/

echo "Installing dep 0.5.0"
../build/binbrew install dep@0.5.0

echo "Installing golangci-lint 1.12.5"
../build/binbrew install golangci-lint@1.12.5

echo "Installing protobuf 3.6.1"
../build/binbrew install protobuf@3.6.1

echo "Installing protoc 3.6.1"
../build/binbrew install protoc@3.6.1

echo "Installing goreleaser 0.95.2"
../build/binbrew install goreleaser@0.95.2

echo "Installing gotestsum 0.3.2"
../build/binbrew install gotestsum@0.3.2

echo "Installing jq 1.6"
../build/binbrew install jq@1.6

echo "Installing protolock 0.10.0"
../build/binbrew install protolock@0.10.0
