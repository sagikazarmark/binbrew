#!/bin/bash

cd "$(dirname "$0")"

rm -rf bin/

echo "Installing dep@0.5.0, golangci-lint@1.12.5, protoc@3.6.1"
../build/binbrew install dep@0.5.0 golangci-lint@1.12.5 protoc@3.6.1
