BIN_DIR := ./bin
EXE_NAME := nyx

.PHONY: init build-cli clean

init:
	mkdir -p $(BIN_DIR)

build-cli: init
	go build -o $(BIN_DIR)/$(EXE_NAME) ./cmd/cli

clean:
	rm -rf $(BIN_DIR)/*
	go clean