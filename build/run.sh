#!/bin/bash -e

cd "$WORKDIR" || exit

# go updates
#go mod download
go mod tidy

# Linters
golangci-lint run ./...

# Build
CGO_ENABLED=0 go build -o "$PROJECT_NAME" -v -ldflags="-X 'main.buildDateTime=$(date)'" "./cmd/$PROJECT_NAME"

# Start app
./"$PROJECT_NAME"

