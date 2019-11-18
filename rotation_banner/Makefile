test_ucb: CGO_ENABLED=0 go test ./pkg/ucb1/

test_postgres:
CGO_ENABLED=0 go test ./internal/databases/postgres/

run_grpc_server: go run main.go grpc_server

run_rest_api_server: go run main.go rest_api_server

generate_proto:
	protoc proto/banner/banner.proto --go_out=plugins=grpc:.
	protoc proto/server/server.proto --go_out=plugins=grpc:.
	protoc proto/slot/slot.proto --go_out=plugins=grpc:.
	protoc proto/soc_dem_group/soc_dem_group.proto --go_out=plugins=grpc:.

test:
    docker-compose

.PHONY: help