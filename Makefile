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
