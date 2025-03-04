# Project name
APP_NAME=SimpleShell
BIN_DIR=bin
CMD_DIR=cmd/SimpleShell

# Build the application
build:
	@go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)/main.go
	@echo "Build complete!"

# Run the application
run:
	@go run $(CMD_DIR)/main.go

# Run tests
test:
	@go test ./...

# Clean up binaries and cache
clean:
	@rm -rf $(BIN_DIR)/*
	@go clean
	@echo "Cleaned up!"

# Format the code
fmt:
	@go fmt ./...

# Install dependencies
deps:
	@go mod tidy

# Help menu
help:
	@echo "Usage:"
	@echo "  make build    - Compile the project"
	@echo "  make run      - Run the application"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Remove built files"
	@echo "  make fmt      - Format the code"
	@echo "  make deps     - Install dependencies"
