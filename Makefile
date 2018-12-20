PATH  := $(PATH):$(PWD)
SHELL := env PATH=$(PATH) /bin/bash

default: run
.PHONY: fmt build run

fmt:
	go fmt ./...

build: fmt
	go install github.com/jinmatt/go-ecs-app

run: build
	$(GOPATH)/bin/go-ecs-app