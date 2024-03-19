VERSION := $(shell git describe --tags --always --dirty="-dev") # generate version like `v1.0.0-2-g2414721`

# clean old builds and rebuild all. run `make all` or just `make`
.PHONY: all
all: clean build

# version. make v or make version. version and v are aliases
.PHONY: version v
version v:
	@echo "Version: ${VERSION}"

# build all under build/app-${version} and set variable `version` in main pacakge.
.PHONY: build
build:
	go build ./...

# clean build
.PHONY: clean
clean:
	rm -rf build/

# run all tests
.PHONY: test
test:
	go test -v ./...

# test coverage
.PHONY: coverage
coverage:
	go test -coverprofile=/tmp/go-sim-lb.cover.tmp ./... vvv
	go tool cover -func /tmp/go-sim-lb.cover.tmp
	unlink /tmp/go-sim-lb.cover.tmp

# format all files
.PHONY: fmt
fmt:
	go fmt ./...
	go mod tidy


