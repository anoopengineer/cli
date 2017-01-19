default: deps fmt vet test build

test:
	@echo "Running go test"
	go test -v -timeout 60s -race ./...
	@echo "All go tests passed"

vet:
	@echo "Running go vet"
	go vet ./...
	@echo "go vet passed"

build:
	@echo "Running go build"
	go build ./...
	@echo "Running go build completed"

fmt:
	@echo "Running go fmt"
	@if [ -n "$$(go fmt ./...)" ]; then echo 'Please run go fmt on your code.' && exit 1; fi
	@echo "Running go fmt completed"

deps:
	@echo "Running go get"
	go get ./...