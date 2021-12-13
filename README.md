# Parco

Hobbyist binary compiler and parser built with as less reflection as possible, highly
extensible and with zero dependencies.

There are plenty packages over the internet which work by leveraging the power of
struct tags and reflection. While sometimes that can be convenient for some
scenarios, that approach leaves little room to define and register custom types in
addition to have an appositive effect on performance.

Do note that `unsafe` is employed (quite isolated though, logging and debugging purposes).

## Usage

#### Parser

```go
data := []byte{4, 72, 79, 76, 65, 42, 2, 9, 10}

parser := parco.NewParserResult().
    SmallVarchar("greet").
    UInt8("life_sense").
    Array("grades", parco.AnyArray(
        parco.UInt8Header(),
        parco.AnyUInt8Body(),
        nil,
    ))

parsed, err := parser.ParseBytes(data)

if err != nil {
    log.Fatal(err)
}

log.Println(parsed.GetString("greet"))
log.Println(parsed.GetUInt8("life_sense"))

grades, err := parsed.GetArray("grades")

if err != nil {
    log.Fatal(err)
}

log.Println(grades.At(0))
log.Println(grades.At(1))

v, _ := grades.At(0)
log.Println(v.GetUInt8())

grades.Range(func(value parco.Value) {
    log.Println(value.GetUInt8())
})

```

#### Compiler

```go
type Example struct {
    Greet     string
    LifeSense uint8
    Grades    []uint8
}

compiler := parco.NewCompiler[Example]().
    SmallVarchar("greet", func(e Example) string {
        return e.Greet
    }).
    UInt8("life_sense", func(e Example) uint8 {
        return e.LifeSense
    }).
    Array("grades", parco.Array[Example, uint8](
        parco.UInt8Header(),
        parco.UInt8Body(),
        func(e Example) parco.Iterable[uint8] {
            return parco.UInt8Iter(e.Grades)
        }),
    )

ex := Example{
    Greet:     "hey",
    LifeSense: 42,
    Grades:    []uint8{5, 6},
}

output := bytes.NewBuffer(nil)
if err := compiler.Compile(ex, output); err != nil {
    log.Fatal(err)
}

log.Println(output.Bytes())

```

For fully functional examples showing the whole API, refer to [Examples](https://github.com/sonirico/parco/tree/master/examples).


## Benchmarks

```shell
make bench # or
CGO_ENABLED=1 go test -v -failfast -race -bench=. -benchtime=1000x -benchmem ./internal/... ./pkg/... -run . -timeout=1m
```

```
goos: linux
goarch: amd64
pkg: github.com/sonirico/parco
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkParcoAlloc_Compile
BenchmarkParcoAlloc_Compile/small_size
BenchmarkParcoAlloc_Compile/small_size-12                1741048               661.4 ns/op              47.00 payload_bytes/op       121 B/op          3 allocs/op
BenchmarkParcoAlloc_Compile/medium_size
BenchmarkParcoAlloc_Compile/medium_size-12                319216              3714 ns/op               338.0 payload_bytes/op        121 B/op          3 allocs/op
BenchmarkParcoAlloc_Compile/large_size
BenchmarkParcoAlloc_Compile/large_size-12                  34654             34305 ns/op              3218 payload_bytes/op          120 B/op          2 allocs/op
BenchmarkParcoDiscard_Compile
BenchmarkParcoDiscard_Compile/small_size
BenchmarkParcoDiscard_Compile/small_size-12              1945192               612.0 ns/op              47.00 payload_bytes/op       121 B/op          3 allocs/op
BenchmarkParcoDiscard_Compile/medium_size
BenchmarkParcoDiscard_Compile/medium_size-12              334830              3411 ns/op               338.0 payload_bytes/op        121 B/op          3 allocs/op
BenchmarkParcoDiscard_Compile/large_size
BenchmarkParcoDiscard_Compile/large_size-12                37971             31599 ns/op              3218 payload_bytes/op          120 B/op          2 allocs/op
BenchmarkJson_Compile
BenchmarkJson_Compile/small_size
BenchmarkJson_Compile/small_size-12                      1938532               607.2 ns/op             116.0 payload_bytes/op        192 B/op          2 allocs/op
BenchmarkJson_Compile/medium_size
BenchmarkJson_Compile/medium_size-12                      367490              3015 ns/op               756.0 payload_bytes/op        832 B/op          2 allocs/op
BenchmarkJson_Compile/large_size
BenchmarkJson_Compile/large_size-12                        44212             27532 ns/op              7071 payload_bytes/op         8263 B/op          2 allocs/op
BenchmarkMsgpack_Compile
BenchmarkMsgpack_Compile/small_size
BenchmarkMsgpack_Compile/small_size-12                   1399462               855.1 ns/op              74.00 payload_bytes/op       320 B/op          4 allocs/op
BenchmarkMsgpack_Compile/medium_size
BenchmarkMsgpack_Compile/medium_size-12                   264453              4176 ns/op               458.0 payload_bytes/op        944 B/op          5 allocs/op
BenchmarkMsgpack_Compile/large_size
BenchmarkMsgpack_Compile/large_size-12                     32216             36783 ns/op              4238 payload_bytes/op         9651 B/op          6 allocs/op
```

## TODO

- Support for all primitive types: boolean, nil...
- Extend interface to include version
- Static code generation
- Replace `encoding/binary` usage by faster implementations (`WriteByte`)
- Custom `Reader` and `Writer` interfaces to implement single byte ops
- Support for nested schema definition