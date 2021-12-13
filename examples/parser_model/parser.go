package main

import (
	"github.com/sonirico/parco"
	"log"
)

type Student struct {
	Name   string
	Age    uint8
	Grades []uint8
}

func parseBytes() {
	data := []byte{4, 72, 79, 76, 65, 42, 2, 9, 10}
	data = []byte{4, 72, 79, 76, 65}

	studentFactory := parco.ObjectFactory[Student]()

	parser := parco.NewModelParser[Student](studentFactory).
		SmallVarchar(func(s *Student, name string) {
			s.Name = name
		})
	//UInt8("life_sense").
	//Array("grades", parco.AnyArray(
	//	parco.UInt8Header(),
	//	parco.AnyUInt8Body(),
	//	nil,
	//))

	parsed, err := parser.ParseBytes(data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsed.Name)
}

func parseStream() {
	data := []byte{4, 72, 79, 76, 65}

	studentFactory := parco.PooledFactory[Student](
		parco.ObjectFactory[Student](),
	)

	parser := parco.NewModelParser[Student](studentFactory).
		SmallVarchar(func(s *Student, name string) {
			s.Name = name
		})
	//UInt8("life_sense").
	//Array("grades", parco.AnyArray(
	//	parco.UInt8Header(),
	//	parco.AnyUInt8Body(),
	//	nil,
	//))
	student, err := parser.ParseBytes(data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(student.Name)
	// DO some work
	// ....

	// Release model
	studentFactory.Put(student)
}

func main() {
	parseStream()
	parseBytes()
}
