#!/usr/bin/env bash

GOARCH=386 go build -ldflags "-H=windowsgui" -o ./bin/boss2_ledscreen.exe main.go