GO=go
GOOS=linux
GOARCH=amd64
BIN=$(shell pwd)/bin

build:
	GOOS=${GOOS} GOARCH=${GOARCH} ${GO} build -o ${BIN}/lierredefenders ./cmd/lierredefenders
