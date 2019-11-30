GOCMD=go
GOTEST=$(GOCMD) test
GOTESTFLAGS=-v

# Test
test_ucb:
	CGO_ENABLED=0  $(GOTEST) $(GOTESTFLAGS) ./pkg/ucb1/

test_postgres:
	CGO_ENABLED=0  $(GOTEST) $(GOTESTFLAGS) ./internal/databases/postgres/

test_api_grpc:
	CGO_ENABLED=0  $(GOTEST) $(GOTESTFLAGS) ./internal/api/grpc

test_api_rest:
	CGO_ENABLED=0  $(GOTEST) $(GOTESTFLAGS) ./internal/api/rest

test: test_ucb test_postgres test_api_grpc test_api_rest

check_govet:
	CGO_ENABLED=0 go vet ./...

check_golint:
	CGO_ENABLED=0 golint ./...

check: check_golint check_govet

run_grpc_server:
	go run main.go grpc_server

run_rest_api_server:
	go run main.go rest_api_server

run:
	go run main.go

generate_proto:
	protoc  --go_out=plugins=grpc,paths=source_relative:. proto/banner/banner.proto
	protoc  --go_out=plugins=grpc,paths=source_relative:. proto/server/server.proto
	protoc  --go_out=plugins=grpc,paths=source_relative:. proto/slot/slot.proto
	protoc  --go_out=plugins=grpc,paths=source_relative:. proto/soc_dem_group/soc_dem_group.proto