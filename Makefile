BIN_DIR := ./bin
SERVER_EXE := nyx-server
BACKEND_DIR := ./backend

.PHONY: init build-cli build-server clean all

init:
	@mkdir -p $(BIN_DIR)

build-server: init
	@echo "Building server interface..."
	@cd $(BACKEND_DIR) && go build -o ../$(BIN_DIR)/$(SERVER_EXE) ./cmd/api

all: build-cli build-server

clean:
	@echo "Cleaning binaries..."
	@rm -rf $(BIN_DIR)/*
	@go clean

run-server: build-server
	@echo "Starting server..."
	@$(BIN_DIR)/$(SERVER_EXE)