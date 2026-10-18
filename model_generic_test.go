//go:build go1.27

package parco

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestModel(arrLen int) *Model[TestStruct] {
	var sliceHeadType IntType
	if arrLen < 256 {
		sliceHeadType = UInt8Header()
	} else {
		sliceHeadType = UInt16HeaderLE()
	}
	return NewModel[TestStruct](ObjectFactory[TestStruct]()).
		Field(Varchar(),
			func(ts *TestStruct) string { return ts.Name },
			func(ts *TestStruct, v string) { ts.Name = v }).
		Field(Varchar(),
			func(ts *TestStruct) string { return ts.Str },
			func(ts *TestStruct, v string) { ts.Str = v }).
		Field(UInt16LE(),
			func(ts *TestStruct) uint16 { return ts.Num },
			func(ts *TestStruct, v uint16) { ts.Num = v }).
		Field(Slice(sliceHeadType, UInt16LE()),
			func(ts *TestStruct) Iterable[uint16] { return SliceView[uint16](ts.Arr) },
			func(ts *TestStruct, v Iterable[uint16]) { ts.Arr = v.Unwrap() }).
		Field(MapType(sliceHeadType, SmallVarchar(), UInt8()),
			func(ts *TestStruct) map[string]uint8 { return ts.Map },
			func(ts *TestStruct, v map[string]uint8) { ts.Map = v }).
		Field(Bool(),
			func(ts *TestStruct) bool { return ts.Bool },
			func(ts *TestStruct, v bool) { ts.Bool = v }).
		Field(Float32(binary.LittleEndian),
			func(ts *TestStruct) float32 { return ts.Float32 },
			func(ts *TestStruct, v float32) { ts.Float32 = v }).
		Field(Float64(binary.LittleEndian),
			func(ts *TestStruct) float64 { return ts.Float64 },
			func(ts *TestStruct, v float64) { ts.Float64 = v })
}

func parcoModelDiscardCompilerFactory(t TestStruct) compileFuncType {
	model := newTestModel(len(t.Arr))
	w := new(discard)
	return func(ts TestStruct) (int, error) {
		defer w.Reset()
		err := model.Compile(t, w)
		return w.Size(), err
	}
}

func BenchmarkParcoModelDiscard_Compile(b *testing.B) {
	benchmarkCompile(b, tests, parcoModelDiscardCompilerFactory)
}

func TestModel(t *testing.T) {
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
		Ptr  *int
	}

	petModel := NewModel[pet](ObjectFactory[pet]()).
		Field(SmallVarchar(),
			func(p *pet) string { return p.Specie },
			func(p *pet, v string) { p.Specie = v }).
		Field(UInt8(),
			func(p *pet) uint8 { return p.Age },
			func(p *pet, v uint8) { p.Age = v })

	m := NewModel[model](ObjectFactory[model]()).
		Field(SmallVarchar(),
			func(m *model) string { return m.Name },
			func(m *model, v string) { m.Name = v }).
		Field(UInt16LE(),
			func(m *model) uint16 { return m.Num },
			func(m *model, v uint16) { m.Num = v }).
		Field(Slice(UInt8Header(), SmallVarchar()),
			func(m *model) Iterable[string] { return SliceView[string](m.Tags) },
			func(m *model, v Iterable[string]) { m.Tags = v.Unwrap() }).
		Field(MapType(UInt8Header(), SmallVarchar(), UInt8()),
			func(m *model) map[string]uint8 { return m.Meta },
			func(m *model, v map[string]uint8) { m.Meta = v }).
		Field(Bool(),
			func(m *model) bool { return m.Ok },
			func(m *model, v bool) { m.Ok = v }).
		Skip(2).
		Field(petModel,
			func(m *model) pet { return m.Pet },
			func(m *model, v pet) { m.Pet = v }).
		Field(Option(Int(binary.LittleEndian)),
			func(m *model) *int { return m.Ptr },
			func(m *model, v *int) { m.Ptr = v })

	seven := 7
	value := model{
		Name: "wire",
		Num:  4242,
		Tags: []string{"go", "binary"},
		Meta: map[string]uint8{"a": 1, "b": 2},
		Ok:   true,
		Pet:  pet{Specie: "cat", Age: 3},
		Ptr:  &seven,
	}

	t.Run("roundtrip preserves every field", func(t *testing.T) {
		var buf bytes.Buffer
		require.NoError(t, m.Compile(value, &buf))

		parsed, err := m.Parse(&buf)

		require.NoError(t, err)
		require.Equal(t, value, parsed)
	})

	t.Run("short writer surfaces ErrCannotWrite", func(t *testing.T) {
		err := m.Compile(value, &shortWriter{maxBytes: 3})

		require.ErrorIs(t, err, ErrCannotWrite)
	})
}
