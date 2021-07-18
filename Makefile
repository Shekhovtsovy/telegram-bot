.PHONY: build run help

BIN_NAME ?= oa
BIN_FULL_PATH ?= "../bin/${BIN_NAME}"
SOURCE_FULL_PATH ?= "./cmd/${BIN_NAME}"

build:
	@echo "Building the binary..."
	@GOOS=linux GOARCH=amd64 go build -o ${BIN_FULL_PATH} ${SOURCE_FULL_PATH}
	@echo "You can now use ${BIN_FULL_PATH} or 'make run'"

run:
	@echo "Running service..."
	@${BIN_FULL_PATH}
