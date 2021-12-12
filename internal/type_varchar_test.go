package internal

import (
	"bytes"
	"encoding/binary"
	"strings"
	"testing"
)

func createString(length int, repeat string) string {
	return strings.Repeat(repeat, length)
}

func TestVarcharType_Compile(t *testing.T) {
	tests := []struct {
		Name         string
		Type         Type
		Payload      interface{}
		Expected     []byte
		ExpectsError bool
	}{
		{
			Name:         "compile with uint8 header should succeed",
			Type:         VarcharType{head: UInt8()},
			Payload:      "HOLA",
			Expected:     []byte{4, 72, 79, 76, 65},
			ExpectsError: false,
		},
		{
			Name:         "compile large string with uint8 header should fail",
			Type:         VarcharType{head: UInt8()},
			Payload:      createString(256, "A"),
			Expected:     nil,
			ExpectsError: true,
		},
		{
			Name:         "compile with uint16 header should succeed",
			Type:         VarcharType{head: UInt16(binary.LittleEndian)},
			Payload:      "HOLA",
			Expected:     []byte{4, 0, 72, 79, 76, 65},
			ExpectsError: false,
		},
		{
			Name:         "compile large string with uint8 header should fail",
			Type:         VarcharType{head: UInt16(binary.LittleEndian)},
			Payload:      createString(66000, "A"),
			Expected:     nil,
			ExpectsError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			output := bytes.NewBuffer(nil)
			err := test.Type.Compile(test.Payload, output)

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
