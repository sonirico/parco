package parco

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArrayType_Int(t *testing.T) {
	tests := []struct {
		name   string
		length int
		values []int
	}{
		{"empty array", 0, []int{}},
		{"single element", 1, []int{42}},
		{"small array", 3, []int{1, 2, 3}},
		{"larger array", 5, []int{10, 20, 30, 40, 50}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrayType := Array[int](tt.length, IntLE())
			var buf bytes.Buffer

			// Test Compile
			err := arrayType.Compile(SliceView[int](tt.values), &buf)
			require.NoError(t, err)

			// Test ByteLength
			expectedByteLength := tt.length * IntLE().ByteLength()
			assert.Equal(t, expectedByteLength, arrayType.ByteLength())

			// Test Parse
			parsed, err := arrayType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.values, []int(parsed.Unwrap()))
		})
	}
}

func TestArrayType_String(t *testing.T) {
	tests := []struct {
		name   string
		length int
		values []string
	}{
		{"empty array", 0, []string{}},
		{"small array", 2, []string{"hello", "world"}},
		{"array with empty strings", 3, []string{"", "test", ""}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrayType := Array[string](tt.length, Varchar())
			var buf bytes.Buffer

			// Test Compile
			err := arrayType.Compile(SliceView[string](tt.values), &buf)
			require.NoError(t, err)

			// Test Parse
			parsed, err := arrayType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.values, []string(parsed.Unwrap()))
		})
	}
}

func TestArrayType_UInt16(t *testing.T) {
	tests := []struct {
		name   string
		length int
		values []uint16
	}{
		{"empty array", 0, []uint16{}},
		{"max values", 3, []uint16{65535, 0, 32768}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrayType := Array[uint16](tt.length, UInt16LE())
			var buf bytes.Buffer

			// Test Compile
			err := arrayType.Compile(SliceView[uint16](tt.values), &buf)
			require.NoError(t, err)

			// Test Parse
			parsed, err := arrayType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.values, []uint16(parsed.Unwrap()))
		})
	}
}

func TestArrayType_Bool(t *testing.T) {
	values := []bool{true, false, true, true, false}
	arrayType := Array[bool](5, Bool())
	var buf bytes.Buffer

	// Test Compile
	err := arrayType.Compile(SliceView[bool](values), &buf)
	require.NoError(t, err)

	// Test Parse
	parsed, err := arrayType.Parse(&buf)
	require.NoError(t, err)
	assert.Equal(t, values, []bool(parsed.Unwrap()))
}

func TestArrayType_ParseError_InsufficientBytes(t *testing.T) {
	arrayType := Array[int](2, IntLE())
	// Only provide 4 bytes, but need 8 bytes (2 ints * 4 bytes/int on 32-bit or 8 bytes on 64-bit)
	buf := bytes.NewBuffer([]byte{0x01, 0x00, 0x00, 0x00})

	_, err := arrayType.Parse(buf)
	assert.Error(t, err)
}

func TestArrayType_ParseError_PartialRead(t *testing.T) {
	arrayType := Array[uint16](3, UInt16LE())
	// Only provide 4 bytes, but need 6 bytes (3 uint16 * 2 bytes)
	buf := bytes.NewBuffer([]byte{0x01, 0x00, 0x02, 0x00})

	_, err := arrayType.Parse(buf)
	assert.Error(t, err)
}

func TestArrayType_RoundTrip(t *testing.T) {
	t.Run("int32 array", func(t *testing.T) {
		values := []int32{-1000, 0, 1000, 2147483647, -2147483648}
		arrayType := Array[int32](5, Int32LE())
		var buf bytes.Buffer

		err := arrayType.Compile(SliceView[int32](values), &buf)
		require.NoError(t, err)

		parsed, err := arrayType.Parse(&buf)
		require.NoError(t, err)
		assert.Equal(t, values, []int32(parsed.Unwrap()))
	})

	t.Run("uint8 array", func(t *testing.T) {
		values := []uint8{0, 127, 255, 42, 200}
		arrayType := Array[uint8](5, UInt8())
		var buf bytes.Buffer

		err := arrayType.Compile(SliceView[uint8](values), &buf)
		require.NoError(t, err)

		parsed, err := arrayType.Parse(&buf)
		require.NoError(t, err)
		assert.Equal(t, values, []uint8(parsed.Unwrap()))
	})
}
