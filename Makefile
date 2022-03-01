.PHONY:
.SILENT:
.DEFAULT_GOAL := run

certs:
	openssl genrsa -out data/certs/id_rsa 4096
	openssl rsa -in data/certs/id_rsa -pubout -out data/certs/id_rsa.pub
run:
	go run ./cmd/app/main.go
swag:
	swag init -g internal/app/app.go
proto generate:
	protoc -I api/proto --go_out=pkg/api --go_opt=paths=source_relative \
    --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative \
    api/proto/auth.proto