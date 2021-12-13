VERSION := $(shell cat VERSION)
XC_OS 	:= linux darwin
XC_OS 	:= linux
XC_ARCH := 386 amd64 arm
XC_ARCH := amd64
LD_FLAGS := -X main.version=$(VERSION) -s -w
SOURCE_FILES ?=./internal/... ./pkg/...
TEST_PATTERN ?=.
BENCH_OPTIONS ?= -bench=. -benchtime=1000x -benchmem
TEST_OPTIONS := -v -failfast -race
CLEAN_OPTIONS ?=-modcache -testcache
CLEAN_OPTIONS :=-testcache
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
	@echo "make clean - clean test cache, build files"
	@echo "make build - build $(PROJECT) for following OS-ARCH constilations: $(XC_OS) / $(XC_ARCH) "
	@echo "make build-dev - build $(PROJECT) for OS-ARCH set by GOOS and GOARCH env variables"
	@echo "make build-docker - build $(PROJECT) for linux-amd64 docker image"
	@echo "make fmt - use gofmt & goimports"
	@echo "make lint - run golangci-lint"
	@echo "make test - run go test including race detection"
	@echo "make coverage - same as test and uses go-junit-report to create report.xml"
	@echo "make dist - build and create packages with hashsums"
	@echo "make build-docker - creates a docker image"
	@echo "make run-docker - creates a docker image"
	@echo "make docker-release/docker-release-latest - creates the docker image and pushes it to the registry (latest pushes also latest tag)"
	@echo "make setup - adds git pre-commit hooks"


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
	@CGO_ENABLED=1 go test ${TEST_OPTIONS} ${BENCH_OPTIONS} ${SOURCE_FILES} -run ${TEST_PATTERN} -timeout=${TEST_TIMEOUT}



