//go:build go1.27

package main

import (
	"bytes"
	"log"
	"time"
)

func main() {
	model := NewExampleModel()
	seven := 7
	value := Example{
		Greet:   "hi",
		Num:     4242,
		Friends: []string{"ana", "bob"},
		Grades:  map[string]uint8{"ana": 9},
		Flags:   [3]bool{true, false, true},
		Pet:     Pet{Specie: "cat", Age: 3},
		Pointer: &seven,
		Raw:     []byte{0xCA, 0xFE},
		When:    time.Date(2026, 10, 18, 12, 0, 0, 0, time.UTC),
	}

	var buf bytes.Buffer
	if err := model.Compile(value, &buf); err != nil {
		log.Fatal(err)
	}
	size := buf.Len()
	parsed, err := model.Parse(&buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d bytes on the wire, parsed %+v", size, parsed)
}
