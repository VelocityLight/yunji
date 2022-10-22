# Variables contain directory/file and so on
GO_BUILD_DIR = ./bin/
WEBSITE_DIR = ./website/

GO_BINARY = ${GO_BUILD_DIR}/yunji
WEBSITE_BINARY = ${WEBSITE_DIR}/build/
YUNJI_MAIN_FILE = ./cmd/yunji/*.go

build: build.web build.server

clean:
	@rm -rf ${GO_BUILD_DIR}
	@rm -rf ${WEBSITE_BINARY}
	@echo "clear all temporary files and folders successful hahaha!"

run: clean build
	./${GO_BINARY}

build.server:
	go fmt ./... && \
	go build -o ${GO_BINARY} ${YUNJI_MAIN_FILE}

build.web:
	cd ${WEBSITE_DIR} && \
	npm ci && \
	npm run build && \
	cd ..

fmt:
	goimports -l -w -local "github.com/VelocityLight/yunji"  .

.PHONY: build run clean help
.PHONY: build.web build.server docker docker.run k8s k8s.clean
