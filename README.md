# Parco

Hobbyist binary compiler and parser built with no reflection, highly
extensible, focused on performance, usability through generics and
with zero dependencies.

There are plenty packages over the internet which work by leveraging the power of
struct tags and reflection. While sometimes that can be convenient for some
scenarios, that approach leaves little room to define and register custom types in
addition to have an appositive effect on performance.

## Usage


```go
type Example struct {
    Greet     string
    LifeSense uint8
    Friends   []string
    Grades    map[string]uint8
}

func (e Example) Equals(other Example) bool {
    return reflect.DeepEqual(e, other)
}

func main() {
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
		Map(
			parco.MapField[Example, string, uint8](
				parco.UInt8Header(),
				parco.SmallVarchar(),
				parco.UInt8(),
				func(s *Example, grades map[string]uint8) {
					s.Grades = grades
				},
				func(s *Example) map[string]uint8 {
					return s.Grades
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
		Grades:    map[string]uint8{"math": 5, "english": 6},
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
}
```

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

## TODO

- Support for all primitive types: boolean, nil...
- Extend interface to include version
- Static code generation
- Replace `encoding/binary` usage by faster implementations (`WriteByte`)
- Custom `Reader` and `Writer` interfaces to implement single byte ops
- Support for nested schema definitions.