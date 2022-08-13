package parco

import (
	"bytes"
	"testing"
)

func createSlice[T any](value T, length int) []T {
	r := make([]T, length)

	for i := 0; i < length; i++ {
		r = append(r, value)
	}

	return r
}

func TestArrayType_Compile(t *testing.T) {
	tests := []struct {
		Name         string
		inner        ArrayType[uint8]
		Payload      Iterable[uint8]
		Expected     []byte
		ExpectsError bool
	}{
		{
			Name:         "compile array(uint8) should succeed",
			inner:        Array[uint8](UInt8Header(), UInt8()),
			Payload:      Slice[uint8]([]uint8{255, 0}),
			Expected:     []byte{2, 255, 0},
			ExpectsError: false,
		},
		{
			Name:         "compile array(uint8) with payload larger than configured should fail",
			inner:        Array[uint8](UInt8Header(), UInt8()),
			Payload:      Slice[uint8](createSlice[uint8](1, 257)),
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
