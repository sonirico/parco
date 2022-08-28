![build](https://github.com/sonirico/parco/actions/workflows/go.yml/badge.svg)
# Parco

Hobbyist binary compiler and parser built with no reflection, highly
extensible, focused on performance, usability through generics and
with zero dependencies.

There are plenty packages over the internet which work by leveraging the power of
struct tags and reflection. While sometimes that can be convenient for some
scenarios, that approach leaves little room to define and register custom types in
addition to have an appositive effect on performance.

## Usage

### Parser & compiler

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

#### Integer

```go
func main () {
  intType := parco.Int(binary.LittleEndian)
  buf := bytes.NewBuffer(nil)
  _ = intType.Compile(math.MaxInt, buf)
  n, _ := intType.Parse(buf)
  log.Println(n == math.MaxInt)
}
```

#### Slice of structs


```go
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
    intType,
    parco.Struct[Animal](animalBuilder),
  )

  payload := []Animal{
    {
      Specie: "cat",
      Age:    32,
    },
    {
      Specie: "dog",
      Age:    12,
    },
  }

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

---

### Supported fields

| Field                 | Status | Size                           |
|-----------------------|--------|--------------------------------|
| byte                  | ✅      | 1                              |
| int8                  | ✅      | 1                              |
| uint8                 | ✅      | 1                              |
| int16                 | ✅      | 2                              |
| uint16                | ✅      | 2                              |
| int32                 | ✅      | 4                              |
| uint32                | ✅      | 4                              |
| int64                 | ✅      | 8                              |
| uint64                | ✅      | 8                              |
| float32               | ✅      | 4                              |
| float64               | ✅      | 8                              |
| int                   | ✅      | 4/8                            |
| bool                  | ✅      | 1                              |
| small varchar         | ✅      | dyn (up to 255)                |
| varchar               | ✅      | dyn (up to 65535)              |
| text                  | ✅      | dyn (up to max uint32 chars)   |
| long text             | ✅      | dyn (up to max uint64 chars)   |
| string                | ✅      | dyn                            |
| bytes (blob)          | ✅      | dyn                            |
| map                   | ✅      | -                              |
| slice                 | ✅      | -                              |
| array (fixed)         | ✅      | -                              |
| struct                | ✅      | -                              |
| time.Time             | ✅      | 8 (+small varchar if TZ aware) |
| optional[T] (pointer) | ✅      | 1 + inner size                 |

For fully functional examples showing the whole API, refer to [Examples](https://github.com/sonirico/parco/tree/master/examples).


## Benchmarks

```shell
make bench

goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
ParcoAlloc_Compile
ParcoAlloc_Compile/small_size
ParcoAlloc_Compile/small_size-12    276934       4015 ns/op      91.00 payload_bytes/op    237 B/op     5 allocs/op
ParcoAlloc_Compile/medium_size
ParcoAlloc_Compile/medium_size-12    48273      24906 ns/op     742.0 payload_bytes/op     239 B/op     5 allocs/op
ParcoAlloc_Compile/large_size
ParcoAlloc_Compile/large_size-12      4705     247203 ns/op    8123 payload_bytes/op       245 B/op     5 allocs/op
ParcoDiscard_Compile
ParcoDiscard_Compile/small_size
ParcoDiscard_Compile/small_size-12  322285       3741 ns/op      91.00 payload_bytes/op    238 B/op     5 allocs/op
ParcoDiscard_Compile/medium_size
ParcoDiscard_Compile/medium_size-12  50703      23336 ns/op     742.0 payload_bytes/op     238 B/op     5 allocs/op
ParcoDiscard_Compile/large_size
ParcoDiscard_Compile/large_size-12    5406     220967 ns/op    8123 payload_bytes/op       241 B/op     5 allocs/op
Json_Compile
Json_Compile/small_size
Json_Compile/small_size-12          213540       5410 ns/op     270.0 payload_bytes/op    1330 B/op    26 allocs/op
Json_Compile/medium_size
Json_Compile/medium_size-12          23980      49912 ns/op    1680 payload_bytes/op     10256 B/op   206 allocs/op
Json_Compile/large_size
Json_Compile/large_size-12            2014     581209 ns/op   16598 payload_bytes/op    101265 B/op  2006 allocs/op
Msgpack_Compile
Msgpack_Compile/small_size
Msgpack_Compile/small_size-12       242005       4760 ns/op     155.0 payload_bytes/op     762 B/op    25 allocs/op
Msgpack_Compile/medium_size
Msgpack_Compile/medium_size-12       33638      35485 ns/op     991.0 payload_bytes/op    4069 B/op   207 allocs/op
Msgpack_Compile/large_size
Msgpack_Compile/large_size-12         3277     357921 ns/op   10171 payload_bytes/op     37448 B/op  2007 allocs/op
```

## Roadmap

- Static code generation.
- Replace `encoding/binary` usage by faster implementations, such as `WriteByte` in order to achieve a zero alloc implementation.
- Custom `Reader` and `Writer` interfaces to implement single byte ops.
