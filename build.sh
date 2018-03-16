#!/bin/sh

# linux 64bit
GOOS=linux GOARCH=amd64 go build -o staticserver

# linux 32bit
GOOS=linux GOARCH=386 go build -o staticserver

# windows 64bit
GOOS=windows GOARCH=amd64 go build -o staticserver.exe

# windows 32bit
GOOS=windows GOARCH=386 go build -o staticserver.exe

# Mac OS X 64bit
GOOS=darwin GOARCH=amd64 go build -o staticserver