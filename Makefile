CURRENT_BIN_DIR=${SRC_DIR}/bin
BIN_DIR := ${SRC_DIR}/bin
VERSION := ${SERVER_VERSION}

default: build

all: build

## server build
server:
	go build -a -v -o ${BIN_DIR}/${SERVER_NAME} ${SRC_DIR}/src/server/*.go

build: bins server

## build clean
clean:
	rm -rf ${CURRENT_BIN_DIR}

## Create output folder
bins:
	mkdir ${BIN_DIR}

## check go style and static analysis
lint:
	golint ./src/...;go vet ./src/...

## go test
test:

## show help
help:
	@make2help $(MAKEFILE_LIST)

## define build target not a file.
.PHONY: all build clean test lint help