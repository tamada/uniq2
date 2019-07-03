GO=go
NAME := uniq2
VERSION := 1.0.1
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)'
	-X 'main.revision=$(REVISION)'

all: test build

setup: update_version
	git submodule update --init

update_version:
	@for i in README.md docs/content/_index.md; do\
	    sed -e 's!Version-[0-9.]*-yellowgreen!Version-${VERSION}-yellowgreen!g' -e 's!tag/v[0-9.]*!tag/v${VERSION}!g' $$i > a ; mv a $$i; \
	done
	@sed 's/const VERSION = .*/const VERSION = "${VERSION}"/g' cmd/uniq2/main.go > a
	@mv a cmd/uniq2/main.go
	@echo "Replace version to \"${VERSION}\""

test: setup
	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./... | grep -v vendor)

build: setup
	cd cmd/uniq2; $(GO) build -o $(NAME) -v

clean:
	$(GO) clean
	rm -rf cmd/uniq2/$(NAME)
