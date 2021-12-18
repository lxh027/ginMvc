.PHONY: all build clean run check

BIN_FILE=mvc

all: check build

build:
	@go mod tidy
	@go build -o "${BIN_FILE}" ./cmd/mvc/main.go

check:
	@go fmt ./...
	@go vet ./...

run:
	./"${BIN_FILE}"
