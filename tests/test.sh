#!/bin/bash

cd "$(dirname "$0")"

rm -rf bin/

echo "Installing gobuffalo/packr 1.21.9"
../build/binbrew install gobuffalo/packr@1.21.9

echo "Installing golang/dep 0.5.0"
../build/binbrew install golang/dep@0.5.0

echo "Installing golangci/golangci-lint 1.12.5"
../build/binbrew install golangci/golangci-lint@1.12.5

echo "Installing google/protobuf 3.6.1"
../build/binbrew install google/protobuf@3.6.1

echo "Installing goph/licensei 0.0.7"
../build/binbrew install goph/licensei@0.0.7

echo "Installing goreleaser/goreleaser 0.95.2"
../build/binbrew install goreleaser/goreleaser@0.95.2

echo "Installing gotestyourself/gotestsum 0.3.2"
../build/binbrew install gotestyourself/gotestsum@0.3.2

echo "Installing golang-migrate/migrate 4.2.1"
../build/binbrew install golang-migrate/migrate@4.2.1

echo "Installing stedolan/jq 1.6"
../build/binbrew install stedolan/jq@1.6

echo "Installing stedolan/jq 1.4"
../build/binbrew install stedolan/jq@1.4
