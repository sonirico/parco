package parco

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
)

func newParser(arrLen int) *Parser[TestStruct] {
	var sliceHeadType IntType
	if arrLen < 256 {
		sliceHeadType = UInt8Header()
	} else {
		sliceHeadType = UInt16HeaderLE()
	}
	return ParserModel[TestStruct](ObjectFactory[TestStruct]()).
		Varchar(func(ts *TestStruct, v string) { ts.Name = v }).
		Varchar(func(ts *TestStruct, v string) { ts.Str = v }).
		UInt16LE(func(ts *TestStruct, v uint16) { ts.Num = v }).
		Slice(SliceField[TestStruct, uint16](
			sliceHeadType,
			UInt16LE(),
			func(ts *TestStruct, v SliceView[uint16]) { ts.Arr = v },
			nil,
		)).
		Map(MapField[TestStruct, string, uint8](
			sliceHeadType,
			SmallVarchar(),
			UInt8(),
			func(ts *TestStruct, v map[string]uint8) { ts.Map = v },
			nil,
		)).
		Bool(func(ts *TestStruct, v bool) { ts.Bool = v }).
		Float32(binary.LittleEndian, func(ts *TestStruct, v float32) { ts.Float32 = v }).
		Float64(binary.LittleEndian, func(ts *TestStruct, v float64) { ts.Float64 = v })
}

func newTestPayload(t testing.TB, ts TestStruct) []byte {
	buf := bytes.NewBuffer(nil)
	require.NoError(t, newCompiler(len(ts.Arr)).Compile(ts, buf))
	return buf.Bytes()
}

type parseFuncType func([]byte) error

type parseFuncFactory func(*testing.B, TestStruct) parseFuncType

func parcoBytesParserFactory(b *testing.B, ts TestStruct) parseFuncType {
	parser := newParser(len(ts.Arr))
	return func(data []byte) error {
		_, err := parser.ParseBytes(data)
		return err
	}
}

func parcoReaderParserFactory(b *testing.B, ts TestStruct) parseFuncType {
	parser := newParser(len(ts.Arr))
	r := bytes.NewReader(nil)
	return func(data []byte) error {
		r.Reset(data)
		_, err := parser.Parse(r)
		return err
	}
}

func benchmarkParse(b *testing.B, tests []TestStruct, factory parseFuncFactory) {
	for _, test := range tests {
		data := newTestPayload(b, test)
		parseFunc := factory(b, test)
		b.Run(test.Name, func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				if err := parseFunc(data); err != nil {
					b.Error(err)
				}
			}
		})
	}
}

func BenchmarkParcoBytes_Parse(b *testing.B) {
	benchmarkParse(b, tests, parcoBytesParserFactory)
}

func BenchmarkParcoReader_Parse(b *testing.B) {
	benchmarkParse(b, tests, parcoReaderParserFactory)
}

func TestParser_Parse(t *testing.T) {
	t.Run("roundtrip preserves every field", func(t *testing.T) {
		want := tests[0]
		data := newTestPayload(t, want)

		got, err := newParser(len(want.Arr)).ParseBytes(data)

		require.NoError(t, err)
		require.Equal(t, want, got)
	})

	t.Run("truncated payload surfaces ErrCannotRead", func(t *testing.T) {
		data := newTestPayload(t, tests[0])

		_, err := newParser(len(tests[0].Arr)).ParseBytes(data[:len(data)-1])

		require.ErrorIs(t, err, ErrCannotRead)
	})
}
