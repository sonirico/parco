# Contributing to Parco

First off, thank you for considering contributing to Parco! 🎉

## Code of Conduct

Be respectful, inclusive, and constructive. We're all here to learn and build cool things together.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check existing issues. When you create a bug report, include:

- **Clear title and description**
- **Steps to reproduce** the behavior
- **Expected vs actual behavior**
- **Go version** (`go version`)
- **OS and architecture**
- **Code sample** (if applicable)

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion:

- **Use a clear and descriptive title**
- **Provide a detailed description** of the enhancement
- **Explain why this would be useful** to most users
- **List any alternative solutions** you've considered

### Pull Requests

1. Fork the repo and create your branch from `main`
2. Write tests for your changes
3. Ensure the test suite passes
4. Make sure your code follows the existing style
5. Write clear, descriptive commit messages
6. Update documentation as needed

## Development Setup

### Prerequisites

- Go 1.19 or higher
- Git

### Getting Started

```bash
# Clone your fork
git clone https://github.com/YOUR-USERNAME/parco.git
cd parco

# Install dev tools
just setup
# or
make setup

# Run tests
just test

# Run linter
just lint
```

## Development Workflow

### Using Just (Recommended)

```bash
# See all commands
just --list

# Run tests with coverage
just test-coverage

# Run benchmarks
just bench

# Format code
just format

# Run all checks
just ci
```

### Using Make

```bash
# Run tests
make test

# Run benchmarks
make bench

# Format code
make format

# Lint
make lint
```

## Testing Guidelines

### Writing Tests

- Test files should be named `*_test.go`
- Use table-driven tests where appropriate
- Test both success and error cases
- Use descriptive test names

Example:
```go
func TestSliceType_Parse(t *testing.T) {
    tests := []struct {
        name    string
        input   []byte
        want    []int
        wantErr bool
    }{
        {"empty slice", []byte{0}, []int{}, false},
        {"single element", []byte{1, 42, 0, 0, 0}, []int{42}, false},
        {"invalid length", []byte{1}, nil, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

### Running Tests

```bash
# All tests
just test

# Specific test
go test -v -run TestSliceType_Parse

# With coverage
just test-coverage

# Race detector
go test -race ./...
```

### Benchmarks

When adding new features, include benchmarks:

```go
func BenchmarkSliceType_Parse(b *testing.B) {
    sliceType := Slice[int](UInt8Header(), IntLE())
    data := /* ... prepare test data ... */

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := sliceType.Parse(bytes.NewReader(data))
        if err != nil {
            b.Fatal(err)
        }
    }
}
```

Run benchmarks:
```bash
just bench
```

## Style Guide

### Code Style

Follow standard Go conventions:

- Use `gofmt` and `goimports`
- Keep functions small and focused
- Document exported functions and types
- Use meaningful variable names

### Documentation

- All exported types and functions must have doc comments
- Doc comments should start with the name of the thing being described
- Use complete sentences

Example:
```go
// SliceType represents a variable-length slice with a header indicating length.
// The header type determines the maximum number of elements (e.g., UInt8Header allows up to 255).
type SliceType[T any] struct {
    header IntType
    inner  Type[T]
}
```

### Commit Messages

Follow conventional commits:

```
<type>(<scope>): <subject>

<body>

<footer>
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Formatting, missing semicolons, etc.
- `refactor`: Code change that neither fixes a bug nor adds a feature
- `perf`: Performance improvement
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Examples:
```
feat(slice): add MaxReasonableSliceLength validation

Add a constant MaxReasonableSliceLength (10M elements) to prevent
excessive memory allocation from malicious or corrupted data.

Closes #123
```

```
fix(pool): correct Put8 using wrong pool

Put8() was incorrectly using pool4 instead of pool8, causing
potential memory issues.
```

## Areas to Contribute

Looking for ideas? Here are some areas where contributions are welcome:

### High Priority

- [ ] More comprehensive tests for edge cases
- [ ] Performance optimizations
- [ ] Documentation improvements
- [ ] Example applications

### Medium Priority

- [ ] Support for more Go types (complex64/128, rune, etc.)
- [ ] Benchmarks against more formats (CBOR, Avro, etc.)
- [ ] Schema evolution helpers
- [ ] Validation utilities

### Low Priority

- [ ] Code generation from struct tags
- [ ] Alternative encodings (varint, zigzag, etc.)
- [ ] Streaming API improvements
- [ ] Custom allocator support

## Performance Considerations

When contributing, keep these in mind:

1. **Avoid allocations** - Use pools, reuse buffers
2. **Profile before optimizing** - Use `just bench-profile`
3. **Benchmark your changes** - Compare before/after
4. **Maintain zero dependencies** - Core API should have no deps

## Questions?

Feel free to open an issue with the `question` label or reach out to the maintainers.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
