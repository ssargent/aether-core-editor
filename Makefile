.PHONY: generate clean build run test

# Generate protobuf code
generate:
	protoc --go_out=gen --go_opt=paths=source_relative \
		--connect-go_out=gen --connect-go_opt=paths=source_relative \
		game_v1.proto

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
