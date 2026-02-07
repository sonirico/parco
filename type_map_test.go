package parco

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapType_IntString(t *testing.T) {
	tests := []struct {
		name  string
		value map[int]string
	}{
		{"empty map", map[int]string{}},
		{"single entry", map[int]string{1: "one"}},
		{"small map", map[int]string{1: "one", 2: "two", 3: "three"}},
		{"larger map", map[int]string{10: "ten", 20: "twenty", 30: "thirty", 40: "forty"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapType := MapType[int, string](IntLEHeader(), IntLE(), Varchar())
			var buf bytes.Buffer

			// Test Compile
			err := mapType.Compile(tt.value, &buf)
			require.NoError(t, err)

			// Test Parse
			parsed, err := mapType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestMapType_StringInt(t *testing.T) {
	tests := []struct {
		name  string
		value map[string]int
	}{
		{"empty map", map[string]int{}},
		{"small map", map[string]int{"one": 1, "two": 2}},
		{"larger map", map[string]int{"ten": 10, "twenty": 20, "thirty": 30}},
		{"keys with spaces", map[string]int{"hello world": 1, "foo bar": 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapType := MapType[string, int](IntLEHeader(), Varchar(), IntLE())
			var buf bytes.Buffer

			// Test Compile
			err := mapType.Compile(tt.value, &buf)
			require.NoError(t, err)

			// Test Parse
			parsed, err := mapType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestMapType_StringString(t *testing.T) {
	tests := []struct {
		name  string
		value map[string]string
	}{
		{"empty map", map[string]string{}},
		{"single entry", map[string]string{"key": "value"}},
		{"multiple entries", map[string]string{"name": "John", "age": "30", "city": "NYC"}},
		{"empty values", map[string]string{"key1": "", "key2": "value"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapType := MapType[string, string](UInt8Header(), SmallVarchar(), SmallVarchar())
			var buf bytes.Buffer

			err := mapType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := mapType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestMapType_UInt16UInt8(t *testing.T) {
	tests := []struct {
		name  string
		value map[uint16]uint8
	}{
		{"empty map", map[uint16]uint8{}},
		{"small map", map[uint16]uint8{1: 10, 2: 20, 3: 30}},
		{"max values", map[uint16]uint8{65535: 255, 0: 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapType := MapType[uint16, uint8](UInt8Header(), UInt16LE(), UInt8())
			var buf bytes.Buffer

			err := mapType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := mapType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestMapType_BoolInt(t *testing.T) {
	tests := []struct {
		name  string
		value map[bool]int
	}{
		{"empty map", map[bool]int{}},
		{"single entry", map[bool]int{true: 1}},
		{"both keys", map[bool]int{true: 1, false: 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapType := MapType[bool, int](UInt8Header(), Bool(), IntLE())
			var buf bytes.Buffer

			err := mapType.Compile(tt.value, &buf)
			require.NoError(t, err)

			parsed, err := mapType.Parse(&buf)
			require.NoError(t, err)
			assert.Equal(t, tt.value, parsed)
		})
	}
}

func TestMapType_ParseError_InsufficientBytesForHeader(t *testing.T) {
	mapType := MapType[int, string](IntLEHeader(), IntLE(), Varchar())
	// Not enough bytes for int header (need at least 4 or 8 bytes depending on platform)
	buf := bytes.NewBuffer([]byte{0x01})

	_, err := mapType.Parse(buf)
	assert.Error(t, err)
}

func TestMapType_ParseError_InsufficientBytesForKey(t *testing.T) {
	mapType := MapType[int, string](IntLEHeader(), IntLE(), Varchar())
	// Header says 1 element, but not enough bytes for the key
	// Assuming IntLE is 8 bytes, provide header but incomplete key
	buf := bytes.NewBuffer([]byte{
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // length = 1
		0x01, // incomplete key (need 8 bytes)
	})

	_, err := mapType.Parse(buf)
	assert.Error(t, err)
}

func TestMapType_ParseError_InsufficientBytesForValue(t *testing.T) {
	mapType := MapType[uint8, string](UInt8Header(), UInt8(), SmallVarchar())
	buf := bytes.NewBuffer([]byte{
		0x01,       // length = 1
		0x05,       // key = 5
		0x03,       // varchar length = 3 (using SmallVarchar with UInt8 header)
		0x61, 0x62, // only 2 bytes of "abc" string
	})

	_, err := mapType.Parse(buf)
	assert.Error(t, err)
}

func TestMapType_RoundTrip(t *testing.T) {
	t.Run("int32 to float32", func(t *testing.T) {
		value := map[int32]float32{
			1:  1.5,
			-1: -1.5,
			0:  0.0,
		}
		mapType := MapType[int32, float32](UInt8Header(), Int32LE(), Float32LE())
		var buf bytes.Buffer

		err := mapType.Compile(value, &buf)
		require.NoError(t, err)

		parsed, err := mapType.Parse(&buf)
		require.NoError(t, err)
		assert.Equal(t, value, parsed)
	})

	t.Run("uint64 to bool", func(t *testing.T) {
		value := map[uint64]bool{
			0:    false,
			1:    true,
			1000: true,
		}
		mapType := MapType[uint64, bool](UInt16HeaderLE(), UInt64LE(), Bool())
		var buf bytes.Buffer

		err := mapType.Compile(value, &buf)
		require.NoError(t, err)

		parsed, err := mapType.Parse(&buf)
		require.NoError(t, err)
		assert.Equal(t, value, parsed)
	})
}

func TestMapType_LargeMaps(t *testing.T) {
	// Test with a larger map to ensure no issues with scaling
	value := make(map[int]int)
	for i := 0; i < 100; i++ {
		value[i] = i * 2
	}

	mapType := MapType[int, int](IntLEHeader(), IntLE(), IntLE())
	var buf bytes.Buffer

	err := mapType.Compile(value, &buf)
	require.NoError(t, err)

	parsed, err := mapType.Parse(&buf)
	require.NoError(t, err)
	assert.Equal(t, value, parsed)
	assert.Equal(t, len(value), len(parsed))
}
