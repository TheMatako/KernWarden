# Binary output name
APP_NAME=warden
# Entry point path
CMD_PATH=./cmd/warden
# Binary destination directory
BIN_DIR=bin

# Phony targets declaration
.PHONY: all build run clean fmt vet test

# Default target
all: fmt vet build

# Compile the code
build:
	@echo "==> Compiling $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	go build -ldflags="-s -w" -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)/main.go
	@echo "==> Binary generated in $(BIN_DIR)/$(APP_NAME)"

# Run on the fly (development mode)
run:
	@echo "==> Running agent in development mode..."
	go run $(CMD_PATH)/main.go

# Strict code formatting (Go Standard)
fmt:
	@echo "==> Formatting code..."
	go fmt ./...

# Static analysis for vulnerabilities and dead code
vet:
	@echo "==> Running static analysis (Go Vet)..."
	go vet ./...

# Execute unit tests
test:
	@echo "==> Running tests..."
	go test -v ./...

# Clean workspace
clean:
	@echo "==> Cleaning binaries..."
	rm -rf $(BIN_DIR)