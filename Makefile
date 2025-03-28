GO_WORKSPACE := $(GOPATH)/src

protoc:
	protoc --proto_path=protos protos/*.proto --go_out=. --go-grpc_out=.
	@echo "Protoc compiled"