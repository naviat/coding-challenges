#!/bin/bash

# Path to your JSON parser executable
PARSER="./"

# Go to the test directory and loop over your test files
cd tests/step1
for testfile in *.json; do
    echo "Testing $testfile..."
    cat $testfile | go run $PARSER
    echo "Exit code: $?"
done
