SHELL := /bin/bash
GITCOMMIT := $(shell git rev-parse HEAD)
VERSION := "$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)"
BINARY_NAME := sane

GO := go

all: clean build

.PHONY: clean
clean:
	-rm -rf target

GO_CFLAGS=-X main.GITCOMMIT=$(GITCOMMIT) -X main.VERSION=$(VERSION)
GO_LDFLAGS=-ldflags "-w $(GO_CFLAGS)"

.PHONY: build
build: build-linux-amd64 build-linux-arm64

.PHONY: build-linux-amd64
build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GO) build ${GO_LDFLAGS} -o target/linux/amd64/$(BINARY_NAME)

.PHONY: build-linux-arm64
build-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GO) build ${GO_LDFLAGS} -o target/linux/arm64/$(BINARY_NAME)

.PHONY: build-darwin-amd64
build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GO) build ${GO_LDFLAGS} -o target/darwin/amd64/$(BINARY_NAME)

.PHONY: build-darwin-arm64
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(GO) build ${GO_LDFLAGS} -o target/darwin/arm64/$(BINARY_NAME)

.PHONY: test
test:
	$(GO) test ./...
