GO=go
NAME := uniq2
VERSION := 0.1.0
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)'
	-X 'main.revision=$(REVISION)'

all: test build

deps:
	$(GO) get golang.org/x/lint/golint
	$(GO) get golang.org/x/tools/cmd/goimports
	$(GO) get github.com/golang/dep/cmd/dep

	$(GO) get golang.org/x/tools/cmd/cover
	$(GO) get github.com/mattn/goveralls

	dep ensure -vendor-only

setup: deps
	git submodule update --init

test: setup 
	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./... | grep -v vendor)

build: setup
	$(GO) build -o $(NAME) -v

clean:
	$(GO) clean
	rm -rf $(NAME)
