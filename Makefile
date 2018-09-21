SHELL=bash

MKDIR_P = mkdir -p
VERSION=$(shell cat VERSION)
GOVERSION=$(shell go version)
BUILDHASH=$(shell git rev-parse --verify --short HEAD)

ROOTDIR=$(shell pwd)
BINDIR=$(ROOTDIR)/bin
DISTDIR=$(ROOTDIR)/dist
TMPDIR=$(ROOTDIR)/tmp

BINARY_NAME=grpc-bitcoin
BINARY=$(BINDIR)/$(BINARY_NAME)

$(BINDIR):
	$(MKDIR_P) $@

$(TMPDIR):
	$(MKDIR_P) $@

$(BINARY): $(BINDIR)
	@go build -o $@ ./cli

.PHONY: setup
## install development packages
setup: $(TMPDIR)
	@if [ -z `which golint 2> /dev/null` ]; then \
		go get github.com/golang/lint/golint; \
		fi
	@if [ -z `which make2help 2> /dev/null` ]; then \
		go get github.com/Songmu/make2help/cmd/make2help; \
		fi
	@if [ -z `which dep 2> /dev/null` ]; then \
		go get github.com/golang/dep/cmd/dep; \
		fi
	@if [ -z `which gnatsd 2> /dev/null` ]; then \
		go get github.com/nats-io/gnatsd; \
		fi
	@if [ -z `which reflex 2> /dev/null` ]; then \
		go get github.com/cespare/reflex; \
		fi
	@if [ -z `which godotenv 2> /dev/null` ]; then \
		go get github.com/joho/godotenv; \
		fi

.PHONY: dep
## install dependencies packages
dep: setup
	dep ensure

.PHONY: latest-dep
## Upgrade dependent packages
latest-dep: setup
	@dep ensure -update

.PHONY: build
## build binary
build: clean dep $(BINARY)

.PHONY: gnatsd
## running nuts server
gnatsd:
	@gnatsd -D -V

.PHONY: clean
## clean up tmp dir and binary
clean:
	@rm -rf $(TMPDIR)/* $(BINARY)
