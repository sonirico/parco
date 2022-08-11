package main

import (
	"bytes"
	"log"

	"github.com/sonirico/parco"
)

var (
	data = []byte{
		3, 104, 101, 121, 42, 4, 7, 64, 98, 111, 108, 105, 114, 105, 8, 64, 100, 97, 110, 105, 114, 111, 100,
		9, 64, 101, 110, 114, 105, 103, 108, 101, 115, 4, 64, 102, 51, 114, 2, 4, 109, 97, 116, 104, 5, 7, 101,
		110, 103, 108, 105, 115, 104, 6, 1, 3, 99, 97, 116, 3,
	}
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

func newAnimalParser() *parco.ModelParser[Animal] {
	return parco.ParserModel[Animal](parco.ObjectFactory[Animal]()).
		SmallVarchar(func(a *Animal, specie string) { a.Specie = specie }).
		UInt8(func(a *Animal, age uint8) { a.Age = age })
}

func newExampleParser(factory parco.Factory[Example]) *parco.ModelParser[Example] {
	return parco.ParserModel[Example](factory).
		SmallVarchar(func(s *Example, greet string) {
			s.Greet = greet
		}).
		UInt8(func(s *Example, lifeSense uint8) {
			s.LifeSense = lifeSense
		}).
		Array(
			parco.ArrayFieldSetter(
				parco.UInt8Header(),
				parco.SmallVarchar(),
				func(s *Example, friends parco.Slice[string]) {
					s.Friends = friends
				},
			),
		).
		Map(
			parco.MapFieldSetter[Example, string, uint8](
				parco.UInt8Header(),
				parco.SmallVarchar(),
				parco.UInt8(),
				func(s *Example, grades map[string]uint8) {
					s.Grades = grades
				},
			),
		).
		Bool(
			func(e *Example, evenOrOdd bool) {
				e.EvenOrOdd = evenOrOdd
			},
		).
		Struct(
			parco.StructFieldSetter[Example, Animal](
				func(e *Example, a Animal) {
					e.Pet = a
				},
				newAnimalParser(),
			),
		)
}

func parseBytes(data []byte) {
	exampleFactory := parco.ObjectFactory[Example]()

	parser := newExampleParser(exampleFactory)

	parsed, err := parser.ParseBytes(data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsed.Greet)
	log.Println(parsed.LifeSense)
	log.Println(parsed.Friends)
	log.Println(parsed.Grades)
	log.Println(parsed.EvenOrOdd)
	log.Println(parsed.Pet)
}

func parseStream(data []byte) {
	exampleFactory := parco.ObjectFactory[Example]()
	parser := newExampleParser(exampleFactory)

	parsed, err := parser.Parse(bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsed.Greet)
	log.Println(parsed.LifeSense)
	log.Println(parsed.Friends)
	log.Println(parsed.Grades)
	log.Println(parsed.EvenOrOdd)
	log.Println(parsed.Pet)
}

func parseWithPool(data []byte) {
	exampleFactory := parco.PooledFactory[Example](
		parco.ObjectFactory[Example](),
	)

	parser := newExampleParser(exampleFactory)

	parsed, err := parser.Parse(bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	// DO some work
	log.Println(parsed.Greet)
	log.Println(parsed.LifeSense)
	log.Println(parsed.Friends)
	log.Println(parsed.Grades)
	log.Println(parsed.EvenOrOdd)
	log.Println(parsed.Pet)
	// ....

	// Release model
	exampleFactory.Put(parsed)
}

func main() {
	parseWithPool(data)
	parseBytes(data)
	parseStream(data)
}
