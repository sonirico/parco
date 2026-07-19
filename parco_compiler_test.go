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

	"github.com/stretchr/testify/require"
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
	for i := range le {
		r[i] = uint16(rand.Intn(math.MaxUint16))
	}
	return r
}

func fillMap(le int) map[string]uint8 {
	r := make(map[string]uint8, le)
	for i := range le {
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
			for range b.N {
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

func TestCompiler_Compile(t *testing.T) {
	type pet struct {
		Specie string
		Age    uint8
	}
	type model struct {
		Name string
		Num  uint16
		Tags []string
		Meta map[string]uint8
		Ok   bool
		Pet  pet
	}

	petBuilder := Builder[pet](ObjectFactory[pet]()).
		SmallVarchar(
			func(p *pet) string { return p.Specie },
			func(p *pet, specie string) { p.Specie = specie },
		).
		UInt8(
			func(p *pet) uint8 { return p.Age },
			func(p *pet, age uint8) { p.Age = age },
		)

	parser, compiler := Builder[model](ObjectFactory[model]()).
		SmallVarchar(
			func(m *model) string { return m.Name },
			func(m *model, name string) { m.Name = name },
		).
		UInt16LE(
			func(m *model) uint16 { return m.Num },
			func(m *model, num uint16) { m.Num = num },
		).
		Slice(SliceField[model, string](
			UInt8Header(),
			SmallVarchar(),
			func(m *model, tags SliceView[string]) { m.Tags = tags },
			func(m *model) SliceView[string] { return m.Tags },
		)).
		Map(MapField[model, string, uint8](
			UInt8Header(),
			SmallVarchar(),
			UInt8(),
			func(m *model, meta map[string]uint8) { m.Meta = meta },
			func(m *model) map[string]uint8 { return m.Meta },
		)).
		Bool(
			func(m *model) bool { return m.Ok },
			func(m *model, ok bool) { m.Ok = ok },
		).
		Struct(StructField[model, pet](
			func(m *model) pet { return m.Pet },
			func(m *model, p pet) { m.Pet = p },
			petBuilder,
		)).
		Parco()

	value := model{
		Name: "wire",
		Num:  4242,
		Tags: []string{"go", "binary"},
		Meta: map[string]uint8{"a": 1, "b": 2},
		Ok:   true,
		Pet:  pet{Specie: "cat", Age: 3},
	}

	t.Run("roundtrip preserves every field", func(t *testing.T) {
		var buf bytes.Buffer
		require.NoError(t, compiler.Compile(value, &buf))

		parsed, err := parser.Parse(&buf)

		require.NoError(t, err)
		require.Equal(t, value, parsed)
	})

	t.Run("short writer surfaces ErrCannotWrite", func(t *testing.T) {
		err := compiler.Compile(value, &shortWriter{maxBytes: 3})

		require.ErrorIs(t, err, ErrCannotWrite)
	})
}
