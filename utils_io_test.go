package parco

import (
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseReader(t *testing.T) {
	type nextCase struct {
		name    string
		reader  *parseReader
		n       int
		want    []byte
		wantErr error
	}
	cases := []nextCase{
		{
			name:   "bytes hands out a window",
			reader: getParseReaderBytes([]byte{1, 2, 3}),
			n:      2,
			want:   []byte{1, 2},
		},
		{
			name:    "bytes past the end is ErrCannotRead",
			reader:  getParseReaderBytes([]byte{1, 2, 3}),
			n:       4,
			wantErr: ErrCannotRead,
		},
		{
			name:   "reader fills its scratch",
			reader: getParseReader(bytes.NewReader([]byte{1, 2, 3})),
			n:      2,
			want:   []byte{1, 2},
		},
		{
			name:    "reader short stream is ErrCannotRead",
			reader:  getParseReader(bytes.NewReader([]byte{1})),
			n:       2,
			wantErr: ErrCannotRead,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Cleanup(func() { putParseReader(tc.reader) })

			got, err := tc.reader.next(tc.n)

			require.ErrorIs(t, err, tc.wantErr)
			require.Equal(t, tc.want, got)
		})
	}

	t.Run("bytes window advances without copying", func(t *testing.T) {
		data := []byte{1, 2, 3, 4}
		pr := getParseReaderBytes(data)
		t.Cleanup(func() { putParseReader(pr) })

		first, err := pr.next(2)
		require.NoError(t, err)
		second, err := pr.next(2)

		require.NoError(t, err)
		require.Equal(t, []byte{3, 4}, second)
		require.Same(t, &data[0], &first[0])
	})

	t.Run("Read over bytes copies and advances", func(t *testing.T) {
		pr := getParseReaderBytes([]byte{1, 2, 3})
		t.Cleanup(func() { putParseReader(pr) })
		p := make([]byte, 2)

		n, err := pr.Read(p)

		require.NoError(t, err)
		require.Equal(t, 2, n)
		require.Equal(t, []byte{1, 2}, p)
		require.Equal(t, 2, pr.pos)
	})

	t.Run("Read over reader delegates", func(t *testing.T) {
		pr := getParseReader(bytes.NewReader([]byte{1, 2, 3}))
		t.Cleanup(func() { putParseReader(pr) })

		got, err := io.ReadAll(pr)

		require.NoError(t, err)
		require.Equal(t, []byte{1, 2, 3}, got)
	})
}

func TestParser_ParseReader(t *testing.T) {
	type envelope struct {
		When time.Time
		Body TestStruct
	}
	inner := newParser(len(tests[0].Arr))
	parser := ParserModel[envelope](ObjectFactory[envelope]()).
		Struct(StructFieldSetter[envelope, time.Time](
			func(e *envelope, v time.Time) { e.When = v }, TimeUTC())).
		Struct(StructFieldSetter[envelope, TestStruct](
			func(e *envelope, v TestStruct) { e.Body = v }, inner))
	when := time.Date(2026, 10, 18, 12, 0, 0, 0, time.UTC)
	buf := bytes.NewBuffer(nil)
	require.NoError(t, TimeUTC().Compile(when, buf))
	buf.Write(newTestPayload(t, tests[0]))

	t.Run("nested parsers share one reader over io.Reader", func(t *testing.T) {
		got, err := parser.Parse(bytes.NewReader(buf.Bytes()))

		require.NoError(t, err)
		require.Equal(t, envelope{When: when, Body: tests[0]}, got)
	})

	t.Run("nested parsers share one reader over bytes", func(t *testing.T) {
		got, err := parser.ParseBytes(buf.Bytes())

		require.NoError(t, err)
		require.Equal(t, envelope{When: when, Body: tests[0]}, got)
	})

	t.Run("does not read past the message on a shared stream", func(t *testing.T) {
		stream := bytes.NewReader(append(bytes.Clone(buf.Bytes()), 0xAA))

		_, err := parser.Parse(stream)

		require.NoError(t, err)
		require.Equal(t, 1, stream.Len())
	})
}
