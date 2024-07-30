proto-gen:
	@echo "Generating Protobuf files..."
	@protoc \
		--proto_path=proto \
		--proto_path=$(GOPATH)/pkg/mod/github.com/cosmos/cosmos-sdk@v0.50.6/proto \
		--proto_path=$(GOPATH)/pkg/mod/github.com/gogo/protobuf@v1.3.2/proto \
		--go_out=. \
		--go-grpc_out=require_unimplemented_servers=false:. \
		proto/identity/identity/tx.proto

.PHONY: proto-gen
