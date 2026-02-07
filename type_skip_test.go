package parco

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkipType(t *testing.T) {
	tests := []struct {
		name string
		pad  int
	}{
		{"skip 1 byte", 1},
		{"skip 2 bytes", 2},
		{"skip 4 bytes", 4},
		{"skip 8 bytes", 8},
		{"skip 16 bytes", 16},
		{"skip 100 bytes", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skipType := SkipType(tt.pad)

			// Create buffer with padding bytes
			paddingData := make([]byte, tt.pad)
			for i := range paddingData {
				paddingData[i] = byte(i % 256)
			}

			// Test Parse - should skip the bytes and return nil
			buf := bytes.NewBuffer(paddingData)
			result, err := skipType.Parse(buf)
			require.NoError(t, err)
			assert.Nil(t, result)
			assert.Equal(t, 0, buf.Len(), "Should have consumed all padding bytes")

			// Test Compile - should write zeros
			var writeBuf bytes.Buffer
			err = skipType.Compile(nil, &writeBuf)
			require.NoError(t, err)
			assert.Equal(t, tt.pad, writeBuf.Len())

			// Verify all bytes are zeros
			written := writeBuf.Bytes()
			for i, b := range written {
				assert.Equal(t, byte(0), b, "Byte at position %d should be zero", i)
			}
		})
	}
}

func TestSkipTypeFactory(t *testing.T) {
	pad := 10
	skipType := SkipTypeFactory(pad)

	var buf bytes.Buffer
	err := skipType.Compile(nil, &buf)
	require.NoError(t, err)
	assert.Equal(t, pad, buf.Len())
}

func TestSkipType_ByteLength(t *testing.T) {
	tests := []int{1, 4, 8, 16, 256}

	for _, pad := range tests {
		t.Run("", func(t *testing.T) {
			skipType := SkipType(pad)
			if bl, ok := skipType.(interface{ ByteLength() int }); ok {
				assert.Equal(t, pad, bl.ByteLength())
			}
		})
	}
}

func TestSkipType_ParseError_InsufficientBytes(t *testing.T) {
	skipType := SkipType(10)

	// Provide only 5 bytes when expecting 10
	buf := bytes.NewBuffer([]byte{0x01, 0x02, 0x03, 0x04, 0x05})
	_, err := skipType.Parse(buf)
	assert.Error(t, err)
}

func TestSkipType_CompileWithValue(t *testing.T) {
	// SkipType should ignore the value passed to Compile
	skipType := SkipType(5)
	var buf bytes.Buffer

	// Pass a non-nil value (should be ignored)
	err := skipType.Compile("ignored value", &buf)
	require.NoError(t, err)
	assert.Equal(t, 5, buf.Len())

	// All bytes should be zeros
	for _, b := range buf.Bytes() {
		assert.Equal(t, byte(0), b)
	}
}

func TestSkipType_RoundTrip(t *testing.T) {
	pad := 8
	skipType := SkipType(pad)

	// Compile
	var writeBuf bytes.Buffer
	err := skipType.Compile(nil, &writeBuf)
	require.NoError(t, err)

	// Parse
	result, err := skipType.Parse(&writeBuf)
	require.NoError(t, err)
	assert.Nil(t, result)
}

func TestSkipType_ZeroPadding(t *testing.T) {
	skipType := SkipType(0)

	var buf bytes.Buffer
	err := skipType.Compile(nil, &buf)
	require.NoError(t, err)
	assert.Equal(t, 0, buf.Len())

	result, err := skipType.Parse(&buf)
	require.NoError(t, err)
	assert.Nil(t, result)
}
