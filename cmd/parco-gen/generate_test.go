package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	t.Run("example package matches its committed output", func(t *testing.T) {
		want, err := os.ReadFile("../../examples/codegen/parco_gen.go")
		require.NoError(t, err)

		got, err := generate("../../examples/codegen", nil)

		require.NoError(t, err)
		// git may check the golden file out with CRLF on Windows.
		require.Equal(t, strings.ReplaceAll(string(want), "\r\n", "\n"), string(got))
	})

	t.Run("explicit type list selects untagged structs", func(t *testing.T) {
		got, err := generate("testdata/untagged", []string{"M"})

		require.NoError(t, err)
		require.Contains(t, string(got), "func NewMModel() *parco.Model[M]")
	})

	type errorCase struct {
		name string
		dir  string
		want string
	}
	cases := []errorCase{
		{name: "unknown tag option", dir: "testdata/badtag", want: `M: unknown tag option "huge"`},
		{name: "unsupported field type", dir: "testdata/badtype", want: "M.A: unsupported type chan int"},
		{name: "no tagged struct", dir: "testdata/untagged", want: "no struct with a parco tag"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := generate(tc.dir, nil)

			require.ErrorContains(t, err, tc.want)
		})
	}
}
