#!make

# version
version := "v0.0.1"

# app name
app := "pttoy"

# platform
darwin := ${app}-darwin.${version}.bin
linux := ${app}-linux.${version}.bin

default: darwin

linux:

	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o "${linux}" cmd/main.go
	upx ${linux}

darwin:

	CGO_ENABLE=0 GOOS=darwin GOARCH=amd64 go build -o "${darwin}" cmd/main.go
	upx ${darwin}
