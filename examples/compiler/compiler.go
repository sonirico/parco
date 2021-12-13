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
}

func main() {
	compiler := parco.NewCompiler[Example]().
		SmallVarchar("greet", func(e Example) string {
			return e.Greet
		}).
		UInt8("life_sense", func(e Example) uint8 {
			return e.LifeSense
		}).
		Array("grades", parco.Array[Example, uint8](
			parco.UInt8Header(),
			parco.UInt8Body(),
			func(e Example) parco.Iterable[uint8] {
				return parco.UInt8Iter(e.Grades)
			},
		))

	ex := Example{
		Greet:     "hey",
		LifeSense: 42,
		Grades:    []uint8{5, 6},
	}

	output := bytes.NewBuffer(nil)
	if err := compiler.Compile(ex, output); err != nil {
		log.Fatal(err)
	}
	log.Println(output.Bytes())
}
