package main

import (
	"bytes"
	"log"

	"github.com/sonirico/parco"
)

var (
	data = []byte{
		3, 104, 101, 121, 42, 1, 4, 7, 64, 98, 111, 108, 105, 114, 105, 8, 64, 100, 97, 110, 105, 114, 111, 100, 9, 64, 101,
		110, 114, 105, 103, 108, 101, 115, 4, 64, 102, 51, 114, 2, 7, 101, 110, 103, 108, 105, 115, 104, 6, 4, 109, 97,
		116, 104, 5,
	}
)

type Example struct {
	Greet     string
	LifeSense uint8
	Friends   []string
	Grades    map[string]uint8
	EvenOrOdd bool
}

func newParser(factory parco.Factory[Example]) *parco.ModelParser[Example] {
	return parco.ParserModel[Example](factory).
		SmallVarchar(func(s *Example, greet string) {
			s.Greet = greet
		}).
		UInt8(func(s *Example, lifeSense uint8) {
			s.LifeSense = lifeSense
		}).
		Bool(
			func(e *Example, evenOrOdd bool) {
				e.EvenOrOdd = evenOrOdd
			},
		).
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
		)
}

func parseBytes(data []byte) {
	exampleFactory := parco.ObjectFactory[Example]()

	parser := newParser(exampleFactory)

	parsed, err := parser.ParseBytes(data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsed.Greet)
	log.Println(parsed.LifeSense)
	log.Println(parsed.Friends)
	log.Println(parsed.Grades)
	log.Println(parsed.EvenOrOdd)
}

func parseStream(data []byte) {
	exampleFactory := parco.ObjectFactory[Example]()
	parser := newParser(exampleFactory)

	parsed, err := parser.Parse(bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsed.Greet)
	log.Println(parsed.LifeSense)
	log.Println(parsed.Friends)
	log.Println(parsed.Grades)
	log.Println(parsed.EvenOrOdd)
}

func parseWithPool(data []byte) {
	exampleFactory := parco.PooledFactory[Example](
		parco.ObjectFactory[Example](),
	)

	parser := newParser(exampleFactory)

	parsed, err := parser.Parse(bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsed.Greet)
	log.Println(parsed.LifeSense)
	log.Println(parsed.Friends)
	log.Println(parsed.Grades)
	log.Println(parsed.EvenOrOdd)

	// DO some work
	// ....

	// Release model
	exampleFactory.Put(parsed)
}

func main() {
	parseWithPool(data)
	parseBytes(data)
	parseStream(data)
}
