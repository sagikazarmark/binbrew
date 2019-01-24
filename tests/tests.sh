#!/bin/bash

cd "$(dirname "$0")"

echo "Running basic tests"
./test.sh

echo "Running vanity tests"
./vanity_test.sh

echo "Running multi tests"
./multi_test.sh

echo "Running no progress tests"
./no_progress_test.sh
