package parco

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFloat32LE(t *testing.T) {
	tests := []struct {
		name  string
		value float32
	}{
		{"zero", 0.0},
		{"positive", 123.45},
		{"negative", -123.45},
		{"max float32", math.MaxFloat32},
		{"min positive float32", math.SmallestNonzeroFloat32},
		{"positive infinity", float32(math.Inf(1))},
		{"negative infinity", float32(math.Inf(-1))},
		{"NaN", float32(math.NaN())},
		{"negative zero", float32(math.Copysign(0, -1))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			floatType := Float32LE()
			testFloat32RoundTrip(t, floatType, binary.LittleEndian, tt.value)
		})
	}
}

func TestFloat32BE(t *testing.T) {
	tests := []struct {
		name  string
		value float32
	}{
		{"zero", 0.0},
		{"positive", 456.78},
		{"negative", -456.78},
		{"max float32", math.MaxFloat32},
		{"NaN", float32(math.NaN())},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			floatType := Float32BE()
			testFloat32RoundTrip(t, floatType, binary.BigEndian, tt.value)
		})
	}
}

func testFloat32RoundTrip(t *testing.T, floatType Type[float32], order binary.ByteOrder, value float32) {
	var buf bytes.Buffer

	// Test Compile
	err := floatType.Compile(value, &buf)
	require.NoError(t, err)

	// Verify byte representation
	expectedBytes := make([]byte, 4)
	order.PutUint32(expectedBytes, math.Float32bits(value))

	if math.IsNaN(float64(value)) {
		// For NaN, we can't compare bytes directly, just verify it's a NaN
		assert.True(t, math.IsNaN(float64(math.Float32frombits(order.Uint32(buf.Bytes())))))
	} else {
		assert.Equal(t, expectedBytes, buf.Bytes())
	}

	// Test Parse
	parsed, err := floatType.Parse(&buf)
	require.NoError(t, err)

	if math.IsNaN(float64(value)) {
		assert.True(t, math.IsNaN(float64(parsed)))
	} else {
		assert.Equal(t, value, parsed)
	}
}

func TestCompileFloat32(t *testing.T) {
	tests := []struct {
		name  string
		value float32
		order binary.ByteOrder
	}{
		{"positive LE", 123.45, binary.LittleEndian},
		{"negative LE", -123.45, binary.LittleEndian},
		{"positive BE", 123.45, binary.BigEndian},
		{"negative BE", -123.45, binary.BigEndian},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box := make([]byte, 4)
			err := CompileFloat32(tt.value, box, tt.order)
			require.NoError(t, err)

			expected := make([]byte, 4)
			tt.order.PutUint32(expected, math.Float32bits(tt.value))
			assert.Equal(t, expected, box)
		})
	}
}

func TestParseFloat32(t *testing.T) {
	t.Run("success LE", func(t *testing.T) {
		data := make([]byte, 4)
		binary.LittleEndian.PutUint32(data, math.Float32bits(123.45))

		parsed, err := ParseFloat32(data, binary.LittleEndian)
		require.NoError(t, err)
		assert.Equal(t, float32(123.45), parsed)
	})

	t.Run("success BE", func(t *testing.T) {
		data := make([]byte, 4)
		binary.BigEndian.PutUint32(data, math.Float32bits(123.45))

		parsed, err := ParseFloat32(data, binary.BigEndian)
		require.NoError(t, err)
		assert.Equal(t, float32(123.45), parsed)
	})

	t.Run("insufficient bytes", func(t *testing.T) {
		data := []byte{0x01, 0x02}
		_, err := ParseFloat32(data, binary.LittleEndian)
		assert.ErrorIs(t, err, NewErrUnSufficientBytesError(4, 0))
	})

	t.Run("empty buffer", func(t *testing.T) {
		data := []byte{}
		_, err := ParseFloat32(data, binary.LittleEndian)
		assert.Error(t, err)
	})
}

func TestFloat32_SpecialValues(t *testing.T) {
	t.Run("positive infinity", func(t *testing.T) {
		floatType := Float32LE()
		var buf bytes.Buffer

		inf := float32(math.Inf(1))
		err := floatType.Compile(inf, &buf)
		require.NoError(t, err)

		parsed, err := floatType.Parse(&buf)
		require.NoError(t, err)
		assert.True(t, math.IsInf(float64(parsed), 1))
	})

	t.Run("negative infinity", func(t *testing.T) {
		floatType := Float32LE()
		var buf bytes.Buffer

		negInf := float32(math.Inf(-1))
		err := floatType.Compile(negInf, &buf)
		require.NoError(t, err)

		parsed, err := floatType.Parse(&buf)
		require.NoError(t, err)
		assert.True(t, math.IsInf(float64(parsed), -1))
	})

	t.Run("NaN", func(t *testing.T) {
		floatType := Float32LE()
		var buf bytes.Buffer

		nan := float32(math.NaN())
		err := floatType.Compile(nan, &buf)
		require.NoError(t, err)

		parsed, err := floatType.Parse(&buf)
		require.NoError(t, err)
		assert.True(t, math.IsNaN(float64(parsed)))
	})
}
