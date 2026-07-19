package parco

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
)

func createSlice[T any](value T, length int) []T {
	r := make([]T, 0, length)

	for range length {
		r = append(r, value)
	}

	return r
}

func TestSliceType_Compile(t *testing.T) {
	tests := []struct {
		Name         string
		inner        SliceType[uint8]
		Payload      Iterable[uint8]
		Expected     []byte
		ExpectsError bool
	}{
		{
			Name:         "compile Slice(uint8) should succeed",
			inner:        Slice[uint8](UInt8Header(), UInt8()),
			Payload:      SliceView[uint8]([]uint8{255, 0}),
			Expected:     []byte{2, 255, 0},
			ExpectsError: false,
		},
		{
			Name:         "compile Slice(uint8) with payload larger than configured should fail",
			inner:        Slice[uint8](UInt8Header(), UInt8()),
			Payload:      SliceView[uint8](createSlice[uint8](1, 257)),
			Expected:     nil,
			ExpectsError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			output := bytes.NewBuffer(nil)
			err := test.inner.Compile(test.Payload, output)

			hasError := err != nil
			if hasError != test.ExpectsError {
				t.Fatalf("unexpected compile error %s", err)
			}

			if test.ExpectsError {
				return
			}

			actual := output.Bytes()
			if bytes.Compare(actual, test.Expected) != 0 {
				t.Errorf("unexpected compile result, want '%s' but have '%s'",
					string(test.Expected), string(actual))
			}
		})
	}
}

func TestSliceType_ParseLyingHeader(t *testing.T) {
	// The header declares 60k elements but no element data follows: parsing
	// must fail instead of allocating room for the declared count.
	data := make([]byte, 2)
	binary.LittleEndian.PutUint16(data, 60000)

	_, err := Slice[uint16](UInt16HeaderLE(), UInt16LE()).Parse(bytes.NewReader(data))

	require.Error(t, err)
}
