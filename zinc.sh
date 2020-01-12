#!/bin/sh

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

clean() {
    echo "Cleaning Project ..."
	_rm ./coverage.html
	_rm ./coverage.txt
}

install() {
    go install ./cmd/zinc
}



$@