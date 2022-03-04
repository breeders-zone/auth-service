.PHONY:
.SILENT:
.DEFAULT_GOAL := run

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage
test.coverage:
	go tool cover -func=cover.out | grep "total"
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
gen:
	mockgen -source=internal/services/services.go -destination=internal/services/mocks/mock.go
	mockgen -source=pkg/api/auth_grpc.pb.go -destination=pkg/api/mocks/mock.go
	