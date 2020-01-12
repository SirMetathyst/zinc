#!/bin/sh

zinc="zinc component add -p kit -n"
float32_2="-d x:float32 -d y:float32"

_rm() {
    if [ -f "$1" ]; then
        rm $1
    fi
}

clean() {
    echo "Cleaning Project ..."
    find . -name "*.go" -type f -delete
}

gen() {
    echo "Generating Kit ..."
    ${zinc} LocalPosition2 ${float32_2}
	${zinc} LocalRotation2 ${float32_2}
	${zinc} LocalScale2    ${float32_2}
	${zinc} Velocity2      ${float32_2}
}

$@