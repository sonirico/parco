package main

import (
	bytes "bytes"
	types "github.com/sonirico/parco/internal"
	parco "github.com/sonirico/parco/pkg"
	"log"
)

type Example struct {
	Greet     string
	LifeSense int
	Grades    []uint8
}

func getGreet(x interface{}) interface{} {
	return x.(Example).Greet
}

func getLifeSense(x interface{}) interface{} {
	return x.(Example).LifeSense
}

func getGrades(x interface{}) interface{} {
	return types.UInt8Iter(x.(Example).Grades)
}

func main() {
	compiler := parco.NewBuilder().
		FieldGet("greet", types.SmallVarchar(), getGreet).
		FieldGet("life_sense", types.UInt8(), getLifeSense).
		FieldGet("grades", types.Array(2, types.UInt8(), types.UInt8()), getGrades).
		Compiler()

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
