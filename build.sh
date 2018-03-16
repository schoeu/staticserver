#!/bin/sh

# linux 64bit
GOOS=linux GOARCH=amd64 go build -o staticserver_64bit

# linux 32bit
GOOS=linux GOARCH=386 go build -o staticserver_32bit

# windows 64bit
GOOS=windows GOARCH=amd64 go build -o staticserver_64bit.exe

# windows 32bit
GOOS=windows GOARCH=386 go build -o staticserver_32bit.exe

# Mac OS X 64bit
GOOS=darwin GOARCH=amd64 go build -o staticserver_mac