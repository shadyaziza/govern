#!/bin/bash

GOOS=windows GOARCH=amd64 go build -o bin/govern-amd64.exe main.go

 GOOS=darwin GOARCH=amd64 go build -o bin/govern-amd64-darwin main.go

 GOOS=linux GOARCH=amd64 go build -o bin/govern-amd64-linux main.go