package parco

import (
	"bytes"
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"math"
	"math/rand"
	"strings"
	"testing"
)

type compileFuncType func(t TestStruct) (int, error)

type compileFuncFactory func(t TestStruct) compileFuncType

type TestStruct struct {
	Name string
	Str  string   `json:"str"`
	Num  uint16   `json:"num"`
	Arr  []uint16 `json:"arr"`
}

func fillSeq(le int) []uint16 {
	r := make([]uint16, le)
	for i := 0; i < le; i++ {
		r[i] = uint16(rand.Intn(math.MaxUint16))
	}
	return r
}

func newCompiler(arrLen int) Compiler[TestStruct] {
	var arrayHeadType IntType
	if arrLen < 256 {
		arrayHeadType = UInt8Header()
	} else {
		arrayHeadType = UInt16LEHeader[TestStruct]()
	}
	return NewCompiler[TestStruct]().
		Varchar("name", func(ts TestStruct) string { return ts.Name }).
		Varchar("str", func(ts TestStruct) string { return ts.Str }).
		UInt16LE("num", func(ts TestStruct) uint16 { return ts.Num }).
		Array("arr", Array[TestStruct, uint16](
			arrayHeadType,
			UInt16BodyLE(),
			func(ts TestStruct) Iterable[uint16] {
				return UInt16Iter(ts.Arr)
			},
		))
}

var tests = []TestStruct{
	{
		Name: "small size",
		Str:  "oh hi Mark",
		Num:  42,
		Arr:  fillSeq(10),
	},
	{
		Name: "medium size",
		Str:  strings.Repeat("oh hi Mark! ", 10),
		Num:  42134,
		Arr:  fillSeq(100),
	},
	{
		Name: "large size",
		Str:  strings.Repeat("oh hi Mark! ", 100),
		Num:  math.MaxUint16,
		Arr:  fillSeq(1000),
	},
}

func jsonCompilerFactory(_ TestStruct) compileFuncType {
	return func(t TestStruct) (int, error) {
		bts, err := json.Marshal(t)
		return len(bts), err
	}
}

func msgPackCompilerFactory(_ TestStruct) compileFuncType {
	return func(t TestStruct) (int, error) {
		bts, err := msgpack.Marshal(t)
		return len(bts), err
	}
}

func parcoCompilerFactory(t TestStruct) compileFuncType {
	compiler := newCompiler(len(t.Arr))
	buf := bytes.NewBuffer(nil)
	return func(ts TestStruct) (int, error) {
		defer buf.Reset()
		err := compiler.Compile(t, buf)
		return buf.Len(), err
	}
}

func parcoDiscardCompilerFactory(t TestStruct) compileFuncType {
	compiler := newCompiler(len(t.Arr))
	w := new(discard)
	return func(ts TestStruct) (int, error) {
		defer w.Reset()
		err := compiler.Compile(t, w)
		return w.Size(), err
	}
}

func benchmarkCompile(b *testing.B, tests []TestStruct, compileFuncFactory compileFuncFactory) {
	for _, test := range tests {
		// creating compiler needs different field types as per different test payloads
		compileFunc := compileFuncFactory(test)
		b.Run(test.Name, func(b *testing.B) {
			var totalBytes int
			for i := 0; i < b.N; i++ {
				n, err := compileFunc(test)
				if err != nil {
					b.Error(err)
				}
				totalBytes += n
			}
			b.ReportMetric(float64(totalBytes/b.N), "payload_bytes/op")
		})
	}
}

func BenchmarkParcoAlloc_Compile(b *testing.B) {
	benchmarkCompile(b, tests, parcoCompilerFactory)
}

func BenchmarkParcoDiscard_Compile(b *testing.B) {
	benchmarkCompile(b, tests, parcoDiscardCompilerFactory)
}

func BenchmarkJson_Compile(b *testing.B) {
	benchmarkCompile(b, tests, jsonCompilerFactory)
}

func BenchmarkMsgpack_Compile(b *testing.B) {
	benchmarkCompile(b, tests, msgPackCompilerFactory)
}
