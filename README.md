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
type (
	Animal struct {
		Age    uint8
		Specie string
	}

	Example struct {
		Greet     string
		LifeSense uint8
		Friends   []string
		Grades    map[string]uint8
		EvenOrOdd bool
		Pet       Animal
		Pointer   *int
		Flags     [5]bool
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
		Parco()

	ex := Example{
		Greet:     "hey",
		LifeSense: 42,
		Grades:    map[string]uint8{"math": 5, "english": 6},
		Friends:   []string{"@boliri", "@danirod", "@enrigles", "@f3r"},
		EvenOrOdd: true,
		Pet:       Animal{Age: 3, Specie: "cat"},
		Pointer:   parco.Ptr(73),
		Flags:     [5]bool{true, false, false, true, false},
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

| Field                 | Status | Size                         |
|-----------------------|--------|------------------------------|
| byte                  | ✅      | 1                            |
| int8                  | ✅      | 1                            |
| uint8                 | ✅      | 1                            |
| int16                 | ✅      | 2                            |
| uint16                | ✅      | 2                            |
| int32                 | ✅      | 4                            |
| uint32                | ✅      | 4                            |
| int64                 | ✅      | 8                            |
| uint64                | ✅      | 8                            |
| float32               | 👷🚧   | 4                            |
| float64               | 👷🚧   | 8                            |
| int                   | ✅      | 4/8                          |
| bool                  | ✅      | 1                            |
| small varchar         | ✅      | dyn (up to 255)              |
| varchar               | ✅      | dyn (up to 65535)            |
| text                  | ✅      | dyn (up to max uint32 chars) |
| long text             | ✅      | dyn (up to max uint64 chars) |
| string                | ✅      | dyn                          |
| bytes (blob)          | ✅      | dyn                          |
| map                   | ✅      | -                            |
| slice                 | ✅      | -                            |
| array (fixed)         | ✅      | -                            |
| struct                | ✅      | -                            |
| time.Time             | 👷🚧   | ?                            |
| optional[T] (pointer) | ✅      | 1 + inner size               |

For fully functional examples showing the whole API, refer to [Examples](https://github.com/sonirico/parco/tree/master/examples).


## Benchmarks

```shell
make bench

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz

ParcoAlloc_Compile
ParcoAlloc_Compile/small_size
ParcoAlloc_Compile/small_size-12                 674432              1777 ns/op                78.00 payload_bytes/op       200 B/op          3 allocs/op
ParcoAlloc_Compile/medium_size
ParcoAlloc_Compile/medium_size-12                 87168             13190 ns/op               729.0 payload_bytes/op        200 B/op          3 allocs/op
ParcoAlloc_Compile/large_size
ParcoAlloc_Compile/large_size-12                   8985            136822 ns/op              8110 payload_bytes/op          200 B/op          3 allocs/op
ParcoDiscard_Compile
ParcoDiscard_Compile/small_size
ParcoDiscard_Compile/small_size-12               661446              1656 ns/op                78.00 payload_bytes/op       200 B/op          3 allocs/op
ParcoDiscard_Compile/medium_size
ParcoDiscard_Compile/medium_size-12               93610             14262 ns/op               729.0 payload_bytes/op        200 B/op          3 allocs/op
ParcoDiscard_Compile/large_size
ParcoDiscard_Compile/large_size-12                 9838            115526 ns/op              8110 payload_bytes/op          200 B/op          3 allocs/op
Json_Compile
Json_Compile/small_size
Json_Compile/small_size-12                       378963              2948 ns/op               200.0 payload_bytes/op       1234 B/op         26 allocs/op
Json_Compile/medium_size
Json_Compile/medium_size-12                       42246             30289 ns/op              1610 payload_bytes/op        10241 B/op        206 allocs/op
Json_Compile/large_size
Json_Compile/large_size-12                         3187            343736 ns/op             16528 payload_bytes/op       101227 B/op       2006 allocs/op
Msgpack_Compile
Msgpack_Compile/small_size
Msgpack_Compile/small_size-12                    487891              2525 ns/op               119.0 payload_bytes/op        490 B/op         24 allocs/op
Msgpack_Compile/medium_size
Msgpack_Compile/medium_size-12                    55425             19528 ns/op               955.0 payload_bytes/op       4053 B/op        207 allocs/op
Msgpack_Compile/large_size
Msgpack_Compile/large_size-12                      6274            191314 ns/op             10135 payload_bytes/op        37432 B/op       2007 allocs/op
```

## Roadmap

- Support for all primitive types:
  - Float32/64
- Extend interface to include version
- Static code generation
- Replace `encoding/binary` usage by faster implementations (`WriteByte`)
- Custom `Reader` and `Writer` interfaces to implement single byte ops
