package parco

import (
	"bytes"
	"strings"
	"testing"
)

func createString(length int, repeat string) string {
	return strings.Repeat(repeat, length)
}

func TestVarcharType_Compile(t *testing.T) {
	tests := []struct {
		Name         string
		Type         Type[string]
		Payload      string
		Expected     []byte
		ExpectsError bool
	}{
		{
			Name:         "compile with uint8 header should succeed",
			Type:         SmallVarchar(),
			Payload:      "HOLA",
			Expected:     []byte{4, 72, 79, 76, 65},
			ExpectsError: false,
		},
		{
			Name:         "compile large string with uint8 header should fail",
			Type:         SmallVarchar(),
			Payload:      createString(256, "A"),
			Expected:     nil,
			ExpectsError: true,
		},
		{
			Name:         "compile with uint16 header should succeed",
			Type:         Varchar(),
			Payload:      "HOLA",
			Expected:     []byte{4, 0, 72, 79, 76, 65},
			ExpectsError: false,
		},
		{
			Name:         "compile large string with uint8 header should fail",
			Type:         SmallVarchar(),
			Payload:      createString(66000, "A"),
			Expected:     nil,
			ExpectsError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := bytes.NewBuffer(nil)
			err := test.Type.Compile(test.Payload, actual)

			hasError := err != nil
			if hasError != test.ExpectsError {
				t.Fatalf("unexpected compile error %s", err)
			}

			if test.ExpectsError {
				return
			}

			if bytes.Compare(actual.Bytes(), test.Expected) != 0 {
				t.Errorf("unexpected compile result, \nwant: '%v' \nhave: '%v'",
					test.Expected, actual)
			}
		})
	}
}
