test_ucb:
    CGO_ENABLED=0 go test ./pkg/ucb1/

generate_proto:
	protoc proto/banner/banner.proto --go_out=plugins=grpc:.
	protoc proto/server/server.proto --go_out=plugins=grpc:.
	protoc proto/slot/slot.proto --go_out=plugins=grpc:.
	protoc proto/soc_dem_group/soc_dem_group.proto --go_out=plugins=grpc:.


.PHONY: test_ucb generate_proto