# Parco

Hobbyist binary compiler and parser built with as less reflection as possible, highly
extensible and with zero dependencies.

There are plenty packages over the internet which work by leveraging the power of
struct tags and reflection. While sometimes that can be convenient for some
scenarios, that approach leaves little room to define and register custom types in
addition to have an appositive effect on performance.

In spite of that, due to the lack of generics, this package still uses reflection
for iterable interfaces in order for types to adhere to external interfaces. Do note
that `unsafe` is also employed (quite isolated though).

## Usage

So the most complete usage would look like this


### Parser + compiler

```go
package main

import (
    "bytes"
    types "github.com/sonirico/parco/internal"
    parco "github.com/sonirico/parco/pkg"
    "log"
)

type Example struct {
    Greet     string
    LifeSense int
    Grades    []uint8
}

func getGreet(x interface{}) interface{} {
    return x.(Example).Greet
}

func getLifeSense(x interface{}) interface{} {
    return x.(Example).LifeSense
}

func getGrades(x interface{}) interface{} {
    return x.(Example).Grades
}

func main() {
    parser, compiler := parco.NewBuilder().
        FieldGet("greet", types.SmallVarchar(), getGreet).
        FieldGet("life_sense", types.UInt8(), getLifeSense).
        FieldGet("grades", types.Array(2, types.UInt8(), types.UInt8()), getGrades).
        ParCo()

    ex := Example{
        Greet:     "hey",
        LifeSense: 42,
        Grades:    []uint8{5, 6},
    }

    output := bytes.NewBuffer(nil)
    if err := compiler.Compile(ex, output); err != nil {
        log.Fatal(err)
    }

    raw := output.Bytes()
    log.Println("raw bytes", raw)

    result, err := parser.ParseBytes(raw)
    if err != nil {
        log.Fatal(err)
    }

    greet, _ := result.GetString("greet")
    lifeSense, _ := result.GetUint8("life_sense")
    grades, _ := result.GetArray("grades")

    log.Println("greet", greet)
    log.Println("life sense", lifeSense)
    log.Println("total grades", grades.Len())

    grades.Range(func(value types.Value) {
        log.Println("grade", value.GetUInt8())
    })
}

```

However, both parser and compiler can be used independently.

#### Parser

```go
raw := []byte{4, 72, 79, 76, 65, 42, 9, 10}

parser := parco.NewBuilder().
    Field("greet", types.SmallVarchar()).
    Field("life_sense", types.UInt8()).
    Field("grades", types.Array(2, types.UInt8(), types.UInt8())).
    Parser()


result, err := parser.ParseBytes(raw)
if err != nil {
    log.Fatal(err)
}

greet, _ := result.GetString("greet")
lifeSense, _ := result.GetUInt8("life_sense")
grades, _ := result.GetArray("grades")

log.Println("greet", greet)
log.Println("life sense", lifeSense)
log.Println("total grades", grades.Len())

grades.Range(func(value types.Value) {
    log.Println("grade", value.GetUInt8())
})

```

#### Compiler

```go
type Example struct {
    Greet     string
    LifeSense int
    Grades    []uint8
}

func getGreet(x interface{}) interface{} {
    return x.(Example).Greet
}

func getLifeSense(x interface{}) interface{} {
    return x.(Example).LifeSense
}

func getGrades(x interface{}) interface{} {
    return x.(Example).Grades
}

func main() {
    compiler := parco.NewBuilder().
        FieldGet("greet", types.SmallVarchar(), getGreet).
        FieldGet("life_sense", types.UInt8(), getLifeSense).
        FieldGet("grades", types.Array(2, types.UInt8(), types.UInt8()), getGrades).
        Compiler()

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
}

```

For fully functional examples showing the whole API, refer to [Examples](https://github.com/sonirico/parco/tree/master/examples).


## Benchmarks

```
goos: linux
goarch: amd64
pkg: github.com/sonirico/parco/pkg
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkParco_Compile
BenchmarkParco_Compile/small_size
BenchmarkParco_Compile/small_size-12         	1000000000	         0.0000043 ns/op	        35.00 bytes/op
BenchmarkParco_Compile/medium_size
BenchmarkParco_Compile/medium_size-12        	1000000000	         0.0000160 ns/op	       325.0 bytes/op
BenchmarkParco_Compile/large_size
BenchmarkParco_Compile/large_size-12         	1000000000	         0.0000568 ns/op	      3206 bytes/op
BenchmarkJson_Compile
BenchmarkJson_Compile/small_size
BenchmarkJson_Compile/small_size-12          	1000000000	         0.0000042 ns/op	       116.0 bytes/op
BenchmarkJson_Compile/medium_size
BenchmarkJson_Compile/medium_size-12         	1000000000	         0.0000159 ns/op	       756.0 bytes/op
BenchmarkJson_Compile/large_size
BenchmarkJson_Compile/large_size-12          	1000000000	         0.0000744 ns/op	      7071 bytes/op
PASS
ok  	github.com/sonirico/parco/pkg	0.010s

```


## TODO

- Support for all primitive types: boolean, nil...
- Extend interface to include version
- Static code generation
- Replace `encoding/binary` usage by faster implementations