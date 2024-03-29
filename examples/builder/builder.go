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
