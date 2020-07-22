#!/usr/bin/bash
echo start to build
GOOS=linux GOARCH=arm go build -o ./bin/boss2_ledscreen.ARM main.go
