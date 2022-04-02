build:
	go build -o bin/server cmd/server/main.go

run:
	go run cmd/server/main.go

test:
	go test -v -race -timeout 30s ./...

generate:
	protoc --go-grpc_out=./internal/delivery/grpc --go_out=./internal/delivery/grpc ./internal/delivery/grpc/book-storage.proto

.DEFAULT_GOAL := run
