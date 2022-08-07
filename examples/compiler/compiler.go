package main

import (
	"bytes"
	"log"

	"github.com/sonirico/parco"
)

type Example struct {
	Greet     string
	LifeSense uint8
	Grades    []uint8
	Friends   []string
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
			parco.ArrayFieldGetter[Example, uint8](
				parco.UInt8Header(), // up to 255 items
				parco.UInt8(),       // each item
				func(e *Example) parco.Slice[uint8] {
					return e.Grades
				},
			),
		).
		Array(
			parco.ArrayFieldGetter[Example, string](
				parco.UInt8Header(),  // up to 255 items
				parco.SmallVarchar(), // each item
				func(e *Example) parco.Slice[string] {
					return e.Friends
				},
			),
		)

	ex := Example{
		Greet:     "hey",
		LifeSense: 42,
		Grades:    []uint8{5, 6},
		Friends:   []string{"@boliri", "@danirod", "@enrigles", "@f3r"},
	}

	output := bytes.NewBuffer(nil)
	if err := compiler.Compile(ex, output); err != nil {
		log.Fatal(err)
	}
	log.Println(output.Bytes())
}
