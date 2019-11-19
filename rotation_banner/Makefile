GOCMD=go
GOTEST=$(GOCMD) test

test:
	test_ucb
	test_postgres

test_ucb:
	CGO_ENABLED=0  $(GOTEST) -v ./pkg/ucb1/

test_postgres:
	CGO_ENABLED=0  $(GOTEST) -v ./internal/databases/postgres/

run_grpc_server: go run main.go grpc_server

run_rest_api_server: go run main.go rest_api_server

generate_proto:
	protoc  --go_out=plugins=grpc,paths=source_relative:. proto/banner/banner.proto
	protoc  --go_out=plugins=grpc,paths=source_relative:. proto/server/server.proto
	protoc  --go_out=plugins=grpc,paths=source_relative:. proto/slot/slot.proto
	protoc  --go_out=plugins=grpc,paths=source_relative:. proto/soc_dem_group/soc_dem_group.proto