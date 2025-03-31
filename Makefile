run: build
	@./bin/api

build:
	@go build -o bin/api ./cmd/server

test:
	@go test -v ./...
