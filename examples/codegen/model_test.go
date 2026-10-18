//go:build go1.27

package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGeneratedModel(t *testing.T) {
	t.Run("roundtrip preserves every tagged field", func(t *testing.T) {
		model := NewExampleModel()
		seven := 7
		value := Example{
			Greet:   "hi",
			Num:     4242,
			Friends: []string{"ana", "bob"},
			Grades:  map[string]uint8{"ana": 9},
			Flags:   [3]bool{true, false, true},
			Pet:     Pet{Specie: "cat", Age: 3},
			Pointer: &seven,
			Raw:     []byte{0xCA, 0xFE},
			When:    time.Date(2026, 10, 18, 12, 0, 0, 0, time.UTC),
			Cache:   "not on the wire",
		}
		var buf bytes.Buffer
		require.NoError(t, model.Compile(value, &buf))

		parsed, err := model.Parse(&buf)

		require.NoError(t, err)
		value.Cache = ""
		require.Equal(t, value, parsed)
	})
}
