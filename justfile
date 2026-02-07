# Parco - Just commands
# https://github.com/casey/just

# List all available commands
default:
    @just --list

# Run tests with race detector
test:
    CGO_ENABLED=1 go test -v -race -failfast -timeout=1m ./...

# Run tests with coverage
test-coverage:
    CGO_ENABLED=1 go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
    go tool cover -html=coverage.txt -o coverage.html
    @echo "Coverage report generated: coverage.html"

# Run benchmarks
bench:
    CGO_ENABLED=1 go test -bench=. -benchmem -run=^$ ./...

# Run benchmarks and compare with previous results
bench-compare:
    CGO_ENABLED=1 go test -bench=. -benchmem -run=^$ ./... | tee new.txt
    @if [ -f old.txt ]; then \
        benchstat old.txt new.txt; \
    fi
    @mv new.txt old.txt

# Run benchmarks with memory and CPU profiling
bench-profile:
    CGO_ENABLED=1 go test -bench=. -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof -run=^$ ./...
    @echo "Profiles generated: cpu.prof, mem.prof"
    @echo "View with: go tool pprof cpu.prof"

# Format code
format:
    @echo "Running gofmt..."
    gofmt -w .
    @echo "Running goimports..."
    goimports -w .
    @echo "Running golines..."
    golines -w --max-len=120 .
    @echo "✓ Code formatted"

# Run linter
lint:
    golangci-lint run --timeout=5m

# Run linter with auto-fix
lint-fix:
    golangci-lint run --timeout=5m --fix

# Build examples
build:
    go build -v ./...
    go build -gcflags='-m -m' examples/compiler/*.go

# Clean build artifacts and caches
clean:
    go clean -cache -testcache -modcache
    rm -f coverage.txt coverage.html cpu.prof mem.prof old.txt new.txt

# Install development tools
setup:
    go install golang.org/x/tools/cmd/goimports@latest
    go install github.com/segmentio/golines@latest
    go install golang.org/x/perf/cmd/benchstat@latest
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    @echo "✓ Development tools installed"

# Run all checks (test, lint, format check)
ci: test lint
    @echo "✅ All checks passed!"

# Update dependencies
update-deps:
    go get -u ./...
    go mod tidy

# Run a specific example
example NAME:
    go run ./examples/{{NAME}}

# Run all examples
examples:
    @for dir in examples/*/; do \
        echo "Running $${dir}..."; \
        go run ./$${dir} || exit 1; \
    done

# Generate test mocks (if using mockgen)
mocks:
    @echo "No mocks to generate yet"

# Show project statistics
stats:
    @echo "Lines of code:"
    @find . -name '*.go' -not -path './vendor/*' | xargs wc -l | tail -1
    @echo "\nTest files:"
    @find . -name '*_test.go' | wc -l
    @echo "\nPackages:"
    @go list ./... | wc -l
