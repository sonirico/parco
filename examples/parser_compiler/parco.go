package main

import (
	"bytes"
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
	parser, compiler := parco.NewBuilder().
		FieldGet("greet", types.SmallVarchar(), getGreet).
		FieldGet("life_sense", types.UInt8(), getLifeSense).
		FieldGet("grades", types.Array(2, types.UInt8(), types.UInt8()), getGrades).
		ParCo()

	ex := Example{
		Greet:     "hey",
		LifeSense: 42,
		Grades:    []uint8{5, 6},
	}

	output := bytes.NewBuffer(nil)
	if err := compiler.Compile(ex, output); err != nil {
		log.Fatal(err)
	}

	raw := output.Bytes()
	log.Println("raw bytes", raw)

	result, err := parser.ParseBytes(raw)
	if err != nil {
		log.Fatal(err)
	}

	greet, _ := result.GetString("greet")
	lifeSense, _ := result.GetString("life_sense")
	grades, _ := result.GetArray("grades")

	log.Println("greet", greet)
	log.Println("life sense", lifeSense)
	log.Println("total grades", grades.Len())

	grades.Range(func(value types.Value) {
		log.Println("grade", value.GetUInt8())
	})
}
