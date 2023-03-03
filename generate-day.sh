#!/bin/bash

if [ -z "$1" ]; then
    >&2 echo "No argument for day"
    exit 1
fi

case $1 in
    ''|*[!0-9]*) >&2 echo "Not integer: $1" && exit 1 ;;
    *) ;;
esac

if [ 1 -gt "$1" ] || [ "$1" -gt 25 ]; then
    >&2 echo "Invalid day: $1"
    exit 1
fi

if [ -e "./days/day$1/$1.go" ]; then
    echo "Already exists: ./days/day$1/$1.go"
    exit 0
fi



echo "pass"
