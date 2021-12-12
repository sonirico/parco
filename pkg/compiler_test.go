package pkg

import (
	bytes "bytes"
	"encoding/json"
	types "github.com/sonirico/parco/internal"
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
	Num  int      `json:"num"`
	Arr  []uint16 `json:"arr"`
}

func getStr(x interface{}) interface{} {
	return x.(TestStruct).Str
}

func getNum(x interface{}) interface{} {
	return x.(TestStruct).Num
}

func getArr(x interface{}) interface{} {
	return types.UInt16Iter(x.(TestStruct).Arr)
}

func fillSeq(le int) []uint16 {
	r := make([]uint16, le)
	for i := 0; i < le; i++ {
		r[i] = uint16(rand.Intn(math.MaxUint16))
	}
	return r
}

func newCompiler(arrLen int) Compiler {
	var arrayHeadType types.Type
	if arrLen < 256 {
		arrayHeadType = types.UInt8()
	} else {
		arrayHeadType = types.UInt16LE()
	}
	return NewBuilder().
		FieldGet("str", types.Varchar(types.UInt16LE()), getStr).
		FieldGet("num", types.UInt16LE(), getNum).
		FieldGet("arr", types.Array(arrLen, arrayHeadType, types.UInt16LE()), getArr).
		Compiler()
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

func benchmarkCompile(b *testing.B, tests []TestStruct, compileFuncFactory compileFuncFactory) {
	for _, test := range tests {
		compileFunc := compileFuncFactory(test)
		for i := 0; i < b.N; i++ {
			b.Run(test.Name, func(b *testing.B) {
				b.ResetTimer()
				b.StartTimer()
				n, err := compileFunc(test)
				b.StopTimer()
				if err != nil {
					b.Fatal(err)
				}
				b.ReportMetric(float64(n), "bytes/op")
			})
		}
	}

}

func BenchmarkParco_Compile(b *testing.B) {
	benchmarkCompile(b, tests, func(t TestStruct) compileFuncType {
		compiler := newCompiler(len(t.Arr))
		return func(tt TestStruct) (int, error) {
			buf := bytes.NewBuffer(nil)
			err := compiler.Compile(t, buf)
			return buf.Len(), err
		}
	})
}

func BenchmarkJson_Compile(b *testing.B) {
	benchmarkCompile(b, tests, func(t TestStruct) compileFuncType {
		return func(t TestStruct) (int, error) {
			bts, err := json.Marshal(t)
			return len(bts), err
		}
	})
}
