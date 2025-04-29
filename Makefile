PROTO_DIR=schema/proto
OUT_DIR=pkg/adapter/grpcgen

PROTO_FILES=$(wildcard $(PROTO_DIR)/*.proto)

PHONY: install
install:
	make install_protoc
	go install github.com/evilmartians/lefthook@latest

.PHONY:install_protoc
install_protoc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

PHONY:format
format:
	gofmt -s -w .
	goimports -local "github.com/t-yamakoshi/go-grpc-sample" -w .
	buf format -w ./schema/proto/

.PHONY:lint
lint:
	golangci-lint run ./...

.PHONY: proto_gen
proto_gen:
	protoc --proto_path=$(PROTO_DIR) \
	--go_out=$(OUT_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(OUT_DIR) \
	--go-grpc_opt=paths=source_relative \
	$(addprefix $(PROTO_DIR)/, $(notdir $(PROTO_FILES)))

.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: actions-dry-run
actions-dry-run:
	act -n
