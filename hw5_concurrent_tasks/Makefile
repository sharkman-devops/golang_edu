SHELL:=/bin/bash
FILES_EXCLUDE_VENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

help:
	@echo "make all"
	@echo "make gofmt_check"
	@echo "make goimports_check"
	@echo "make govet_check"
	@echo "make golint_check"
	@echo "make tests"

.PHONY: all
all: gofmt_check goimports_check govet_check golint_check tests

# checks
.PHONY: gofmt_check
gofmt_check:
	@echo "Running gofmt..."
	@if [ `gofmt -l ${FILES_EXCLUDE_VENDOR} | wc -l` -ne 0 ]; then \
		echo "Failed on files: " ; \
		gofmt -l ${FILES_EXCLUDE_VENDOR} ; \
		exit 1 ;\
	fi
	@echo "Success!"

.PHONY: goimports_check
goimports_check:
	@echo "Running goimports..."
	@if [ `goimports -l ${FILES_EXCLUDE_VENDOR} | wc -l` -ne 0 ]; then \
		echo "Failed on files: " ; \
		goimports -l ${FILES_EXCLUDE_VENDOR} ; \
		exit 1 ;\
	fi
	@echo "Success!"

.PHONY: govet_check
govet_check:
	@echo "Running go vet..."
	go vet
	@echo "Success!"

.PHONY: golint_check
golint_check:
	@echo "Running golint..."
	golint
	@echo "Success!"

.PHONY: tests
tests:
	go test -v
	go test -race -v
