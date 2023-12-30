help:
	@echo "Please use \`make <target>' where <target> is one of:"
	@echo "  build         Compiles and builds the application."
	@echo "  run           Runs the application."
	@echo "  test          Run unit tests."

all: build test

build:
	go build -v ./...

run:
	go run .

test:
	go test -v ./... -race -covermode=atomic -coverprofile=coverage.out


.PHONY: help build dev test