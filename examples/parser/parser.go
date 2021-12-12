package main

import (
	"bytes"
	types "github.com/sonirico/parco/internal"
	parco "github.com/sonirico/parco/pkg"
	"log"
)

func parseBytes() {
	data := []byte{4, 72, 79, 76, 65, 42, 9, 10}

	parser := parco.NewBuilder().
		Field("greet", types.SmallVarchar()).
		Field("life_sense", types.UInt8()).
		Field("grades", types.Array(2, types.UInt8(), types.UInt8())).
		Parser()

	parsed, err := parser.ParseBytes(data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsed.GetString("greet"))
	log.Println(parsed.GetUInt8("life_sense"))

	grades, _ := parsed.GetArray("grades")

	log.Println(grades.At(0))
	log.Println(grades.At(1))

	v, _ := grades.At(0)
	log.Println(v.GetUInt8())

	grades.Range(func(value types.Value) {
		log.Println(value.GetUInt8())
	})
}

func parseStream() {
	data := bytes.NewBuffer([]byte{4, 72, 79, 76, 65, 42, 9, 10})

	parser := parco.NewBuilder().
		Field("greet", types.SmallVarchar()).
		Field("life_sense", types.UInt8()).
		Field("grades", types.Array(2, types.UInt8(), types.UInt8())).
		Parser()

	parsed, err := parser.Parse(data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsed.GetString("greet"))
	log.Println(parsed.GetUInt8("life_sense"))

	grades, _ := parsed.GetArray("grades")

	log.Println(grades.At(0))
	log.Println(grades.At(1))

	v, _ := grades.At(0)
	log.Println(v.GetUInt8())

	grades.Range(func(value types.Value) {
		log.Println(value.GetUInt8())
	})
}

func main() {
	parseStream()
	parseBytes()
}
