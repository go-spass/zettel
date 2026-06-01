.DEFAULT_GOAL := all

BUILD_DIR ?= ./build

fmt:
	go fmt ./...

lint: fmt
	golangci-lint run

vet: lint
	go vet ./...

test:
	go test -v -cover ./...

bench:
	go test -bench . ./...

clean:
	rm -r $(BUILD_DIR)

build: vet
	go build -o build/zettel ./cmd/web

run:
	go run ./cmd/web

all: build

.PHONY: all build run
