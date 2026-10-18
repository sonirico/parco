// parco-gen reads the structs of a package and writes a parco.Model for
// every struct that carries at least one `parco` struct tag.
//
// Usage:
//
//	parco-gen [-output parco_gen.go] [-type Name,Name] [dir]
//
// The wire type of every field is inferred from its Go type. The tag only
// adjusts it:
//
//	parco:"-"          skip the field
//	parco:"be"         big-endian integers and floats (default: little-endian)
//	parco:"hdr=uint8"  length header for strings, slices, arrays and maps
//	parco:"location"   time.Time with its location (default: UTC)
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	output := flag.String("output", "parco_gen.go", "output file, relative to dir")
	types := flag.String("type", "", "comma-separated struct names (default: every struct with a parco tag)")
	flag.Parse()

	dir := "."
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	}

	var names []string
	if *types != "" {
		names = strings.Split(*types, ",")
	}

	code, err := generate(dir, names)
	if err != nil {
		fmt.Fprintln(os.Stderr, "parco-gen:", err)
		os.Exit(1)
	}

	if err := os.WriteFile(filepath.Join(dir, *output), code, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "parco-gen:", err)
		os.Exit(1)
	}
}
