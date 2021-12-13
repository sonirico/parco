VERSION := $(shell cat VERSION)
XC_OS 	:= linux darwin
XC_OS 	:= linux
XC_ARCH := 386 amd64 arm
XC_ARCH := amd64
LD_FLAGS := -X main.version=$(VERSION) -s -w
SOURCE_FILES ?=./internal/... ./pkg/...
TEST_OPTIONS := -v -failfast -race
TEST_PATTERN ?=.
BENCH_OPTIONS ?= -v -bench=. -benchmem
BENCH_TEST_OPTIONS := -v -failfast -race
CLEAN_OPTIONS ?=-modcache -testcache
TEST_TIMEOUT ?=1m
LINT_VERSION := 1.40.1

export CGO_ENABLED=0
export XC_OS
export XC_ARCH
export VERSION
export PROJECT
export GO111MODULE=on
export LD_FLAGS
export SOURCE_FILES
export TEST_PATTERN
export TEST_OPTIONS
export TEST_TIMEOUT
export LINT_VERSION
export MallocNanoZone ?= 0

.PHONY: all
all: help

.PHONY: help
help:
	@echo "make fmt - use gofmt & goimports"
	@echo "make lint - run golangci-lint"
	@echo "make test - run go test including race detection"
	@echo "make bench - run go test including benchmarking"


.PHONY: format
format:
	$(info: Make: Format)
	@echo gofmt -w -w pkg/* internal/*
	@gofmt -w -w pkg/* internal/*
	@echo goimport -w pkg/* internal/*
	@goimports -w pkg/* internal/*

.PHONY: lint
lint:
	$(info: Make: Lint)
	@golangci-lint run --tests=false


.PHONY: test
test:
	@echo CGO_ENABLED=1 go test ${TEST_OPTIONS} ${SOURCE_FILES} -run ${TEST_PATTERN} -timeout=${TEST_TIMEOUT}
	@CGO_ENABLED=1 go test ${TEST_OPTIONS} ${SOURCE_FILES} -run ${TEST_PATTERN} -timeout=${TEST_TIMEOUT}

.PHONY: bench
bench:
	@echo CGO_ENABLED=1 go test ${TEST_OPTIONS} ${BENCH_OPTIONS} ${SOURCE_FILES} -run ${TEST_PATTERN} -timeout=${TEST_TIMEOUT}
	@CGO_ENABLED=1 go test ${BENCH_OPTIONS} ${SOURCE_FILES} -run ${TEST_PATTERN} -timeout=${TEST_TIMEOUT}



