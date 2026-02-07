package parco

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

type compileFuncType func(t TestStruct) (int, error)

type compileFuncFactory func(t TestStruct) compileFuncType

type TestStruct struct {
	Name    string
	Str     string           `json:"str"`
	Num     uint16           `json:"num"`
	Arr     []uint16         `json:"arr"`
	Map     map[string]uint8 `json:"map"`
	Bool    bool             `json:"bool"`
	Float32 float32          `json:"float32"`
	Float64 float64          `json:"float64"`
}

func fillSeq(le int) []uint16 {
	r := make([]uint16, le)
	for i := 0; i < le; i++ {
		r[i] = uint16(rand.Intn(math.MaxUint16))
	}
	return r
}

func fillMap(le int) map[string]uint8 {
	r := make(map[string]uint8, le)
	for i := 0; i < le; i++ {
		r[strconv.FormatInt(int64(i), 10)] = uint8(rand.Intn(math.MaxUint8))
	}
	return r
}

func newCompiler(arrLen int) *Compiler[TestStruct] {
	var SliceHeadType IntType
	if arrLen < 256 {
		SliceHeadType = UInt8Header()
	} else {
		SliceHeadType = UInt16HeaderLE()
	}
	return CompilerModel[TestStruct]().
		Varchar(func(ts *TestStruct) string { return ts.Name }).
		Varchar(func(ts *TestStruct) string { return ts.Str }).
		UInt16LE(func(ts *TestStruct) uint16 { return ts.Num }).
		Slice(SliceField[TestStruct, uint16](
			SliceHeadType,
			UInt16LE(),
			nil,
			func(ts *TestStruct) SliceView[uint16] {
				return ts.Arr
			},
		)).
		Map(MapField[TestStruct, string, uint8](
			SliceHeadType,
			SmallVarchar(),
			UInt8(),
			nil,
			func(ts *TestStruct) map[string]uint8 {
				return ts.Map
			},
		)).
		Bool(func(ts *TestStruct) bool { return ts.Bool }).
		Float32(binary.LittleEndian, func(ts *TestStruct) float32 { return ts.Float32 }).
		Float64(binary.LittleEndian, func(ts *TestStruct) float64 { return ts.Float64 })
}

var tests = []TestStruct{
	{
		Name:    "small size",
		Str:     "oh hi Mark",
		Num:     42,
		Arr:     fillSeq(10),
		Map:     fillMap(10),
		Bool:    true,
		Float32: math.MaxFloat32,
		Float64: math.MaxFloat64,
	},
	{
		Name:    "medium size",
		Str:     strings.Repeat("oh hi Mark! ", 10),
		Num:     42134,
		Arr:     fillSeq(100),
		Map:     fillMap(100),
		Bool:    true,
		Float32: math.MaxFloat32,
		Float64: math.MaxFloat64,
	},
	{
		Name:    "large size",
		Str:     strings.Repeat("oh hi Mark! ", 100),
		Num:     math.MaxUint16,
		Arr:     fillSeq(1000),
		Map:     fillMap(1000),
		Bool:    true,
		Float32: math.MaxFloat32,
		Float64: math.MaxFloat64,
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
		// creating Compiler needs different field types as per different test payloads
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
