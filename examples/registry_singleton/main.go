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

	Flat struct {
		Price   float32
		Address string
	}
)

const (
	AnimalType int = 0
	FlatType       = 1
)

func (a Animal) ParcoID() int {
	return AnimalType
}

func (a Flat) ParcoID() int {
	return FlatType
}

func main() {
	animalBuilder := parco.Builder[Animal](parco.ObjectFactory[Animal]()).
		SmallVarchar(
			func(a *Animal) string { return a.Specie },
			func(a *Animal, specie string) { a.Specie = specie },
		).
		UInt8(
			func(a *Animal) uint8 { return a.Age },
			func(a *Animal, age uint8) { a.Age = age },
		)

	flatBuilder := parco.Builder[Flat](parco.ObjectFactory[Flat]()).
		Float32(
			binary.LittleEndian,
			func(f *Flat) float32 { return f.Price },
			func(f *Flat, price float32) { f.Price = price },
		).
		SmallVarchar(
			func(f *Flat) string { return f.Address },
			func(f *Flat, address string) { f.Address = address },
		)

	_ = parco.RegisterModel(AnimalType, animalBuilder)
	_ = parco.RegisterModel(FlatType, flatBuilder)

	buf := bytes.NewBuffer(nil)

	_ = parco.CompileModel(Animal{Age: 10, Specie: "monkeys"}, buf)
	_ = parco.CompileModel(Flat{Price: 42, Address: "Plaza mayor"}, buf)
	_ = parco.CompileAnyModel(AnimalType, Animal{Age: 7, Specie: "felix catus"}, buf)

	id, something, _ := parco.ParseModel(buf)
	Print(id, something)
	id, something, _ = parco.ParseModel(buf)
	Print(id, something)
	id, something, _ = parco.ParseModel(buf)
	Print(id, something)
}

func Print(id int, x any) {
	switch id {
	case AnimalType:
		animal := x.(Animal)
		log.Println("animal:", animal)
	case FlatType:
		flat := x.(Flat)
		log.Println("flat", flat)
	}
}
