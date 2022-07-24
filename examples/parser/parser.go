package main

import (
	"bytes"
	"log"

	"github.com/sonirico/parco"
)

func parseBytes() {
	data := []byte{4, 72, 79, 76, 65, 42, 2, 9, 10}

	parser := parco.NewParserResult().
		SmallVarchar("greet").
		UInt8("life_sense").
		Array("grades", parco.AnyArray(
			parco.UInt8Header(),
			parco.AnyUInt8Body(),
			nil,
		))

	parsed, err := parser.ParseBytes(data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsed.GetString("greet"))
	log.Println(parsed.GetUInt8("life_sense"))

	grades, err := parsed.GetArray("grades")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(grades.At(0))
	log.Println(grades.At(1))

	v, _ := grades.At(0)
	log.Println(v.GetUInt8())

	grades.Range(func(value parco.Value) {
		log.Println(value.GetUInt8())
	})
}

func parseStream() {
	data := bytes.NewBuffer([]byte{4, 72, 79, 76, 65, 42, 2, 9, 10})

	parser := parco.NewParserResult().
		SmallVarchar("greet").
		UInt8("life_sense").
		Array("grades", parco.AnyArray(
			parco.UInt8Header(),
			parco.AnyUInt8Body(),
			nil,
		))

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

	grades.Range(func(value parco.Value) {
		log.Println(value.GetUInt8())
	})
}

func main() {
	parseStream()
	//parseBytes()
}
