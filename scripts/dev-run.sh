#!/bin/bash -e

function WorkDirectory {
    dir=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
    echo "$dir/.."
}

workDir=$(WorkDirectory)
cd "$workDir"

golangci-lint run ./...
go run ./... "$1" "$2"
