.DEFAULT_GOAL := build

fmt:
	goimports -l -w .

.PHONY: fmt

lint:fmt
	golangci-lint run

build:lint
	go mod tidy
	go build main.go
