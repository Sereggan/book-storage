run:
	go run cmd/grpc/server/main.go

test:
	go test -v -race -timeout 30s ./...

generate:
	protoc --go-grpc_out=./internal/delivery/grpc --go_out=./internal/delivery/grpc ./internal/delivery/grpc/book-storage.proto&&go generate ./...

run-database:
	docker-compose up -d

run-test-client:
	go run cmd/grpc/client/main.go

.DEFAULT_GOAL := run
