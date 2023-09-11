#!/bin/bash -e

PROJECT_NAME="PersonalPlanner"

golangci-lint run ./...
go run $PROJECT_NAME "$1" "$2"
