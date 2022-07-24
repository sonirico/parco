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
		Type         ArrayType[any, uint8]
		Payload      Iterable[uint8]
		Expected     []byte
		ExpectsError bool
	}{
		{
			Name:         "compile array(uint8) should succeed",
			Type:         Array[any, uint8](UInt8Header(), UInt8Body(), nil),
			Payload:      UInt8Iter([]uint8{255, 0}),
			Expected:     []byte{2, 255, 0},
			ExpectsError: false,
		},
		{
			Name:         "compile array(uint8) with payload larger than configured should fail",
			Type:         Array[any, uint8](UInt8Header(), UInt8Body(), nil),
			Payload:      UInt8Iter(createSlice[uint8](1, 257)),
			Expected:     nil,
			ExpectsError: true,
		},
		//{
		//	Name:         "compile array(uint16) should succeed",
		//	Type:         Array[any, uint8](UInt8Header(), UInt8Body(), nil),
		//	Payload:      UInt16Iter([]uint16{65535, 512}),
		//	Expected:     []byte{2, 255, 255, 0, 2},
		//	ExpectsError: false,
		//},
		//{
		//	Name:         "compile array(uint16) with payload larger than configured should fail",
		//	Type:         Array( UInt8(), UInt16LE()),
		//	Payload:      UInt16Iter(createSlice(1, 257)),
		//	Expected:     nil,
		//	ExpectsError: true,
		//},
		//{
		//	Name:         "compile array(uint16, uint16) should success",
		//	Type:         Array( UInt16BE(), UInt16LE()),
		//	Payload:      UInt16Iter([]uint16{1, 2}),
		//	Expected:     []byte{0, 2, 1, 0, 2, 0},
		//	ExpectsError: false,
		//},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			output := bytes.NewBuffer(nil)
			err := test.Type.CompileIterable(test.Payload, output)

			hasError := err != nil
			if hasError != test.ExpectsError {
				t.Fatalf("unexpected compile error %s", err)
			}

			if test.ExpectsError {
				return
			}

			//actual := output.Bytes()
			//if bytes.Compare(actual, test.Expected) != 0 {
			//	t.Errorf("unexpected compile result, want '%s' but have '%s'",
			//		string(test.Expected), string(actual))
			//}
		})
	}
}
