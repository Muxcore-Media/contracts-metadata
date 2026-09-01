.PHONY: build test lint clean fmt tidy proto ci

GO ?= go
PROTOC_GEN_GO_VERSION ?= v1.36.6
PROTOC_GEN_GO_GRPC_VERSION ?= v1.5.1

build:
	$(GO) build ./...

test:
	$(GO) test -race -count=1 -timeout 60s ./...

lint:
	golangci-lint run --timeout 120s ./...

# Published stubs under muxcore/metadata/v1/ are committed; do not delete them.
clean:
	$(GO) clean -cache -testcache

fmt:
	$(GO) fmt ./...

tidy:
	$(GO) mod tidy

proto:
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)
	PATH="$$(go env GOPATH)/bin:$$PATH" protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		-I proto proto/muxcore/metadata/v1/metadata.proto

ci: lint test build
