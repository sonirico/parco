//go:build go1.27

package main

import "time"

//go:generate go run github.com/sonirico/parco/cmd/parco-gen

type Pet struct {
	Specie string `parco:"hdr=uint8"`
	Age    uint8
}

type Example struct {
	Greet   string `parco:"hdr=uint8"`
	Num     uint16 `parco:"be"`
	Friends []string
	Grades  map[string]uint8 `parco:"hdr=uint8"`
	Flags   [3]bool
	Pet     Pet
	Pointer *int
	Raw     []byte
	When    time.Time `parco:"location"`
	Cache   string    `parco:"-"`
}
