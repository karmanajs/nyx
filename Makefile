BIN_DIR := ./bin
CLI_EXE := nyx-cli
SERVER_EXE := nyx-server

.PHONY: init build-cli build-server clean all

init:
	@mkdir -p $(BIN_DIR)

build-cli: init
	@echo "Building CLI interface..."
	@go build -o $(BIN_DIR)/$(CLI_EXE) ./cmd/cli

build-server: init
	@echo "Building server interface..."
	@go build -o $(BIN_DIR)/$(SERVER_EXE) ./cmd/server

all: build-cli build-server

clean:
	@echo "Cleaning binaries..."
	@rm -rf $(BIN_DIR)/*
	@go clean

run-server: build-server
	@echo "Starting server..."
	@$(BIN_DIR)/$(SERVER_EXE)