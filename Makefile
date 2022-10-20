# Variables contain directory/file and so on
GO_BUILD_DIR = ./bin/

GO_BINARY = ${GO_BUILD_DIR}/yunji
YUNJI_MAIN_FILE = ./cmd/yunji/*.go

build: build.server

build.server:
	go fmt ./... && \
	go build -o ${GO_BINARY} ${YUNJI_MAIN_FILE}


clean:
	@rm -rf ${GO_BUILD_DIR}
	@echo "clear all temporary files and folders successful hahaha!"

run: clean build
	./${GO_BINARY}

