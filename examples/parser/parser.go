package main

import (
	"bytes"
	"log"

	"github.com/sonirico/parco"
)

type Student struct {
	Name   string
	Age    uint8
	Grades []uint8
}

func newParser(factory parco.Factory[Student]) *parco.ModelParser[Student] {
	return parco.ParserModel[Student](factory).
		SmallVarchar(func(s *Student, name string) {
			s.Name = name
		}).
		UInt8(func(s *Student, age uint8) {
			s.Age = age
		}).
		Array(parco.ArrayFieldSetter(
			parco.UInt8Header(),
			parco.UInt8(),
			func(s *Student, items parco.Slice[uint8]) {
				s.Grades = items
			},
		))
}

func parseBytes() {
	data := []byte{4, 72, 79, 76, 65, 42, 2, 9, 10}

	studentFactory := parco.ObjectFactory[Student]()

	parser := newParser(studentFactory)

	parsed, err := parser.ParseBytes(data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsed.Name)
	log.Println(parsed.Age)
	log.Println(parsed.Grades)
}

func parseStream() {
	data := []byte{4, 72, 79, 76, 65, 42, 2, 9, 10}

	studentFactory := parco.ObjectFactory[Student]()
	parser := newParser(studentFactory)

	student, err := parser.Parse(bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	log.Println(student.Name)
	log.Println(student.Age)
	log.Println(student.Grades)

}

func parseWithPool() {
	data := []byte{4, 72, 79, 76, 65, 42, 2, 9, 10}

	studentFactory := parco.PooledFactory[Student](
		parco.ObjectFactory[Student](),
	)

	parser := newParser(studentFactory)

	student, err := parser.Parse(bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	log.Println(student.Name)
	log.Println(student.Age)
	log.Println(student.Grades)
	// DO some work
	// ....

	// Release model
	studentFactory.Put(student)
}

func main() {
	parseWithPool()
	parseBytes()
	parseStream()
}
