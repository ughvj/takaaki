@echo off
set GOOS=linux
set GOARCH=amd64
go build -o out
echo Done building for Linux