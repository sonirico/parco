# Performance Analysis

Complete benchmark comparison and honest discussion about what these numbers actually mean.

## Benchmark Results

**Test Environment**
- CPU: Intel(R) Core(TM) i7-8750H @ 2.20GHz
- OS: Linux
- Go Version: 1.23

```
BenchmarkParcoAlloc_Compile/small_size-12     513920   2373 ns/op    91 B payload   184 B/op   3 allocs/op
BenchmarkParcoAlloc_Compile/medium_size-12     82522  15727 ns/op   742 B payload   184 B/op   3 allocs/op
BenchmarkParcoAlloc_Compile/large_size-12       7939 154303 ns/op  8123 B payload   184 B/op   3 allocs/op

BenchmarkJson_Compile/small_size-12           353646   2972 ns/op   269 B payload   970 B/op  23 allocs/op
BenchmarkJson_Compile/medium_size-12           43119  28305 ns/op  1681 B payload  7688 B/op 203 allocs/op
BenchmarkJson_Compile/large_size-12             3609 319851 ns/op 16637 B payload 76527 B/op 2003 allocs/op

BenchmarkMsgpack_Compile/small_size-12        336514   3020 ns/op   155 B payload   762 B/op  25 allocs/op
BenchmarkMsgpack_Compile/medium_size-12        54040  20051 ns/op   991 B payload  4070 B/op 207 allocs/op
BenchmarkMsgpack_Compile/large_size-12          5658 206755 ns/op 10171 B payload 37451 B/op 2007 allocs/op
```

## What These Numbers Mean

**Speed**

| Payload Size | Parco | JSON | MessagePack | vs JSON | vs MessagePack |
|--------------|-------|------|-------------|---------|----------------|
| Small (91B)  | 2.4µs | 3.0µs | 3.0µs     | 1.25x faster | 1.27x faster |
| Medium (742B)| 15.7µs | 28.3µs | 20.1µs  | 1.80x faster | 1.27x faster |
| Large (8KB)  | 154µs | 320µs | 207µs    | 2.07x faster | 1.34x faster |

Parco's advantage grows with payload size. This isn't magic, it's what happens when you skip all the metadata overhead.

**Memory**

| Format | Small | Medium | Large |
|--------|-------|--------|-------|
| Parco | 184 B (3 allocs) | 184 B (3 allocs) | 184 B (3 allocs) |
| JSON | 970 B (23 allocs) | 7,688 B (203 allocs) | 76,527 B (2003 allocs) |
| MessagePack | 762 B (25 allocs) | 4,070 B (207 allocs) | 37,451 B (2007 allocs) |

Parco uses constant memory. JSON and MessagePack allocate proportionally to payload size. This matters for garbage collection--more allocations mean more GC pressure.

**Payload Size**

| Payload Size | Parco | JSON | MessagePack | vs JSON | vs MessagePack |
|--------------|-------|------|-------------|---------|----------------|
| Small | 91 B | 269 B | 155 B | 66% smaller | 41% smaller |
| Medium | 742 B | 1,681 B | 991 B | 56% smaller | 25% smaller |
| Large | 8,123 B | 16,637 B | 10,171 B | 51% smaller | 20% smaller |

## Why Parco is Faster

**No Reflection**

JSON uses runtime reflection to inspect structs. This is slow. Parco uses a builder API that generates direct function calls at compile time.

```go
// JSON - uses reflection (slow)
json.Marshal(obj)

// Parco - direct function calls (fast)
compiler.Compile(obj, writer)
```

**Minimal Allocations**

Parco uses a pool for temporary buffers and reuses them. Three allocations per serialization, regardless of payload size. JSON creates new allocations for every field and nested structure.

**No Field Names**

```
JSON:  {"name":"Alice","age":30}  → 28 bytes
Parco: [5]Alice[30]               → 7 bytes
```

JSON transmits field names with every message. Parco assumes both sides know the schema.

**Efficient Integer Encoding**

JSON encodes numbers as ASCII strings. `42` becomes `"42"` (2+ bytes). Parco writes `42` as a single byte.

**No Type Tags**

JSON includes type information for every value (string, number, array, object). Parco defines types at compile time and doesn't transmit them.

## Common Questions

**Q: Is this a fair comparison?**

Yes, with caveats. It's like comparing a sports car to a pickup truck--both are vehicles, but they're designed for different purposes.

Parco is optimized for performance when you control both ends and know the schema. JSON is optimized for flexibility when you don't. Both are valid design choices.

**Q: Why not just use Protocol Buffers?**

| Feature | Parco | Protobuf |
|---------|-------|----------|
| Schema definition | Go code | .proto files + codegen |
| Type safety | Compile-time (generics) | Compile-time (generated) |
| Dependencies | Zero | protoc compiler + plugins |
| Learning curve | Go knowledge | Proto syntax + tooling |
| Flexibility | Full Go control | Proto language limits |

Use Parco if you prefer code over DSL and don't want a codegen pipeline. Use Protobuf if you need cross-language support.

**Q: What about schema evolution?**

Parco supports schema evolution through:

**Optional fields**: Add new fields as `Option[T]`
```go
// v1
Builder[UserV1]().
    Varchar(...).  // name
    UInt8(...)     // age

// v2
Builder[UserV2]().
    Varchar(...).  // name
    UInt8(...).    // age
    Option(        // email (new)
        OptionField[UserV2, string](Varchar(), ...),
    )
```

**Skip padding**: Reserve bytes for future use
```go
Builder[User]().
    Varchar(...).
    UInt8(...).
    Skip(10)  // reserve 10 bytes for future fields
```

**Multi-model registry**: Version models with IDs
```go
const (
    UserV1 = 1
    UserV2 = 2
)

MultiBuilder(UInt8Header()).
    MustRegister(UserV1, userV1Builder).
    MustRegister(UserV2, userV2Builder)
```

## Optimization Tips

**Reuse Parsers and Compilers**

```go
// Bad: create on every call
func Serialize(user User) []byte {
    _, compiler := Builder[User](...).Parco()
    // ...
}

// Good: create once, reuse
var userCompiler = Builder[User](...).Compiler()

func Serialize(user User) []byte {
    userCompiler.Compile(user, writer)
}
```

**Use Appropriate Header Sizes**

```go
// For up to 255 items
Slice(UInt8Header(), ...)    // 1 byte

// For up to 65,535 items
Slice(UInt16HeaderLE(), ...) // 2 bytes

// Don't use UInt32 if UInt8 suffices
```

**Pool Writers**

```go
var bufferPool = sync.Pool{
    New: func() any { return new(bytes.Buffer) },
}

func Serialize(user User) []byte {
    buf := bufferPool.Get().(*bytes.Buffer)
    defer func() {
        buf.Reset()
        bufferPool.Put(buf)
    }()
    compiler.Compile(user, buf)
    return buf.Bytes()
}
```

**Profile Before Optimizing**

```bash
# CPU profile
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof

# Memory profile
go test -bench=. -memprofile=mem.prof
go tool pprof mem.prof
```

## Further Reading

- [Go performance tips](https://github.com/dgryski/go-perfbook)
- [Protocol Buffers encoding](https://protobuf.dev/programming-guides/encoding/)
- [Benchmark methodology](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
