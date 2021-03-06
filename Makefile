export GO111MODULE=on
export GOPROXY=

export PATH := $(GOPATH)/bin:$(PATH)

BINARY_VERSION?=0.0.1
BINARY_OUTPUT?=motionerd
EXTRA_FLAGS?=-mod=vendor

.PHONY: all install build test clean deps upgrade print

all: deps build

install:
	go install -v $(EXTRA_FLAGS) -ldflags "-X main.Version=$(BINARY_VERSION)"

build:
	go build -v $(EXTRA_FLAGS) -ldflags "-X main.Version=$(BINARY_VERSION)" -o $(BINARY_OUTPUT)

test:
	go test -v $(EXTRA_FLAGS) -race -coverprofile=coverage.txt -covermode=atomic ./...

clean:
	go clean
	rm -f $(BINARY_NAME)

deps:
	go build -v $(EXTRA_FLAGS) ./...

upgrade:
	go get -u ./...
	go mod vendor

tidy:
	go mod tidy

print:
	@echo "PWD: $(shell pwd)"
	@echo "PATH: $(PATH)"
	@echo "GOPATH: $(GOPATH)"
	@echo "GOBIN: $(GOBIN)"
	@echo "GOPROXY: $(GOPROXY)"
	@echo "GO111MODULE: $(GO111MODULE)"
