[![CI](https://github.com/sonirico/parco/actions/workflows/ci.yml/badge.svg)](https://github.com/sonirico/parco/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sonirico/parco)](https://goreportcard.com/report/github.com/sonirico/parco)
[![GoDoc](https://godoc.org/github.com/sonirico/parco?status.svg)](https://godoc.org/github.com/sonirico/parco)
[![License](https://img.shields.io/github/license/sonirico/parco)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/sonirico/parco)](go.mod)

# Parco

**Parco** is a high-performance binary serializer and deserializer for Go: no reflection, highly extensible, focused on speed and usability through generics, with zero external dependencies for the core API.

**Performance**: 2-3x faster than JSON, payloads 40% smaller than MessagePack, 10x fewer allocations.

**A note on honesty**: These numbers might make you suspicious, and you should be. Parco achieves this by trading flexibility for speed: it uses a known schema approach (like Protocol Buffers) rather than self-describing formats (like JSON). Both sender and receiver must agree on the data structure beforehand. This isn't a sneaky benchmark trick--it's a fundamental design decision for specific use cases. If you need to inspect unknown JSON from the wild, Parco won't help you. If you control both ends of the wire and want it fast, keep reading. See the [benchmarks section](#benchmarks) for an honest comparison and guidance on when to use what.

Many packages rely on struct tags and reflection, which is convenient until you need custom types or care about performance. Parco uses a builder API with getters and setters: you describe how each field is read and written, and the library generates efficient, allocation-conscious code. Think of it as Go code that writes Go code, except you're the one writing it.


## Table of contents

- [Features](#features)
- [Installation](#installation)
- [Concepts](#concepts)
- [Quick start](#quick-start)
- [Usage](#usage)
  - [Parser & compiler (single model)](#parser--compiler-single-model)
  - [Single types](#single-types)
  - [Multi-model (registry)](#multi-model-parsers--compilers)
- [Supported types](#supported-types)
- [Error handling](#error-handling)
- [Examples](#examples)
- [Benchmarks](#benchmarks)
- [Development](#development)
- [Contributing](#contributing)
- [Roadmap](#roadmap)
- [License](#license)
- [Additional Resources](#additional-resources)


## Features

- **No reflection** -- Types and fields are described via a fluent builder and generics.
- **Composable** -- Use primitives, slices, maps, structs, and optional types; nest and reuse builders.
- **Multi-model on the same stream** -- Register multiple models with a type ID and parse/compile from a single reader/writer.
- **Control over layout** -- Choose byte order (e.g. `binary.LittleEndian`) for integers and floats; pick header types for length-prefixed data (e.g. `UInt8Header()` for up to 255 elements).
- **Reusable types** -- Build a type once (e.g. `parco.Int(binary.LittleEndian)`), then use it in structs, slices, maps, or standalone.
- **Optional factory** -- Use `ObjectFactory[T]()` for default zero values or `PooledFactory` for pooling instances when parsing.


## When to use Parco?

**Use Parco when:**
- You control both client and server (shared schema)
- Performance and bandwidth are critical
- You want compile-time safety (no reflection)
- Working with IoT, microservices, or game servers
- Need predictable memory usage (no hidden allocations)

**Consider alternatives when:**
- Schema changes frequently without versioning
- Need human-readable debug output
- Third-party integrations require JSON
- Dynamic/exploratory data analysis

In other words: if you're building a public REST API that strangers will consume, Parco is probably overkill (and arguably the wrong tool). If you're building a game server where every millisecond counts and you control the client, Parco makes sense. Choose your battles.

## Why Parco instead of Protocol Buffers or Cap'n Proto?

Fair question. Protobuf and Cap'n Proto are mature, battle-tested, and widely used. Here's why you might still choose Parco:

**When Parco makes sense:**

**Pure Go projects**: If your entire stack is Go and you don't need cross-language support, Protobuf's complexity (`.proto` files, `protoc` compiler, generated code, versioning generated files) is unnecessary overhead. Parco is just Go code.

**No codegen pipeline**: Protobuf requires writing `.proto` files, running `protoc`, managing generated code, and keeping it in sync with your actual code. Parco's builder API lives directly in your Go code. Change your schema, change your builder, done.

**Learning curve**: Proto syntax is another DSL to learn. Parco is just Go--if you know Go, you know Parco. No new syntax, no `protoc` flags, no plugin versions to juggle.

**Full Go flexibility**: Protobuf limits you to what the Proto language supports. Want custom serialization logic? Conditional fields? Complex validation inline? With Parco, you have the full power of Go (closures, interfaces, generics) directly in your builder.

**Zero dependencies (core API)**: Parco's core has zero external dependencies. Protobuf requires the protobuf runtime library, and Cap'n Proto needs its own runtime. This matters for small binaries, embedded systems, or when you just want fewer moving parts.

**Tighter control**: With Protobuf, you're at the mercy of what the code generator produces. With Parco, you write exactly what you want. Want to serialize a field conditionally based on runtime state? Just write the code.

**When Protobuf/Cap'n Proto make more sense:**

**Cross-language support**: If you need to talk to Rust, Python, C++, or any non-Go language, Protobuf is the obvious choice. Parco is Go-only and will stay that way.

**Industry standard**: Protobuf is used by Google, gRPC, and thousands of companies. It's proven at massive scale. Parco is a hobby project by comparison.

**Tooling ecosystem**: Protobuf has validators, linters, documentation generators, schema registries, and IDE plugins. Parco has... this README.

**Schema evolution guarantees**: Protobuf's backward/forward compatibility rules are well-documented and enforced by the toolchain. With Parco, you're responsible for not breaking things.

**Team size**: If you have multiple teams working on different services in different languages, Protobuf's schema-as-contract model makes sense. If it's just you or a small Go team, Parco's simpler.

**Realistic comparison:**

Think of Parco as the "write SQL queries directly" approach, while Protobuf is the "use an ORM" approach. ORMs (Protobuf) give you safety, abstraction, and cross-database (cross-language) support. Raw SQL (Parco) gives you control, performance, and simplicity when you don't need the abstraction.

Neither is universally better. Choose based on your constraints:
- Pure Go + want simplicity → Parco
- Multi-language + need tooling → Protobuf
- Maximum performance + control → Cap'n Proto (zero-copy) or Parco
- Industry standard + large team → Protobuf

**Quick decision matrix:**

| Your Situation | Recommendation |
|----------------|----------------|
| Pure Go, small team, want simplicity | **Parco** |
| Multi-language microservices | **Protocol Buffers** |
| Go + some Python/Rust/JS | **Protocol Buffers** |
| Maximum performance, Go-only, willing to manage schema manually | **Parco** |
| Need industry standard with proven tooling | **Protocol Buffers** |
| Want zero-copy, can handle complexity | **FlatBuffers** or **Cap'n Proto** |
| Public API, need flexibility | **JSON** or **MessagePack** |
| Debugging/logging/config files | **JSON** or **YAML** |

**Bottom line**: If you're already using Protobuf successfully, there's probably no reason to switch. If you're starting fresh on a pure Go project and find Protobuf's machinery overkill, give Parco a look.


## Installation

```bash
go get github.com/sonirico/parco
```

Requires **Go 1.19+**.


## Concepts

| Concept | Description |
|--------|-------------|
| **Type** | A `Type[T]` knows how to `Parse` and `Compile` values of type `T` (e.g. `UInt8()`, `Int(binary.LittleEndian)`). |
| **Builder** | `Builder[T](factory)` defines the layout of a struct (or model) by adding fields. Each field has a getter (compile) and setter (parse). |
| **Parser** | Produced by the builder; reads bytes from an `io.Reader` and fills a `T` (using the factory to create the value). |
| **Compiler** | Produced by the builder; writes a `T` to an `io.Writer`. |
| **Factory** | `Factory[T]` provides new instances when parsing (e.g. `ObjectFactory[T]()` returns zero value; you can use a pool for reuse). |
| **Header** | For variable-length data (slices, maps, varchar), a “header” type encodes the length (e.g. `UInt8Header()` for 0–255, `UInt16HeaderLE()` for 0–65535). |

You typically build once, then reuse the same parser and compiler for many reads/writes.


## Quick start

```go
package main

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/sonirico/parco"
)

func main() {
	type Point struct {
		X, Y int32
	}

	factory := parco.ObjectFactory[Point]()
	parser, compiler := parco.Builder[Point](factory).
		Int32(binary.LittleEndian,
			func(p *Point) int32 { return p.X },
			func(p *Point, x int32) { p.X = x },
		).
		Int32(binary.LittleEndian,
			func(p *Point) int32 { return p.Y },
			func(p *Point, y int32) { p.Y = y },
		).
		Parco()

	value := Point{X: 1, Y: 2}
	var buf bytes.Buffer
	_ = compiler.Compile(value, &buf)

	parsed, _ := parser.Parse(&buf)
	log.Println(parsed) // {1 2}
}
```


## Usage

### Parser & compiler (single model)

Define a model and use the builder to add fields in order. Call `Parco()` to obtain the parser and compiler.

```go
package main

import (
  "bytes"
  "encoding/binary"
  "encoding/json"
  "log"
  "reflect"
  "time"

  "github.com/sonirico/parco"
)

type (
  Animal struct {
    Age    uint8
    Specie string
  }

  Example struct {
    Greet              string
    LifeSense          uint8
    Friends            []string
    Grades             map[string]uint8
    EvenOrOdd          bool
    Pet                Animal
    Pointer            *int
    Flags              [5]bool
    Balance            float32
    MorePreciseBalance float64
    CreatedAt          time.Time
  }
)

func (e Example) String() string {
  bts, _ := json.MarshalIndent(e, "", "\t")
  return string(bts)
}

func main() {
  animalBuilder := parco.Builder[Animal](parco.ObjectFactory[Animal]()).
    SmallVarchar(
      func(a *Animal) string { return a.Specie },
      func(a *Animal, specie string) { a.Specie = specie },
    ).
    UInt8(
      func(a *Animal) uint8 { return a.Age },
      func(a *Animal, age uint8) { a.Age = age },
    )

  exampleFactory := parco.ObjectFactory[Example]()

  exampleParser, exampleCompiler := parco.Builder[Example](exampleFactory).
    SmallVarchar(
      func(e *Example) string { return e.Greet },
      func(e *Example, s string) { e.Greet = s },
    ).
    UInt8(
      func(e *Example) uint8 { return e.LifeSense },
      func(e *Example, lifeSense uint8) { e.LifeSense = lifeSense },
    ).
    Map(
      parco.MapField[Example, string, uint8](
        parco.UInt8Header(),
        parco.SmallVarchar(),
        parco.UInt8(),
        func(s *Example, grades map[string]uint8) { s.Grades = grades },
        func(s *Example) map[string]uint8 { return s.Grades },
      ),
    ).
    Slice(
      parco.SliceField[Example, string](
        parco.UInt8Header(),  // up to 255 items
        parco.SmallVarchar(), // each item's type
        func(e *Example, friends parco.SliceView[string]) { e.Friends = friends },
        func(e *Example) parco.SliceView[string] { return e.Friends },
      ),
    ).
    Bool(
      func(e *Example) bool { return e.EvenOrOdd },
      func(e *Example, evenOrOdd bool) { e.EvenOrOdd = evenOrOdd },
    ).
    Struct(
      parco.StructField[Example, Animal](
        func(e *Example) Animal { return e.Pet },
        func(e *Example, a Animal) { e.Pet = a },
        animalBuilder,
      ),
    ).
    Option(
      parco.OptionField[Example, int](
        parco.Int(binary.LittleEndian),
        func(e *Example, value *int) { e.Pointer = value },
        func(e *Example) *int { return e.Pointer },
      ),
    ).
    Array(
      parco.ArrayField[Example, bool](
        5,
        parco.Bool(),
        func(e *Example, flags parco.SliceView[bool]) {
          copy(e.Flags[:], flags)
        },
        func(e *Example) parco.SliceView[bool] {
          return e.Flags[:]
        },
      ),
    ).
    Float32(
      binary.LittleEndian,
      func(e *Example) float32 {
        return e.Balance
      },
      func(e *Example, balance float32) {
        e.Balance = balance
      },
    ).
    Float64(
      binary.LittleEndian,
      func(e *Example) float64 {
        return e.MorePreciseBalance
      },
      func(e *Example, balance float64) {
        e.MorePreciseBalance = balance
      },
    ).
    TimeUTC(
      func(e *Example) time.Time {
        return e.CreatedAt
      },
      func(e *Example, createdAt time.Time) {
        e.CreatedAt = createdAt
      },
    ).
    Parco()

  ex := Example{
    Greet:              "hey",
    LifeSense:          42,
    Grades:             map[string]uint8{"math": 5, "english": 6},
    Friends:            []string{"@boliri", "@danirod", "@enrigles", "@f3r"},
    EvenOrOdd:          true,
    Pet:                Animal{Age: 3, Specie: "cat"},
    Pointer:            parco.Ptr(73),
    Flags:              [5]bool{true, false, false, true, false},
    Balance:            234.987,
    MorePreciseBalance: 1234243.5678,
    CreatedAt:          time.Now().UTC(),
  }

  output := bytes.NewBuffer(nil)
  if err := exampleCompiler.Compile(ex, output); err != nil {
    log.Fatal(err)
  }

  log.Println(parco.FormatBytes(output.Bytes()))

  parsed, err := exampleParser.ParseBytes(output.Bytes())

  if err != nil {
    log.Fatal(err)
  }

  log.Println(parsed.String())

  if !reflect.DeepEqual(ex, parsed) {
    panic("not equals")
  }
}
```

### Single types

You can use primitive and composite types without a full struct builder.

#### Integer

```go
func main() {
  intType := parco.Int(binary.LittleEndian)
  buf := bytes.NewBuffer(nil)
  _ = intType.Compile(math.MaxInt, buf)
  n, _ := intType.Parse(buf)
  log.Println(n == math.MaxInt)
}
```

#### Slice of structs

```go
import (
  "bytes"
  "log"

  "github.com/sonirico/parco"
)

type (
  Animal struct {
    Age    uint8
    Specie string
  }
)

func main() {
  animalBuilder := parco.Builder[Animal](parco.ObjectFactory[Animal]()).
    SmallVarchar(
      func(a *Animal) string { return a.Specie },
      func(a *Animal, specie string) { a.Specie = specie },
    ).
    UInt8(
      func(a *Animal) uint8 { return a.Age },
      func(a *Animal, age uint8) { a.Age = age },
    )

  animalsType := parco.Slice[Animal](
    parco.UInt8Header(), // length prefix: up to 255 items
    parco.Struct[Animal](animalBuilder),
  )

  payload := []Animal{
    {Specie: "cat", Age: 32},
    {Specie: "dog", Age: 12},
  }

  buf := bytes.NewBuffer(nil)
  _ = animalsType.Compile(parco.SliceView[Animal](payload), buf)
  log.Println(buf.Bytes())

  res, _ := animalsType.Parse(buf)
  log.Println(res.Len())
  _ = res.Range(func(animal Animal) error {
    log.Println(animal)
    return nil
  })
}
```

### Multi-model (parsers & compilers)

Serialize and deserialize different models on the same stream by registering them with a type ID. The first field written/read is the model ID (e.g. from `UInt8Header()`), then the payload.

Models that implement `ParcoID() T` can use `Compile(item, w)`; otherwise use `CompileAny(id, item, w)`.

```go
type (
  Animal struct {
    Age    uint8
    Specie string
  }

  Flat struct {
    Price   float32
    Address string
  }
)

const (
  AnimalType int = 0
  FlatType       = 1
)

func (a Animal) ParcoID() int { return AnimalType }
func (a Flat) ParcoID() int   { return FlatType }

func main() {
  animalBuilder := parco.Builder[Animal](parco.ObjectFactory[Animal]()).
    SmallVarchar(
      func(a *Animal) string { return a.Specie },
      func(a *Animal, specie string) { a.Specie = specie },
    ).
    UInt8(
      func(a *Animal) uint8 { return a.Age },
      func(a *Animal, age uint8) { a.Age = age },
    )

  flatBuilder := parco.Builder[Flat](parco.ObjectFactory[Flat]()).
    Float32(
      binary.LittleEndian,
      func(f *Flat) float32 { return f.Price },
      func(f *Flat, price float32) { f.Price = price },
    ).
    SmallVarchar(
      func(f *Flat) string { return f.Address },
      func(f *Flat, address string) { f.Address = address },
    )

  parCo := parco.MultiBuilder(parco.UInt8Header()). // up to 255 model types
    MustRegister(AnimalType, animalBuilder).
    MustRegister(FlatType, flatBuilder)

  buf := bytes.NewBuffer(nil)

  // Compile when model implements ParcoID():
  _ = parCo.Compile(Animal{Age: 10, Specie: "monkeys"}, buf)
  _ = parCo.Compile(Flat{Price: 42, Address: "Plaza mayor"}, buf)

  // Or specify ID explicitly:
  _ = parCo.CompileAny(AnimalType, Animal{Age: 7, Specie: "felix catus"}, buf)

  id, something, _ := parCo.Parse(buf)
  Print(id, something)
  id, something, _ = parCo.Parse(buf)
  Print(id, something)
  id, something, _ = parCo.Parse(buf)
  Print(id, something)
}

func Print(id int, x any) {
  switch id {
  case AnimalType:
    log.Println("animal:", x.(Animal))
  case FlatType:
    log.Println("flat:", x.(Flat))
  }
}
```


## Supported types

| Field                 | Size                           |
|-----------------------|--------------------------------|
| byte                  | 1                              |
| int8                  | 1                              |
| uint8                 | 1                              |
| int16                 | 2                              |
| uint16                | 2                              |
| int32                 | 4                              |
| uint32                | 4                              |
| int64                 | 8                              |
| uint64                | 8                              |
| float32               | 4                              |
| float64               | 8                              |
| int                   | 4/8 (platform)                 |
| bool                  | 1                              |
| small varchar         | dyn (up to 255)                |
| varchar               | dyn (up to 65535)              |
| text                  | dyn (up to max uint32 chars)   |
| long text             | dyn (up to max uint64 chars)   |
| string                | dyn                            |
| bytes (blob)          | dyn                            |
| map                   | variable                       |
| slice                 | variable                       |
| array (fixed)         | length × element size          |
| struct                | sum of field sizes             |
| time.Time             | 8 (+ small varchar if TZ aware)|
| optional[T] (pointer) | 1 + inner size                 |

For length-prefixed types (slices, maps, varchars), use a **header** type for the length: `UInt8Header()`, `UInt16HeaderLE()`, `Int32LEHeader()`, etc., depending on the maximum count you need.


## Error handling

The package uses sentinel errors and a few custom types you can check with `errors.Is` / type assertions:

| Error | Description |
|-------|-------------|
| `parco.ErrNotIntegerType` | Integer type assertion failed. |
| `parco.ErrOverflow` | Value exceeds the range of the header or type. |
| `parco.ErrCannotRead` | Fewer bytes read than required. |
| `parco.ErrCannotWrite` | Fewer bytes written than required. |
| `parco.ErrAlreadyRegistered` | Multi-model: type ID already registered. |
| `parco.ErrUnknownType` | Multi-model: unknown type ID when parsing, or `CompileAny` with wrong type. |
| `parco.ErrUnSufficientBytes` | Not enough bytes (want/have). |
| `parco.ErrFieldNotFound` | Field lookup failed. |
| `parco.ErrTypeAssertion` | Type assertion failed (expected/actual). |
| `parco.ErrCompile` | Generic compile-time error (reason string). |


## Examples

Run the examples from the repo root:

| Example | Description |
|--------|-------------|
| [examples/builder](examples/builder) | Full builder API: parser + compiler from a single builder. |
| [examples/compiler](examples/compiler) | Compiler-only usage (no parser). |
| [examples/parser](examples/parser) | Parser-only usage (no compiler). |
| [examples/registry](examples/registry) | Multi-model: several types on the same stream with `MultiBuilder`. |
| [examples/registry_singleton](examples/registry_singleton) | Global multi-model registry. |

```bash
go run ./examples/registry/
```


## Benchmarks

Run benchmarks:

```bash
# Using just
just bench

# Direct go test
go test -bench=. -benchmem
```

### Performance Comparison

Parco vs JSON vs MessagePack on an Intel i7-8750H @ 2.20GHz:

| Format | Small (91 bytes) | Medium (742 bytes) | Large (8123 bytes) |
|--------|-----------------|-------------------|-------------------|
| **Parco** | **2373 ns/op** <br> 91 bytes <br> 3 allocs | **15727 ns/op** <br> 742 bytes <br> 3 allocs | **154303 ns/op** <br> 8123 bytes <br> 3 allocs |
| JSON | 2972 ns/op (1.25x slower) <br> 269 bytes (2.96x larger) <br> 23 allocs (7.6x more) | 28305 ns/op (1.8x slower) <br> 1681 bytes (2.27x larger) <br> 203 allocs (67x more) | 319851 ns/op (2.07x slower) <br> 16637 bytes (2.05x larger) <br> 2003 allocs (667x more) |
| MessagePack | 3020 ns/op (1.27x slower) <br> 155 bytes (1.70x larger) <br> 25 allocs (8.3x more) | 20051 ns/op (1.27x slower) <br> 991 bytes (1.34x larger) <br> 207 allocs (69x more) | 206755 ns/op (1.34x slower) <br> 10171 bytes (1.25x larger) <br> 2007 allocs (669x more) |

**Key Takeaways:**
- Consistently faster: 25-100% faster than alternatives
- Minimal allocations: Only 3 allocations regardless of payload size
- Compact payloads: Up to 3x smaller than JSON
- Scales linearly: Performance degrades predictably with size

### Why is Parco faster?

**Design Philosophy: Known Schema vs Self-Describing**

| Aspect | Parco | JSON/MessagePack |
|--------|-------|------------------|
| **Schema** | Known at compile time | Embedded in each message |
| **Field names** | Not transmitted | Included in payload |
| **Type info** | Pre-defined | Encoded per value |
| **Use case** | Client-server with shared schema | Dynamic/exploratory data |
| **Similar to** | Protocol Buffers, FlatBuffers | Self-describing formats |

**This is not a "trick" - it's a deliberate trade-off:**
- If sender and receiver know the schema → Parco saves bandwidth/CPU
- If schema is unknown or changes frequently → JSON/MessagePack are more flexible

Think of it like HTTP/1.1 (text headers) vs HTTP/2 (binary with HPACK compression) - both valid, different use cases.

### Understanding These Benchmarks (or: Why You Should Be Skeptical)

These results might seem "too good to be true"--and in some sense, they are. Let me explain what Parco doesn't do, by design, which is why it's faster:

**What Parco Does NOT Include:**

**No field names in the wire format**
JSON sends `{"name":"Alice","age":30}` with the actual strings `"name"` and `"age"` embedded. Parco just sends the values. Sender and receiver must agree on field order beforehand, or things will break spectacularly.

**No type information per value**
JSON includes type metadata for every value (string, number, array, etc.). Parco assumes you defined the types at compile time and doesn't waste bytes transmitting them.

**No schema discovery**
You can `curl` a random JSON endpoint and understand the structure. With Parco, you need the schema definition or you're just staring at binary gibberish.

**No reflection at runtime**
JSON uses reflection to inspect structs (slow but flexible). Parco uses a builder API that generates direct code paths at compile time (fast but rigid).

#### What This Means:

```go
// JSON payload example (self-describing):
{
  "user": {
    "id": 42,
    "name": "Alice",
    "active": true
  }
}
// Size: ~60 bytes, includes all metadata

// Parco payload (binary, schema known):
[42]["Alice"][1]
// Size: ~8 bytes, just the raw data
```

**When Parco Makes Sense:**

**Microservices**: Both ends control the schema, deploy together.

**Game networking**: Performance critical, schema updates via patches. Nobody wants to send field names when you're already fighting latency.

**IoT devices**: Bandwidth constrained, firmware includes schema. Battery life matters.

**High-frequency trading**: Every microsecond matters, and you're probably already using FIX or worse.

**Internal APIs**: Teams coordinate schema changes. If breaking changes require a meeting, you have bigger problems than serialization.

**When JSON/MessagePack Make More Sense:**

**Public REST APIs**: Unknown clients, self-documentation expected. You want developers to `curl | jq` your API and understand it.

**Configuration files**: Human readable, easy to edit. Good luck hand-editing binary files.

**Debugging/logging**: Need to inspect data without specialized tools. `cat log.json` beats `xxd log.bin | head -c 1000`.

**Polyglot systems**: Multiple languages without shared schema. Convincing your Rust/Python/JS teams to use the same schema definition is hard enough without binary formats.

**Rapid prototyping**: Schema changes frequently. If your schema changes three times a day, schema-based serialization will drive you insane.

**The Honest Truth:**

Parco is not universally better than JSON. It's optimized for a specific use case where:
- You control both client and server
- Performance matters more than flexibility
- Schema is relatively stable or versioned
- Binary format is acceptable

If that's not you, don't use Parco. Seriously.

**If you need JSON's flexibility but better performance:**
- **MessagePack**: Good middle ground (still self-describing but more compact)
- **JSON with schema validation**: Use JSON Schema for safety, keep the flexibility
- **CBOR**: Like MessagePack but RFC-standardized (if you care about that)

**If you want Parco-level performance with wider tooling:**
- **Protocol Buffers**: Industry standard, multi-language, mature tooling. Trade-off: `.proto` files, codegen pipeline, learning curve.
- **FlatBuffers**: Zero-copy deserialization, very fast. Trade-off: more complex to use, schema must be stable.
- **Cap'n Proto**: Protobuf's successor by its original author, simpler model, fast. Trade-off: smaller ecosystem than Protobuf.

**Bottom line**: These benchmarks compare different design philosophies honestly. Parco trades flexibility for speed. If that trade works for your use case, great. If not, there are plenty of other tools that might suit you better. Choose wisely.


## Development

Parco uses [just](https://github.com/casey/just) for task automation. If you don't have it: `cargo install just`

Common commands:

| Command | Description |
|--------|-------------|
| `just test` | Run tests with race detector |
| `just test-coverage` | Generate HTML coverage report |
| `just bench` | Run benchmarks |
| `just bench-profile` | Profile CPU and memory |
| `just lint` | Run golangci-lint |
| `just format` | Format code |
| `just ci` | Run all checks (test + lint) |
| `just setup` | Install dev tools |
| `just example builder` | Run specific example |

See `just --list` for all available commands.


## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

Quick start:
```bash
# Setup development tools
just setup

# Run tests
just test

# Run benchmarks
just bench

# Run all checks
just ci
```


## Roadmap

**Short-term**
- Comprehensive test coverage (41.8% and climbing)
- Memory safety improvements (done: limits on allocations)
- Schema evolution utilities
- Validation helpers

**Long-term**
- Static code generation from struct tags or DSL
- Replace `encoding/binary` with faster, zero-alloc primitives
- Custom `Reader`/`Writer` interfaces for single-byte operations
- Cross-language support (C, Rust bindings)

Contributions welcome, but please open an issue before starting significant work.


## Additional Resources

- **[Performance Guide](PERFORMANCE.md)** - Detailed benchmark analysis and optimization tips
- **[Contributing Guide](CONTRIBUTING.md)** - How to contribute to Parco
- **[Examples](examples/)** - Practical usage examples
- **[Go Package Docs](https://pkg.go.dev/github.com/sonirico/parco)** - API reference


## License

MIT License - see [LICENSE](LICENSE) for details.


## Acknowledgments

Thanks to all [contributors](https://github.com/sonirico/parco/graphs/contributors).
