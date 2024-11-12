# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
AVS_BINARY=avs
OPERATOR_BINARY=operator
EXOKEY_BINARY=exokey

# Go version
GO_VERSION=1.22

# Build flags
LDFLAGS=-ldflags "-s -w"

# Default target
all: clean build

# Build all binaries
build: avs operator exokey

# AVS build
avs:
	$(GOBUILD) $(LDFLAGS) -o $(AVS_BINARY) avs/cmd/main.go

# Operator build
operator:
	$(GOBUILD) $(LDFLAGS) -o $(OPERATOR_BINARY) operator/cmd/main.go

# ExoKey build
exokey:
	$(GOBUILD) $(LDFLAGS) -o $(EXOKEY_BINARY) exokey/cmd/main.go

# Clean build artifacts
clean:
	rm -f $(AVS_BINARY) $(OPERATOR_BINARY) $(EXOKEY_BINARY)

# Run tests
test:
	$(GOTEST) ./...

# Install dependencies
deps:
	$(GOGET) -v ./...

# Lint code
lint:
	golangci-lint run

# Cross-platform builds
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(AVS_BINARY)-linux avs/cmd/main.go
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(OPERATOR_BINARY)-linux operator/cmd/main.go

build-darwin:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(AVS_BINARY)-darwin avs/cmd/main.go
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(OPERATOR_BINARY)-darwin operator/cmd/main.go

# Import key command
import-key:
	./$(EXOKEY_BINARY) import --key-type ecdsa $(PRI_KEY)

# Phony targets
.PHONY: all build avs operator exokey clean test deps lint build-linux build-darwin import-key