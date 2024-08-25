GOCMD = go
GORUN = $(GOCMD) run
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTOOL = $(GOCMD) tool
GOMOD = $(GOCMD) mod
GOINSTALL = $(GOCMD) install

WORKDIR=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
OUTPUT_PATH=build
BINARY_NAME=app
BIN_OUTPUT=$(OUTPUT_PATH)/$(BINARY_NAME)

CONFIG_FILE_PATH=internal/configs
OUTPUT_CONFIG_FILES_PATH=$(OUTPUT_PATH)/$(CONFIG_FILE_PATH)
CONFIG_FILES=$(wildcard $(CONFIG_FILE_PATH)/*.yaml)

run:
	@echo "===== Running... ====="
	$(GORUN) .

dev:
	@echo ">> If you have a problem with reflex, please run 'make install' to install it <<"
	@echo "===== Running in development mode... ====="
	@reflex -r '\.go$$' -s -- sh -c 'make run'

install:
	@echo "===== Installing... ====="
	$(GOMOD) download
	$(GOINSTALL) github.com/cespare/reflex@latest

build: clean _build copy_config

clean:
	$(GOCLEAN)
	@rm -rf $(OUTPUT_PATH)

copy_config:
	@mkdir -p $(OUTPUT_CONFIG_FILES_PATH)
	@for file in $(CONFIG_FILES); do \
		cp $$file $(OUTPUT_CONFIG_FILES_PATH); \
	done

_build:
	@echo "===== Building... ====="
	CGO_ENABLED=0 $(GOBUILD) -a -installsuffix cgo -o $(BIN_OUTPUT) main.go
	@echo "===== Built successfully to '$(BIN_OUTPUT)' ====="
	@echo "===== Run './$(BIN_OUTPUT)' to start the application ===="

