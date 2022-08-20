package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"time"

	"github.com/sonirico/parco"
)

var (
	data = []byte{
		3,
		104,
		101,
		121,
		42,
		4,
		7,
		64,
		98,
		111,
		108,
		105,
		114,
		105,
		8,
		64,
		100,
		97,
		110,
		105,
		114,
		111,
		100,
		9,
		64,
		101,
		110,
		114,
		105,
		103,
		108,
		101,
		115,
		4,
		64,
		102,
		51,
		114,
		2,
		4,
		109,
		97,
		116,
		104,
		5,
		7,
		101,
		110,
		103,
		108,
		105,
		115,
		104,
		6,
		1,
		3,
		99,
		97,
		116,
		3,
		1,
		1,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		1,
		0,
		0,
		1,
		0,
		172,
		252,
		106,
		67,
		63,
		87,
		91,
		145,
		67,
		213,
		50,
		65,
		160,
		80,
		242,
		234,
		205,
		43,
		13,
		23,
	}
)

type (
	Animal struct {
		Age    uint8
		Specie string
	}

	Example struct {
		Greet              string
		LifeSense          uint8
		Friends            []string
		Grades             map[string]uint8
		EvenOrOdd          bool
		Pet                Animal
		Pointer            *int
		Flags              [5]bool
		Balance            float32
		MorePreciseBalance float64
		CreatedAt          time.Time
	}
)

func (e Example) String() string {
	bts, _ := json.MarshalIndent(e, "", "\t")
	return string(bts)
}

func newAnimalParser() *parco.Parser[Animal] {
	return parco.ParserModel[Animal](parco.ObjectFactory[Animal]()).
		SmallVarchar(func(a *Animal, specie string) { a.Specie = specie }).
		UInt8(func(a *Animal, age uint8) { a.Age = age })
}

func newExampleParser(factory parco.Factory[Example]) *parco.Parser[Example] {
	return parco.ParserModel[Example](factory).
		SmallVarchar(func(s *Example, greet string) {
			s.Greet = greet
		}).
		UInt8(func(s *Example, lifeSense uint8) {
			s.LifeSense = lifeSense
		}).
		Slice(
			parco.SliceFieldSetter(
				parco.UInt8Header(),
				parco.SmallVarchar(),
				func(s *Example, friends parco.SliceView[string]) {
					s.Friends = friends
				},
			),
		).
		Map(
			parco.MapFieldSetter[Example, string, uint8](
				parco.UInt8Header(),
				parco.SmallVarchar(),
				parco.UInt8(),
				func(s *Example, grades map[string]uint8) {
					s.Grades = grades
				},
			),
		).
		Bool(
			func(e *Example, evenOrOdd bool) {
				e.EvenOrOdd = evenOrOdd
			},
		).
		Struct(
			parco.StructFieldSetter[Example, Animal](
				func(e *Example, a Animal) {
					e.Pet = a
				},
				newAnimalParser(),
			),
		).
		Option(
			parco.OptionFieldSetter[Example, int](
				parco.Int(binary.LittleEndian),
				func(e *Example, value *int) { e.Pointer = value },
			),
		).
		Array(
			parco.ArrayFieldSetter[Example, bool](
				5,
				parco.Bool(),
				func(e *Example, flags parco.SliceView[bool]) {
					copy(e.Flags[:], flags)
				},
			),
		).
		Float32(
			binary.LittleEndian,
			func(e *Example, balance float32) {
				e.Balance = balance
			},
		).
		Float64(
			binary.LittleEndian,
			func(e *Example, balance float64) {
				e.MorePreciseBalance = balance
			},
		).
		TimeUTC(
			func(e *Example, createdAt time.Time) {
				e.CreatedAt = createdAt
			},
		)
}

func parseBytes(data []byte) {
	exampleFactory := parco.ObjectFactory[Example]()

	parser := newExampleParser(exampleFactory)

	parsed, err := parser.ParseBytes(data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsed.String())
}

func parseStream(data []byte) {
	exampleFactory := parco.ObjectFactory[Example]()
	parser := newExampleParser(exampleFactory)

	parsed, err := parser.Parse(bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	log.Println(parsed.String())
}

func parseWithPool(data []byte) {
	exampleFactory := parco.PooledFactory[Example](
		parco.ObjectFactory[Example](),
		parco.WithResetFunc(func(e *Example) {
			e.EvenOrOdd = false
			e.Greet = ""
			e.Friends = nil
			e.Grades = make(map[string]uint8)
			e.Pet.Age = 0
			e.Pet.Specie = ""
			e.Pointer = nil
		}),
	)

	parser := newExampleParser(exampleFactory)

	parsed, err := parser.Parse(bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	// DO some work
	log.Println(parsed.String())
	// ....

	// Release model
	exampleFactory.Put(parsed)
}

func main() {
	parseWithPool(data)
	parseBytes(data)
	parseStream(data)
}
