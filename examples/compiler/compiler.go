package main

import (
	"bytes"
	"log"

	"github.com/sonirico/parco"
)

type Example struct {
	Greet     string
	LifeSense uint8
	Friends   []string
	Grades    map[string]uint8
	EvenOrOdd bool
}

func main() {
	compiler := parco.CompilerModel[Example]().
		SmallVarchar(func(e *Example) string {
			return e.Greet
		}).
		UInt8(func(e *Example) uint8 {
			return e.LifeSense
		}).
		Array(
			parco.ArrayFieldGetter[Example, string](
				parco.UInt8Header(),  // up to 255 items
				parco.SmallVarchar(), // each item
				func(e *Example) parco.Slice[string] {
					return e.Friends
				},
			),
		).
		Map(
			parco.MapFieldGetter[Example, string, uint8](
				parco.UInt8Header(),  // up to 255 items
				parco.SmallVarchar(), // key type
				parco.UInt8(),        // value type
				func(e *Example) map[string]uint8 {
					return e.Grades
				},
			),
		).
		Bool(func(e *Example) bool {
			return e.EvenOrOdd
		})

	ex := Example{
		Greet:     "hey",
		LifeSense: 42,
		Grades: map[string]uint8{
			"math":    5,
			"english": 6,
		},
		Friends:   []string{"@boliri", "@danirod", "@enrigles", "@f3r"},
		EvenOrOdd: true,
	}

	output := bytes.NewBuffer(nil)
	if err := compiler.Compile(ex, output); err != nil {
		log.Fatal(err)
	}
	log.Println(output.Bytes())
}
