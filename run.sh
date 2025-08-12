#!/bin/bash

set -e

THIS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

function 01_go_mod_init {
    echo "Initializing Go module..."
    cd "${THIS_DIR}/01-fiber"
    go mod init github.com/sunil-ytla/01-fiber
}

function 01_go_mod_tidy {
    echo "Tidying Go module..."
    cd "${THIS_DIR}/01-fiber"
    go mod tidy
}

function 01_go_run {
    echo "Running Go application..."
    cd "${THIS_DIR}/01-fiber"
    go run main.go
}

function 02_go_test {
    echo "running test"
    cd "${THIS_DIR}/01-fiber"
    go test -v
}

function proj_go_mod_init {
    echo "Initializing Go module for project..."
    cd "${THIS_DIR}/APIPROJECT"
    go mod init github.com/sunil-ytla/APIPROJECT
}

function proj_go_mod_tidy {
    echo "Tidying Go module..."
    cd "${THIS_DIR}/APIPROJECT"
    go mod tidy
}

function proj_run {
    echo "Running Go application..."
    cd "${THIS_DIR}/APIPROJECT"
    go run main.go
}

function proj_generate_sqlc {
    echo "Generating SQLC files..."
    cd "${THIS_DIR}/APIPROJECT"
    sqlc generate
}



function help {
    echo "Usage: $0 [options]"
    echo "Options:"
    echo "  --help          Show this help message"
    echo "  --version       Show the script version"
}


TIMEFORMAT="Task completed in %3lR"
time ${@:-help}