.PHONY: generate clean build run test generate-buf generate-protoc

# Check if buf is available
BUF_AVAILABLE := $(shell command -v buf 2> /dev/null)

# Generate protobuf code (automatically choose buf or protoc)
generate:
ifdef BUF_AVAILABLE
	@echo "Using buf for protobuf generation..."
	@$(MAKE) generate-buf
else
	@echo "Using protoc for protobuf generation..."
	@$(MAKE) generate-protoc
endif

# Generate using buf (preferred)
generate-buf:
	buf generate

# Generate using protoc (fallback)
generate-protoc:
	mkdir -p gen/game/v1
	# First compile common.proto
	protoc --proto_path=. --go_out=. --go_opt=paths=source_relative protos/game/v1/common.proto
	# Then compile all service protos
	for proto in protos/game/v1/*_service.proto; do \
		protoc --proto_path=. --go_out=. --go_opt=paths=source_relative \
			--connect-go_out=. --connect-go_opt=paths=source_relative \
			"$$proto"; \
	done
	# Move generated files to correct location
	mkdir -p gen/game/v1
	find protos/game/v1 -name "*.pb.go" -exec mv {} gen/game/v1/ \;
	find protos/game/v1 -name "*connect" -type d -exec mv {} gen/game/v1/ \;

# Lint protobuf files
lint:
ifdef BUF_AVAILABLE
	buf lint
else
	@echo "buf not available, skipping lint"
endif

# Format protobuf files
format:
ifdef BUF_AVAILABLE
	buf format -w
else
	@echo "buf not available, skipping format"
endif

# Install dependencies
deps:
	go mod download
	go mod tidy

# Clean generated files
clean:
	rm -rf gen/

# Build the server
build:
	go build -o bin/server ./cmd/server

# Run the server
run:
	go run ./cmd/server

# Run tests
test:
	go test ./...

# Install protobuf tools
tools:
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Install buf (preferred protobuf tool)
install-buf:
	go install github.com/bufbuild/buf/cmd/buf@latest

# Setup all tools
setup: tools install-buf
