# Parco

Hobbyist binary compiler and parser built with no reflection, highly
extensible, focused on performance, usability through generics and
with zero dependencies.

There are plenty packages over the internet which work by leveraging the power of
struct tags and reflection. While sometimes that can be convenient for some
scenarios, that approach leaves little room to define and register custom types in
addition to have an appositive effect on performance.

## Usage

#### Parser

```go
type Student struct {
    Name   string
    Age    uint8
    Grades []uint8
}

data := []byte{4, 72, 79, 76, 65, 42, 2, 9, 10}

studentFactory := parco.ObjectFactory[Student]()

parser := parco.ParserModel[Student](studentFactory).
    SmallVarchar(func(s *Student, name string) {
        s.Name = name
    }).
    UInt8(func(s *Student, age uint8) {
        s.Age = age
    }).
    Array(parco.ArrayFieldSetter(
        parco.UInt8Header(), // up to 255 items in the array
        parco.UInt8(), // type of the array
        func(s *Student, items parco.Slice[uint8]) {
            s.Grades = items
        },
    ))

parsed, err := parser.ParseBytes(data)

if err != nil {
    log.Fatal(err)
}

log.Println(parsed.Name)
log.Println(parsed.Age)
log.Println(parsed.Grades)

```

#### Compiler

```go
type Example struct {
    Greet     string
    LifeSense uint8
    Grades    []uint8
    Friends   []string
}

compiler := parco.CompilerModel[Example]().
    SmallVarchar(func(e *Example) string {
        return e.Greet
    }).
    UInt8(func(e *Example) uint8 {
        return e.LifeSense
    }).
    Array(
        parco.ArrayFieldGetter[Example, uint8](
            parco.UInt8Header(), // up to 255 items
            parco.UInt8(),       // each item
            func(e *Example) parco.Slice[uint8] {
                return e.Grades
            },
        ),
    ).
    Array(
        parco.ArrayFieldGetter[Example, string](
            parco.UInt8Header(),  // up to 255 items
            parco.SmallVarchar(), // each item
            func(e *Example) parco.Slice[string] {
                return e.Friends
            }, 
        ), 
    )

ex := Example{
    Greet:     "hey",
    LifeSense: 42,
    Grades:    []uint8{5, 6},
    Friends:   []string{"@boliri", "@danirod", "@enrigles", "@f3r"},
}

output := bytes.NewBuffer(nil)
if err := compiler.Compile(ex, output); err != nil {
    log.Fatal(err)
}

log.Println(output.Bytes())

```


#### Builder

Both parser and compiler can be defined at the same time:

```go
exampleFactory := parco.ObjectFactory[Example]()

parser, compiler := parco.Builder[Example](exampleFactory).
    SmallVarchar(
        func(e *Example) string {
            return e.Greet
        },
        func(e *Example, s string) {
            e.Greet = s
        },
    ).
    UInt8(
        func(e *Example) uint8 {
            return e.LifeSense
        },
        func(e *Example, lifeSense uint8) {
            e.LifeSense = lifeSense
        },
    ).
    Array(
        parco.ArrayField[Example, uint8](
            parco.UInt8Header(), // up to 255 items
            parco.UInt8(),       // each item's type
            func(e *Example, grades parco.Slice[uint8]) {
                e.Grades = grades
            },
            func(e *Example) parco.Slice[uint8] {
                return e.Grades
            },
        ),
    ).
    Array(
        parco.ArrayField[Example, string](
            parco.UInt8Header(),  // up to 255 items
            parco.SmallVarchar(), // each item's type
            func(e *Example, friends parco.Slice[string]) {
                e.Friends = friends
            },
            func(e *Example) parco.Slice[string] {
                return e.Friends
            },
        ),
    ).
    ParCo()

ex := Example{
    Greet:     "hey",
    LifeSense: 42,
    Grades:    []uint8{5, 6},
    Friends:   []string{"@boliri", "@danirod", "@enrigles", "@f3r"},
}

output := bytes.NewBuffer(nil)
if err := compiler.Compile(ex, output); err != nil {
    log.Fatal(err)
}

log.Println(output.Bytes())

parsed, err := parser.ParseBytes(output.Bytes())

if err != nil {
    log.Fatal(err)
}

if !ex.Equals(parsed) {
    panic("not equals")
}
```

For fully functional examples showing the whole API, refer to [Examples](https://github.com/sonirico/parco/tree/master/examples).


## Benchmarks

```shell
make bench

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkParcoAlloc_Compile
BenchmarkParcoAlloc_Compile/small_size
BenchmarkParcoAlloc_Compile/small_size-12                1627336               771.7 ns/op              47.00 payload_bytes/op       184 B/op          3 allocs/op
BenchmarkParcoAlloc_Compile/medium_size
BenchmarkParcoAlloc_Compile/medium_size-12                299202              3812 ns/op               338.0 payload_bytes/op        184 B/op          3 allocs/op
BenchmarkParcoAlloc_Compile/large_size
BenchmarkParcoAlloc_Compile/large_size-12                  34354             35841 ns/op              3218 payload_bytes/op          184 B/op          3 allocs/op
BenchmarkParcoDiscard_Compile
BenchmarkParcoDiscard_Compile/small_size
BenchmarkParcoDiscard_Compile/small_size-12              1625402               704.0 ns/op              47.00 payload_bytes/op       184 B/op          3 allocs/op
BenchmarkParcoDiscard_Compile/medium_size
BenchmarkParcoDiscard_Compile/medium_size-12              353724              3478 ns/op               338.0 payload_bytes/op        184 B/op          3 allocs/op
BenchmarkParcoDiscard_Compile/large_size
BenchmarkParcoDiscard_Compile/large_size-12                39127             31379 ns/op              3218 payload_bytes/op          184 B/op          3 allocs/op
BenchmarkJson_Compile
BenchmarkJson_Compile/small_size
BenchmarkJson_Compile/small_size-12                      1935265               690.0 ns/op             116.0 payload_bytes/op        192 B/op          2 allocs/op
BenchmarkJson_Compile/medium_size
BenchmarkJson_Compile/medium_size-12                      367641              3206 ns/op               756.0 payload_bytes/op        832 B/op          2 allocs/op
BenchmarkJson_Compile/large_size
BenchmarkJson_Compile/large_size-12                        42553             28492 ns/op              7071 payload_bytes/op         8262 B/op          2 allocs/op
BenchmarkMsgpack_Compile
BenchmarkMsgpack_Compile/small_size
BenchmarkMsgpack_Compile/small_size-12                   1296073               918.7 ns/op              74.00 payload_bytes/op       320 B/op          4 allocs/op
BenchmarkMsgpack_Compile/medium_size
BenchmarkMsgpack_Compile/medium_size-12                   237603              4778 ns/op               458.0 payload_bytes/op        944 B/op          5 allocs/op
BenchmarkMsgpack_Compile/large_size
BenchmarkMsgpack_Compile/large_size-12                     28405             40949 ns/op              4238 payload_bytes/op         9651 B/op          6 allocs/op

```

## TODO

- Support for all primitive types: boolean, nil...
- Extend interface to include version
- Static code generation
- Replace `encoding/binary` usage by faster implementations (`WriteByte`)
- Custom `Reader` and `Writer` interfaces to implement single byte ops
- Support for nested schema definitions.