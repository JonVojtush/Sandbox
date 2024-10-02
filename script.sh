#!/bin/bash
greet() {
    local name="$1"
    echo "Hello ${name}"
}

default_name="World"

if [ -n "$1" ]; then
    greet "$1"
else
    greet "$default_name"
fi
