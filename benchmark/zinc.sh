#!/bin/sh

bench() {
    echo "Running Benchmarks ..."
    go test -bench . -benchmem -benchtime 15s -cpu 1,4 -count 2
}

$@