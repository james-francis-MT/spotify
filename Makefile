.PHONY: run build test lint

run:
	go run ./cmd/textEditor/

build:
	go build -o bin/textEditor ./cmd/textEditor/

test:
	go test ./...

format:
	go fmt ./...
