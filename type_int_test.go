package parco

import (
	"bytes"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ==================== UInt8 Tests ====================

func TestUInt8(t *testing.T) {
	tests := []struct {
		name  string
		value uint8
	}{
		{"zero", 0},
		{"max", 255},
		{"mid", 127},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uintType := UInt8()
			var buf bytes.Buffer

			err := uintType.Compile(tt.value, &buf)
			require.NoError(t, err)
			assert.Equal(t, 1, buf.Len())

			parsed, err := uintType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestByte(t *testing.T) {
	byteType := Byte()
	var buf bytes.Buffer

	value := byte(123)
	err := byteType.Compile(value, &buf)
	require.NoError(t, err)

	parsed, err := byteType.Parse(&buf)
	require.NoError(t, err)
	assert.Equal(t, value, parsed)
}

func TestUInt8Header(t *testing.T) {
	tests := []struct {
		name    string
		value   int
		wantErr bool
	}{
		{"zero", 0, false},
		{"max valid", 255, false},
		{"mid", 127, false},
		{"overflow", 256, true},
		{"negative", -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headerType := UInt8Header()
			var buf bytes.Buffer

			err := headerType.Compile(tt.value, &buf)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			parsed, err := headerType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

// ==================== Int8 Tests ====================

func TestInt8(t *testing.T) {
	tests := []struct {
		name  string
		value int8
	}{
		{"zero", 0},
		{"max", 127},
		{"min", -128},
		{"positive", 42},
		{"negative", -42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intType := Int8()
			var buf bytes.Buffer

			err := intType.Compile(tt.value, &buf)
			require.NoError(t, err)
			assert.Equal(t, 1, buf.Len())

			parsed, err := intType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestInt8Header(t *testing.T) {
	tests := []struct {
		name    string
		value   int
		wantErr bool
	}{
		{"zero", 0, false},
		{"max valid", 127, false},
		{"min valid", -128, false},
		{"overflow positive", 128, true},
		{"overflow negative", -129, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headerType := Int8Header()
			var buf bytes.Buffer

			err := headerType.Compile(tt.value, &buf)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			parsed, err := headerType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

// ==================== UInt16 Tests ====================

func TestUInt16LE(t *testing.T) {
	tests := []struct {
		name  string
		value uint16
	}{
		{"zero", 0},
		{"max", 65535},
		{"mid", 32768},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uintType := UInt16LE()
			var buf bytes.Buffer

			err := uintType.Compile(tt.value, &buf)
			require.NoError(t, err)
			assert.Equal(t, 2, buf.Len())

			parsed, err := uintType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestUInt16BE(t *testing.T) {
	tests := []struct {
		name  string
		value uint16
	}{
		{"zero", 0},
		{"max", 65535},
		{"specific", 0x1234},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uintType := UInt16BE()
			var buf bytes.Buffer

			err := uintType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := uintType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

// ==================== Int16 Tests ====================

func TestInt16LE(t *testing.T) {
	tests := []struct {
		name  string
		value int16
	}{
		{"zero", 0},
		{"max", 32767},
		{"min", -32768},
		{"positive", 12345},
		{"negative", -12345},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intType := Int16LE()
			var buf bytes.Buffer

			err := intType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := intType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestInt16BE(t *testing.T) {
	tests := []struct {
		name  string
		value int16
	}{
		{"zero", 0},
		{"max", 32767},
		{"min", -32768},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intType := Int16BE()
			var buf bytes.Buffer

			err := intType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := intType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

// ==================== UInt32 Tests ====================

func TestUInt32LE(t *testing.T) {
	tests := []struct {
		name  string
		value uint32
	}{
		{"zero", 0},
		{"max", 4294967295},
		{"mid", 2147483648},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uintType := UInt32LE()
			var buf bytes.Buffer

			err := uintType.Compile(tt.value, &buf)
			require.NoError(t, err)
			assert.Equal(t, 4, buf.Len())

			parsed, err := uintType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestUInt32BE(t *testing.T) {
	tests := []struct {
		name  string
		value uint32
	}{
		{"zero", 0},
		{"max", 4294967295},
		{"specific", 0x12345678},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uintType := UInt32BE()
			var buf bytes.Buffer

			err := uintType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := uintType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

// ==================== Int32 Tests ====================

func TestInt32LE(t *testing.T) {
	tests := []struct {
		name  string
		value int32
	}{
		{"zero", 0},
		{"max", 2147483647},
		{"min", -2147483648},
		{"positive", 1234567},
		{"negative", -1234567},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intType := Int32LE()
			var buf bytes.Buffer

			err := intType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := intType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestInt32BE(t *testing.T) {
	tests := []struct {
		name  string
		value int32
	}{
		{"zero", 0},
		{"max", 2147483647},
		{"min", -2147483648},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intType := Int32BE()
			var buf bytes.Buffer

			err := intType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := intType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

// ==================== UInt64 Tests ====================

func TestUInt64LE(t *testing.T) {
	tests := []struct {
		name  string
		value uint64
	}{
		{"zero", 0},
		{"max", 18446744073709551615},
		{"mid", 9223372036854775808},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uintType := UInt64LE()
			var buf bytes.Buffer

			err := uintType.Compile(tt.value, &buf)
			require.NoError(t, err)
			assert.Equal(t, 8, buf.Len())

			parsed, err := uintType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestUInt64BE(t *testing.T) {
	tests := []struct {
		name  string
		value uint64
	}{
		{"zero", 0},
		{"max", 18446744073709551615},
		{"specific", 0x123456789ABCDEF0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uintType := UInt64BE()
			var buf bytes.Buffer

			err := uintType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := uintType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

// ==================== Int64 Tests ====================

func TestInt64LE(t *testing.T) {
	tests := []struct {
		name  string
		value int64
	}{
		{"zero", 0},
		{"max", 9223372036854775807},
		{"min", -9223372036854775808},
		{"positive", 123456789012345},
		{"negative", -123456789012345},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intType := Int64LE()
			var buf bytes.Buffer

			err := intType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := intType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestInt64BE(t *testing.T) {
	tests := []struct {
		name  string
		value int64
	}{
		{"zero", 0},
		{"max", 9223372036854775807},
		{"min", -9223372036854775808},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intType := Int64BE()
			var buf bytes.Buffer

			err := intType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := intType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

// ==================== Int (platform-dependent) Tests ====================

func TestIntLE(t *testing.T) {
	tests := []struct {
		name  string
		value int
	}{
		{"zero", 0},
		{"max", math.MaxInt},
		{"min", math.MinInt},
		{"positive", 123456},
		{"negative", -123456},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intType := IntLE()
			var buf bytes.Buffer

			err := intType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := intType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestIntBE(t *testing.T) {
	tests := []struct {
		name  string
		value int
	}{
		{"zero", 0},
		{"positive", 123456},
		{"negative", -123456},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intType := IntBE()
			var buf bytes.Buffer

			err := intType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := intType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestUIntLE(t *testing.T) {
	tests := []struct {
		name  string
		value uint
	}{
		{"zero", 0},
		{"max", math.MaxUint},
		{"mid", math.MaxUint / 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uintType := UIntLE()
			var buf bytes.Buffer

			err := uintType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := uintType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

// ==================== Bool Tests ====================

func TestBool(t *testing.T) {
	tests := []struct {
		name  string
		value bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			boolType := Bool()
			var buf bytes.Buffer

			err := boolType.Compile(tt.value, &buf)
			require.NoError(t, err)
			assert.Equal(t, 1, buf.Len())

			parsed, err := boolType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

// ==================== Error Cases ====================

func TestParseInt_InsufficientBytes(t *testing.T) {
	tests := []struct {
		name     string
		typeFunc func() Type[int]
		data     []byte
	}{
		{"int8", func() Type[int] { return Int8Header() }, []byte{}},
		{"int16", func() Type[int] { return Int16LEHeader() }, []byte{0x01}},
		{"int32", func() Type[int] { return Int32LEHeader() }, []byte{0x01, 0x02}},
		{"int64", func() Type[int] { return Int64LEHeader() }, []byte{0x01, 0x02, 0x03}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intType := tt.typeFunc()
			buf := bytes.NewBuffer(tt.data)
			_, err := intType.Parse(buf)
			assert.Error(t, err)
		})
	}
}
