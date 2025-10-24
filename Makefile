.PHONY: build install clean test run

# Binary name
BINARY=sniprun
VERSION?=0.1.0

# Build variables
BUILD_DIR=bin
MAIN_PATH=main.go

# Go build flags
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

# Build the binary
build:
	@echo "Building $(BINARY)..."
	@mkdir -p $(BUILD_DIR)
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY) $(MAIN_PATH)
	@echo "Built: $(BUILD_DIR)/$(BINARY)"

# Build for all platforms
build-all:
	@echo "Building for all platforms..."
	@mkdir -p $(BUILD_DIR)
	
	@echo "Building for Linux (amd64)..."
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-linux-amd64 $(MAIN_PATH)
	
	@echo "Building for macOS (amd64)..."
	@GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-darwin-amd64 $(MAIN_PATH)
	
	@echo "Building for macOS (arm64)..."
	@GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-darwin-arm64 $(MAIN_PATH)
	
	@echo "Building for Windows (amd64)..."
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-windows-amd64.exe $(MAIN_PATH)
	
	@echo "All builds complete!"

# Install to system
install: build
	@echo "Installing $(BINARY)..."
	@sudo cp $(BUILD_DIR)/$(BINARY) /usr/local/bin/
	@echo "Installed to /usr/local/bin/$(BINARY)"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "Clean complete"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run the application
run:
	@go run $(MAIN_PATH) $(ARGS)

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

# Show help
help:
	@echo "Available targets:"
	@echo "  build      - Build the binary"
	@echo "  build-all  - Build for all platforms"
	@echo "  install    - Install to /usr/local/bin"
	@echo "  clean      - Remove build artifacts"
	@echo "  test       - Run tests"
	@echo "  run        - Run the application (use ARGS=... for arguments)"
	@echo "  fmt        - Format code"
	@echo "  lint       - Run linter"
	@echo "  deps       - Download dependencies"
	@echo "  help       - Show this help"