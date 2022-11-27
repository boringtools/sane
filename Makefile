SHELL := /bin/bash
GITCOMMIT := $(shell git rev-parse HEAD)
VERSION := 0.0.1-alpha

GO := go

all: clean build

.PHONY: clean
clean:
	-rm sane

GO_CFLAGS=-X main.GITCOMMIT=$(GITCOMMIT) -X main.VERSION=$(VERSION)
GO_LDFLAGS=-ldflags "-w $(GO_CFLAGS)"

.PHONY: build
build:
	$(GO) build ${GO_LDFLAGS}

.PHONY: test
test:
	$(GO) test ./...
