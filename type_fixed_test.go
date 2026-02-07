package parco

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFixedType_Int(t *testing.T) {
	tests := []struct {
		name          string
		byteLength    int
		value         int
		expectedBytes []byte
	}{
		{
			name:       "4-byte int",
			byteLength: 4,
			value:      12345,
			expectedBytes: func() []byte {
				b := make([]byte, 4)
				binary.LittleEndian.PutUint32(b, uint32(12345))
				return b
			}(),
		},
		{
			name:       "8-byte int",
			byteLength: 8,
			value:      67890,
			expectedBytes: func() []byte {
				b := make([]byte, 8)
				binary.LittleEndian.PutUint64(b, uint64(67890))
				return b
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixedType := NewFixedType[int](
				tt.byteLength,
				func(data []byte) (int, error) {
					if tt.byteLength == 4 {
						return int(binary.LittleEndian.Uint32(data)), nil
					}
					return int(binary.LittleEndian.Uint64(data)), nil
				},
				func(value int, box []byte) error {
					if tt.byteLength == 4 {
						binary.LittleEndian.PutUint32(box, uint32(value))
					} else {
						binary.LittleEndian.PutUint64(box, uint64(value))
					}
					return nil
				},
			)

			// Test Compile
			var compileBuf bytes.Buffer
			err := fixedType.Compile(tt.value, &compileBuf)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedBytes, compileBuf.Bytes())

			// Test Parse
			parseBuf := bytes.NewBuffer(tt.expectedBytes)
			parsedValue, err := fixedType.Parse(parseBuf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsedValue)

			// Test ByteLength - use type assertion to interface with ByteLength method
			if bl, ok := fixedType.(interface{ ByteLength() int }); ok {
				assert.Equal(t, tt.byteLength, bl.ByteLength())
			}
		})
	}
}

func TestFixedType_UInt32(t *testing.T) {
	tests := []struct {
		name  string
		value uint32
	}{
		{"zero", 0},
		{"max uint32", 4294967295},
		{"random value", 123456789},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixedType := NewFixedType[uint32](
				4,
				func(data []byte) (uint32, error) {
					return binary.LittleEndian.Uint32(data), nil
				},
				func(value uint32, box []byte) error {
					binary.LittleEndian.PutUint32(box, value)
					return nil
				},
			)

			var buf bytes.Buffer
			err := fixedType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := fixedType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestFixedType_ParseError_InsufficientBytes(t *testing.T) {
	fixedType := NewFixedType[int](
		4,
		func(data []byte) (int, error) {
			return int(binary.LittleEndian.Uint32(data)), nil
		},
		func(value int, box []byte) error {
			binary.LittleEndian.PutUint32(box, uint32(value))
			return nil
		},
	)

	// Provide a buffer smaller than byteLength
	buf := bytes.NewBuffer([]byte{0x01, 0x02})
	_, err := fixedType.Parse(buf)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrCannotRead)
}

func TestFixedType_ParseError_EmptyBuffer(t *testing.T) {
	fixedType := NewFixedType[uint16](
		2,
		func(data []byte) (uint16, error) {
			return binary.LittleEndian.Uint16(data), nil
		},
		func(value uint16, box []byte) error {
			binary.LittleEndian.PutUint16(box, value)
			return nil
		},
	)

	buf := bytes.NewBuffer([]byte{})
	_, err := fixedType.Parse(buf)
	assert.Error(t, err)
}

func TestFixedType_CompileError_WriteFails(t *testing.T) {
	fixedType := NewFixedType[int](
		4,
		func(data []byte) (int, error) { return 0, nil },
		func(value int, box []byte) error { return nil },
	)

	// Use a writer that always returns an error
	w := &errorWriter{}
	err := fixedType.Compile(123, w)
	assert.Error(t, err)
}

func TestFixedType_CompileError_ShortWrite(t *testing.T) {
	fixedType := NewFixedType[int](
		4,
		func(data []byte) (int, error) { return 0, nil },
		func(value int, box []byte) error { return nil },
	)

	// Use a writer that writes fewer bytes than expected
	w := &shortWriter{maxBytes: 2}
	err := fixedType.Compile(123, w)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrCannotWrite)
}

func TestFixedType_ParseBytes(t *testing.T) {
	fixedType := NewFixedType[uint16](
		2,
		func(data []byte) (uint16, error) {
			return binary.BigEndian.Uint16(data), nil
		},
		func(value uint16, box []byte) error {
			binary.BigEndian.PutUint16(box, value)
			return nil
		},
	)

	data := []byte{0x12, 0x34}
	// Use type assertion to interface with ParseBytes method
	if pb, ok := fixedType.(interface{ ParseBytes([]byte) (uint16, error) }); ok {
		parsed, err := pb.ParseBytes(data)
		require.NoError(t, err)
		assert.Equal(t, uint16(0x1234), parsed)
	}
}

func TestFixedType_CompileBytes(t *testing.T) {
	fixedType := NewFixedType[uint16](
		2,
		func(data []byte) (uint16, error) {
			return binary.BigEndian.Uint16(data), nil
		},
		func(value uint16, box []byte) error {
			binary.BigEndian.PutUint16(box, value)
			return nil
		},
	)

	box := make([]byte, 2)
	// Use type assertion to interface with CompileBytes method
	if cb, ok := fixedType.(interface{ CompileBytes(uint16, []byte) error }); ok {
		err := cb.CompileBytes(0x1234, box)
		require.NoError(t, err)
		assert.Equal(t, []byte{0x12, 0x34}, box)
	}
}

// errorWriter is a mock io.Writer that always returns an error
type errorWriter struct{}

func (w *errorWriter) Write(p []byte) (n int, err error) {
	return 0, io.ErrShortWrite
}

// shortWriter is a mock io.Writer that writes fewer bytes than requested
type shortWriter struct {
	maxBytes int
}

func (w *shortWriter) Write(p []byte) (n int, err error) {
	if len(p) > w.maxBytes {
		return w.maxBytes, nil
	}
	return len(p), nil
}
