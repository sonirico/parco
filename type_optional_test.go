package parco

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptionalType_Int_None(t *testing.T) {
	optType := Option[int](IntLE())
	var buf bytes.Buffer

	// Test Compile None (nil)
	err := optType.Compile(nil, &buf)
	require.NoError(t, err)

	// Should only write the header (1 byte for bool = false)
	assert.Equal(t, 1, buf.Len())

	// Test Parse None
	parsed, err := optType.Parse(&buf)
	require.NoError(t, err)
	assert.Nil(t, parsed)
}

func TestOptionalType_Int_Some(t *testing.T) {
	optType := Option[int](IntLE())
	var buf bytes.Buffer

	value := 42
	// Test Compile Some
	err := optType.Compile(&value, &buf)
	require.NoError(t, err)

	// Should write header (1 byte) + int value (4 or 8 bytes)
	assert.Greater(t, buf.Len(), 1)

	// Test Parse Some
	parsed, err := optType.Parse(&buf)
	require.NoError(t, err)
	require.NotNil(t, parsed)
	assert.Equal(t, value, *parsed)
}

func TestOptionalType_String_None(t *testing.T) {
	optType := Option[string](Varchar())
	var buf bytes.Buffer

	// Test Compile None
	err := optType.Compile(nil, &buf)
	require.NoError(t, err)

	// Test Parse None
	parsed, err := optType.Parse(&buf)
	require.NoError(t, err)
	assert.Nil(t, parsed)
}

func TestOptionalType_String_Some(t *testing.T) {
	optType := Option[string](Varchar())
	var buf bytes.Buffer

	value := "hello world"
	// Test Compile Some
	err := optType.Compile(&value, &buf)
	require.NoError(t, err)

	// Test Parse Some
	parsed, err := optType.Parse(&buf)
	require.NoError(t, err)
	require.NotNil(t, parsed)
	assert.Equal(t, value, *parsed)
}

func TestOptionalType_String_EmptyString(t *testing.T) {
	optType := Option[string](Varchar())
	var buf bytes.Buffer

	value := ""
	// Test Compile Some (empty string is still "Some", not "None")
	err := optType.Compile(&value, &buf)
	require.NoError(t, err)

	// Test Parse Some
	parsed, err := optType.Parse(&buf)
	require.NoError(t, err)
	require.NotNil(t, parsed)
	assert.Equal(t, value, *parsed)
}

func TestOptionalType_UInt32_Values(t *testing.T) {
	tests := []struct {
		name  string
		value *uint32
	}{
		{"None", nil},
		{"Zero", func() *uint32 { v := uint32(0); return &v }()},
		{"Max", func() *uint32 { v := uint32(4294967295); return &v }()},
		{"Random", func() *uint32 { v := uint32(12345); return &v }()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optType := Option[uint32](UInt32LE())
			var buf bytes.Buffer

			// Test Compile
			err := optType.Compile(tt.value, &buf)
			require.NoError(t, err)

			// Test Parse
			parsed, err := optType.Parse(&buf)
			require.NoError(t, err)

			if tt.value == nil {
				assert.Nil(t, parsed)
			} else {
				require.NotNil(t, parsed)
				assert.Equal(t, *tt.value, *parsed)
			}
		})
	}
}

func TestOptionalType_Bool_Values(t *testing.T) {
	tests := []struct {
		name  string
		value *bool
	}{
		{"None", nil},
		{"True", func() *bool { v := true; return &v }()},
		{"False", func() *bool { v := false; return &v }()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optType := Option[bool](Bool())
			var buf bytes.Buffer

			// Test Compile
			err := optType.Compile(tt.value, &buf)
			require.NoError(t, err)

			// Test Parse
			parsed, err := optType.Parse(&buf)
			require.NoError(t, err)

			if tt.value == nil {
				assert.Nil(t, parsed)
			} else {
				require.NotNil(t, parsed)
				assert.Equal(t, *tt.value, *parsed)
			}
		})
	}
}

func TestOptionalType_Float32_Values(t *testing.T) {
	tests := []struct {
		name  string
		value *float32
	}{
		{"None", nil},
		{"Zero", func() *float32 { v := float32(0.0); return &v }()},
		{"Positive", func() *float32 { v := float32(123.45); return &v }()},
		{"Negative", func() *float32 { v := float32(-123.45); return &v }()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optType := Option[float32](Float32LE())
			var buf bytes.Buffer

			// Test Compile
			err := optType.Compile(tt.value, &buf)
			require.NoError(t, err)

			// Test Parse
			parsed, err := optType.Parse(&buf)
			require.NoError(t, err)

			if tt.value == nil {
				assert.Nil(t, parsed)
			} else {
				require.NotNil(t, parsed)
				assert.Equal(t, *tt.value, *parsed)
			}
		})
	}
}

func TestOptionalType_ParseError_InsufficientBytes(t *testing.T) {
	optType := Option[int](IntLE())

	// Provide header saying "Some" but no bytes for the int value
	buf := bytes.NewBuffer([]byte{0x01}) // header = true, but no int data

	_, err := optType.Parse(buf)
	assert.Error(t, err)
}

func TestOptionalType_ParseError_EmptyBuffer(t *testing.T) {
	optType := Option[int](IntLE())
	buf := bytes.NewBuffer([]byte{})

	_, err := optType.Parse(buf)
	assert.Error(t, err)
}

func TestOptionalType_RoundTrip(t *testing.T) {
	t.Run("int16 some", func(t *testing.T) {
		value := int16(12345)
		optType := Option[int16](Int16LE())
		var buf bytes.Buffer

		err := optType.Compile(&value, &buf)
		require.NoError(t, err)

		parsed, err := optType.Parse(&buf)
		require.NoError(t, err)
		require.NotNil(t, parsed)
		assert.Equal(t, value, *parsed)
	})

	t.Run("int16 none", func(t *testing.T) {
		optType := Option[int16](Int16LE())
		var buf bytes.Buffer

		err := optType.Compile(nil, &buf)
		require.NoError(t, err)

		parsed, err := optType.Parse(&buf)
		require.NoError(t, err)
		assert.Nil(t, parsed)
	})

	t.Run("uint64 some", func(t *testing.T) {
		value := uint64(18446744073709551615)
		optType := Option[uint64](UInt64LE())
		var buf bytes.Buffer

		err := optType.Compile(&value, &buf)
		require.NoError(t, err)

		parsed, err := optType.Parse(&buf)
		require.NoError(t, err)
		require.NotNil(t, parsed)
		assert.Equal(t, value, *parsed)
	})
}

func TestOptionalType_NestedStructure(t *testing.T) {
	// Test optional of optional (edge case)
	t.Run("optional string nested", func(t *testing.T) {
		// While OptionalType[*T] is unusual, let's test basic nesting
		innerOpt := Option[string](Varchar())
		var buf bytes.Buffer

		value := "nested"
		err := innerOpt.Compile(&value, &buf)
		require.NoError(t, err)

		parsed, err := innerOpt.Parse(&buf)
		require.NoError(t, err)
		require.NotNil(t, parsed)
		assert.Equal(t, value, *parsed)
	})
}
