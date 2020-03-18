#!/bin/bash

ZINC="zinc -package kit -name"
FLOAT32_2="-var x:float32 -var y:float32"

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

component_clean() {
    cd ./kit
    echo "Cleaning Components ..."
    find . -name "*.go" -type f
    read -n 1 -p "... Files will be deleted. Proceed? [y/n]" choice
    echo ""
    case "$choice" in 
    y|Y ) find . -name "*.go" -type f -delete;;
    n|N ) echo "Skipping ...";;
    * ) ;;
    esac
}

component_gen() {
    component_clean
    echo "Generating Components ..."
    # MultiComponentDataTemplate
    ${ZINC} LocalPosition2 ${FLOAT32_2}
	${ZINC} LocalRotation2 ${FLOAT32_2}
	${ZINC} LocalScale2    ${FLOAT32_2}
	${ZINC} Velocity2      ${FLOAT32_2}
    # SingleComponentDataTemplate
    ${ZINC} UnixTime -var time:int
    # FlagComponentDataTemplate
    ${ZINC} Active
    # UniqueFlagComponentDataTemplate
    ${ZINC} Running -unique
    # UniqueSingleComponentDataTemplate
    ${ZINC} LogService -unique -var logger:interface{}
    # UniqueMultiComponentDataTemplate
    ${ZINC} TimeService -unique -var min:int -var sec:int
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