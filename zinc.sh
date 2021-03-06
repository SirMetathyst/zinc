#!/bin/bash

ZINC="zinc component add -p kit -n"
FLOAT32_2="-d x:float32 -d y:float32"


_rm() {
    if [ -f "$1" ]; then
        rm $1
    fi
}

test() {
    echo "Running Go Tests"
    go test -race -coverprofile=coverage.txt -covermode=atomic
    go tool cover -html=coverage.txt -o ./coverage.html
}

kit_clean() {
    cd ./kit
    echo "Cleaning Kit Project ..."
    find . -name "*.go" -type f
    read -n 1 -p "... Files will be deleted. Proceed? [y/n]" choice
    echo ""
    case "$choice" in 
    y|Y ) find . -name "*.go" -type f -delete;;
    n|N ) echo "Skipping ...";;
    * ) ;;
    esac
}

kit_gen() {
    kit_clean
    echo "Generating Kit ..."
    ${ZINC} LocalPosition2 ${FLOAT32_2}
	${ZINC} LocalRotation2 ${FLOAT32_2}
	${ZINC} LocalScale2    ${FLOAT32_2}
	${ZINC} Velocity2      ${FLOAT32_2}
}

bench() {
    cd ./benchmark
    echo "Running Benchmarks ..."
    go test -bench . -benchmem -benchtime 15s -cpu 1,4 -count 2
}

clean() {
    echo "Cleaning Project ..."
	_rm ./coverage.html
	_rm ./coverage.txt
}

install() {
    go install ./cmd/zinc
}

$@