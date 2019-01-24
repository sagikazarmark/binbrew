#!/bin/bash

cd "$(dirname "$0")"

rm -rf bin/

echo "Installing dep@0.5.0"
../build/binbrew install --no-progress dep@0.5.0
