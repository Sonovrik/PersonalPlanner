#!/bin/bash -e

PROJECT_NAME="PersonalPlanner"

function WorkDirectory {
    dir=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
    echo "$dir/.."
}

workDir=$(WorkDirectory)
cd "$workDir"

golangci-lint run ./...

CGO_ENABLED=0 go build -o "$PROJECT_NAME" -v -ldflags="-X 'main.buildDateTime=$(date)'" "./cmd/$PROJECT_NAME"
./"$PROJECT_NAME"
