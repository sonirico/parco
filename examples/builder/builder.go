package main

import (
	"bytes"
	"log"
	"reflect"

	"github.com/sonirico/parco"
)

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
		Array(
			parco.ArrayField[Example, string](
				parco.UInt8Header(),  // up to 255 items
				parco.SmallVarchar(), // each item's type
				func(e *Example, friends parco.Slice[string]) { e.Friends = friends },
				func(e *Example) parco.Slice[string] { return e.Friends },
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
		ParCo()

	ex := Example{
		Greet:     "hey",
		LifeSense: 42,
		Grades:    map[string]uint8{"math": 5, "english": 6},
		Friends:   []string{"@boliri", "@danirod", "@enrigles", "@f3r"},
		EvenOrOdd: true,
		Pet:       Animal{Age: 3, Specie: "cat"},
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
