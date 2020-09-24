# cross parameters
SHELL:=/bin/bash -O extglob
BINARY=agileEngine
VERSION=0.1.0

LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# Build step, generates the binary.
build: clean
	@go build ${LDFLAGS} -o ${BINARY} cmd/web/main/*.go

clean: ## Clean the project, set it up for a new build
	@rm -rf internal/templates/assets/bindata.go
	@rm -rf ${BINARY}

# Web is a mask to run the web interface, in our case the main function will start the http server.
web:
	@clear
	@go run cmd/web/main/!(*_test).go

# Run go formatter
fmt:
	@gofmt -w .

# Run the test for all the directories.
test:
	@clear
	@go test -v ./...
