package main

import (
	"bytes"
	"log"
	"reflect"

	"github.com/sonirico/parco"
)

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
