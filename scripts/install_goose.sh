#!/bin/bash

cd ../db/goose
go mod tidy
go build -o goose ./cmd/goose