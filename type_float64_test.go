package parco

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFloat64LE(t *testing.T) {
	tests := []struct {
		name  string
		value float64
	}{
		{"zero", 0.0},
		{"positive", 123.456789},
		{"negative", -123.456789},
		{"max float64", math.MaxFloat64},
		{"min positive float64", math.SmallestNonzeroFloat64},
		{"positive infinity", math.Inf(1)},
		{"negative infinity", math.Inf(-1)},
		{"NaN", math.NaN()},
		{"negative zero", math.Copysign(0, -1)},
		{"pi", math.Pi},
		{"e", math.E},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			floatType := Float64LE()
			testFloat64RoundTrip(t, floatType, binary.LittleEndian, tt.value)
		})
	}
}

func TestFloat64BE(t *testing.T) {
	tests := []struct {
		name  string
		value float64
	}{
		{"zero", 0.0},
		{"positive", 456.789012},
		{"negative", -456.789012},
		{"max float64", math.MaxFloat64},
		{"NaN", math.NaN()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			floatType := Float64BE()
			testFloat64RoundTrip(t, floatType, binary.BigEndian, tt.value)
		})
	}
}

func testFloat64RoundTrip(t *testing.T, floatType Type[float64], order binary.ByteOrder, value float64) {
	var buf bytes.Buffer

	// Test Compile
	err := floatType.Compile(value, &buf)
	require.NoError(t, err)

	// Verify byte representation
	expectedBytes := make([]byte, 8)
	order.PutUint64(expectedBytes, math.Float64bits(value))

	if math.IsNaN(value) {
		// For NaN, we can't compare bytes directly, just verify it's a NaN
		assert.True(t, math.IsNaN(math.Float64frombits(order.Uint64(buf.Bytes()))))
	} else {
		assert.Equal(t, expectedBytes, buf.Bytes())
	}

	// Test Parse
	parsed, err := floatType.Parse(&buf)
	require.NoError(t, err)

	if math.IsNaN(value) {
		assert.True(t, math.IsNaN(parsed))
	} else {
		assert.Equal(t, value, parsed)
	}
}

func TestCompileFloat64(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		order binary.ByteOrder
	}{
		{"positive LE", 123.456789, binary.LittleEndian},
		{"negative LE", -123.456789, binary.LittleEndian},
		{"positive BE", 123.456789, binary.BigEndian},
		{"negative BE", -123.456789, binary.BigEndian},
		{"very small", 1e-300, binary.LittleEndian},
		{"very large", 1e300, binary.LittleEndian},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box := make([]byte, 8)
			err := CompileFloat64(tt.value, box, tt.order)
			require.NoError(t, err)

			expected := make([]byte, 8)
			tt.order.PutUint64(expected, math.Float64bits(tt.value))
			assert.Equal(t, expected, box)
		})
	}
}

func TestParseFloat64(t *testing.T) {
	t.Run("success LE", func(t *testing.T) {
		data := make([]byte, 8)
		binary.LittleEndian.PutUint64(data, math.Float64bits(123.456789))

		parsed, err := ParseFloat64(data, binary.LittleEndian)
		require.NoError(t, err)
		assert.Equal(t, 123.456789, parsed)
	})

	t.Run("success BE", func(t *testing.T) {
		data := make([]byte, 8)
		binary.BigEndian.PutUint64(data, math.Float64bits(123.456789))

		parsed, err := ParseFloat64(data, binary.BigEndian)
		require.NoError(t, err)
		assert.Equal(t, 123.456789, parsed)
	})

	t.Run("insufficient bytes", func(t *testing.T) {
		data := []byte{0x01, 0x02, 0x03, 0x04}
		_, err := ParseFloat64(data, binary.LittleEndian)
		assert.ErrorIs(t, err, NewErrUnSufficientBytesError(8, 0))
	})

	t.Run("empty buffer", func(t *testing.T) {
		data := []byte{}
		_, err := ParseFloat64(data, binary.LittleEndian)
		assert.Error(t, err)
	})

	t.Run("single byte", func(t *testing.T) {
		data := []byte{0x01}
		_, err := ParseFloat64(data, binary.LittleEndian)
		assert.Error(t, err)
	})
}

func TestFloat64_SpecialValues(t *testing.T) {
	t.Run("positive infinity", func(t *testing.T) {
		floatType := Float64LE()
		var buf bytes.Buffer

		inf := math.Inf(1)
		err := floatType.Compile(inf, &buf)
		require.NoError(t, err)

		parsed, err := floatType.Parse(&buf)
		require.NoError(t, err)
		assert.True(t, math.IsInf(parsed, 1))
	})

	t.Run("negative infinity", func(t *testing.T) {
		floatType := Float64LE()
		var buf bytes.Buffer

		negInf := math.Inf(-1)
		err := floatType.Compile(negInf, &buf)
		require.NoError(t, err)

		parsed, err := floatType.Parse(&buf)
		require.NoError(t, err)
		assert.True(t, math.IsInf(parsed, -1))
	})

	t.Run("NaN", func(t *testing.T) {
		floatType := Float64LE()
		var buf bytes.Buffer

		nan := math.NaN()
		err := floatType.Compile(nan, &buf)
		require.NoError(t, err)

		parsed, err := floatType.Parse(&buf)
		require.NoError(t, err)
		assert.True(t, math.IsNaN(parsed))
	})
}

func TestFloat64_Precision(t *testing.T) {
	// Test that we don't lose precision on round-trip
	tests := []float64{
		1.7976931348623157e+308, // Near max
		2.2250738585072014e-308, // Near min positive
		math.Pi,
		math.E,
		math.Sqrt2,
		1.0 / 3.0,
		0.1 + 0.2, // Classic floating point edge case
	}

	floatType := Float64LE()
	for _, value := range tests {
		t.Run("", func(t *testing.T) {
			var buf bytes.Buffer
			err := floatType.Compile(value, &buf)
			require.NoError(t, err)

			parsed, err := floatType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, value, parsed)
		})
	}
}
