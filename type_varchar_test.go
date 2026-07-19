package parco

import (
	"bytes"
	"encoding/binary"
	"io"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/require"
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

func TestVarcharType_Parse(t *testing.T) {
	tests := []struct {
		Name    string
		Type    Type[string]
		Payload string
		Wrap    func(io.Reader) io.Reader
	}{
		{
			Name:    "parse string larger than the pooled buffer",
			Type:    Text(binary.LittleEndian),
			Payload: createString(10_000, "A"),
		},
		{
			Name:    "parse from a reader delivering one byte at a time",
			Type:    SmallVarchar(),
			Payload: "HOLA",
			Wrap:    func(r io.Reader) io.Reader { return iotest.OneByteReader(r) },
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			require.NoError(t, test.Type.Compile(test.Payload, buf))

			var r io.Reader = buf
			if test.Wrap != nil {
				r = test.Wrap(buf)
			}

			actual, err := test.Type.Parse(r)

			require.NoError(t, err)
			require.Equal(t, test.Payload, actual)
		})
	}
}

func TestVarcharType_ParseLyingHeader(t *testing.T) {
	// The header declares 1MB but only 4 bytes follow: parsing must fail
	// with an error instead of panicking or allocating the declared size.
	data := make([]byte, 8)
	binary.LittleEndian.PutUint32(data, 1<<20)
	copy(data[4:], "HOLA")

	_, err := Text(binary.LittleEndian).Parse(bytes.NewReader(data))

	require.ErrorIs(t, err, ErrCannotRead)
}

func TestBlobType_ParseReturnsStableBytes(t *testing.T) {
	blob := Blob(UInt8Header())

	first := parseTestBlob(t, blob, []byte("AAAA"))
	parseTestBlob(t, blob, []byte("BBBB"))

	require.Equal(t, []byte("AAAA"), first)
}

func parseTestBlob(t *testing.T, typ Type[[]byte], payload []byte) []byte {
	t.Helper()

	buf := bytes.NewBuffer(nil)
	require.NoError(t, typ.Compile(payload, buf))

	res, err := typ.Parse(buf)
	require.NoError(t, err)

	return res
}
