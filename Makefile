# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=kori
BINARY_UNIX=$(BINARY_NAME)_unix

# Build parameters
BUILD_DIR=build
MAIN_PATH=cmd/main.go

.PHONY: all build test clean run deps dev

all: test build

build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v $(MAIN_PATH)

test:
	$(GOTEST) -v ./...

clean:
	rm -rf $(BUILD_DIR)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	nodemon -w . -e go,yaml,toml --exec "make build"

deps:
	$(GOMOD) download
	$(GOMOD) verify

docs:
	@echo "Generating docs..."
	@swag init -g internal/*
	
# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_UNIX) -v $(MAIN_PATH)

dev:
	nodemon

helper:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_UNIX) -v cmd/helper/main.go
