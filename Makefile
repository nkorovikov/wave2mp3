BIN_DIR := bin
APP_NAME := wave2mp3
SRC_DIR := cmd

LINT_BIN_DIR := $(shell go env GOPATH)/bin
GOLANGCI_LINT := $(LINT_BIN_DIR)/golangci-lint
VERSION := latest

# Assembly for the current OS
.PHONY: build
build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) ./$(SRC_DIR)
	@echo "Built: $(BIN_DIR)/$(APP_NAME)"

# Cross-compilation
.PHONY: release
release:
	@mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)-linux ./$(SRC_DIR)
	GOOS=darwin GOARCH=arm64 go build -o $(BIN_DIR)/$(APP_NAME)-macos ./$(SRC_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)-win.exe ./$(SRC_DIR)
	@echo "Cross-compilation complete!"

# Cleaning binaries
.PHONY: clean
clean:
	@rm -rf $(BIN_DIR)
	@echo "Cleaned $(BIN_DIR)"

# Installing dependencies
.PHONY: deps
deps:
	go mod download
	go mod tidy

# Linter launch
.PHONY: lint
lint: $(GOLANGCI_LINT)
	@echo "🛠 Launch golangci-lint..."
	@$(GOLANGCI_LINT) run --config .golangci.yml

$(GOLANGCI_LINT):
	@echo "📦 Установка golangci-lint $(VERSION)..."
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LINT_BIN_DIR) $(VERSION)
	@echo "✅ Done! The linter is installed in $(GOLANGCI_LINT)"

# Linter deleting
.PHONY: cleanlint
cleanlint:
	@rm -f $(GOLANGCI_LINT)
	@echo "🧹 Linter deleted"

# Help
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build    - Build for current platform"
	@echo "  make release  - Cross compilation for all platforms"
	@echo "  make clean    - Clear binaries"
	@echo "  make deps     - Install dependencies"
	@echo "  make lint     - Launch linter"
	@echo "  make cleanlint   - Delete linter"