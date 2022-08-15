package main

import (
	"bytes"
	"encoding/binary"
	"log"

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
		Pointer   *int
	}
)

func main() {
	animalCompiler := parco.CompilerModel[Animal]().
		SmallVarchar(func(e *Animal) string {
			return e.Specie
		}).
		UInt8(func(e *Animal) uint8 {
			return e.Age
		})

	exampleCompiler := parco.CompilerModel[Example]().
		SmallVarchar(func(e *Example) string {
			return e.Greet
		}).
		UInt8(func(e *Example) uint8 {
			return e.LifeSense
		}).
		Slice(
			parco.SliceFieldGetter[Example, string](
				parco.UInt8Header(),  // up to 255 items
				parco.SmallVarchar(), // each item
				func(e *Example) parco.SliceView[string] {
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
		}).
		Struct(
			parco.StructFieldGetter[Example, Animal](
				func(e *Example) Animal {
					return e.Pet
				},
				animalCompiler,
			),
		).
		Option(
			parco.OptionFieldGetter[Example, int](
				parco.Int(binary.LittleEndian),
				func(e *Example) *int { return e.Pointer },
			),
		)

	ex := Example{
		Greet:     "hey",
		LifeSense: 42,
		Grades: map[string]uint8{
			"math":    5,
			"english": 6,
		},
		Friends:   []string{"@boliri", "@danirod", "@enrigles", "@f3r"},
		EvenOrOdd: true,
		Pet: Animal{
			Age:    3,
			Specie: "cat",
		},
		Pointer: parco.Ptr(1),
	}

	output := bytes.NewBuffer(nil)
	if err := exampleCompiler.Compile(ex, output); err != nil {
		log.Fatal(err)
	}
	log.Println(parco.FormatBytes(output.Bytes()))
}
