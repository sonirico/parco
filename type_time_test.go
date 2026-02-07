package parco

import (
	"bytes"
	"encoding/binary"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeUTC(t *testing.T) {
	tests := []struct {
		name  string
		value time.Time
	}{
		{"epoch", time.Unix(0, 0).UTC()},
		{"now", time.Now().UTC()},
		{"future", time.Date(2030, 6, 15, 12, 30, 0, 0, time.UTC)},
		{"past", time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeType := TimeUTC()
			var buf bytes.Buffer

			// Test Compile
			err := timeType.Compile(tt.value, &buf)
			require.NoError(t, err)
			assert.Equal(t, 8, buf.Len()) // time is 8 bytes (int64)

			// Test Parse
			parsed, err := timeType.Parse(&buf)
			require.NoError(t, err)

			// Compare Unix timestamps (nanosecond precision)
			assert.Equal(t, tt.value.Unix(), parsed.Unix())
			assert.Equal(t, tt.value.UnixNano(), parsed.UnixNano())
		})
	}
}

func TestTimeUTC_ByteLength(t *testing.T) {
	timeType := TimeUTC()
	if tl, ok := timeType.(interface{ ByteLength() int }); ok {
		assert.Equal(t, 8, tl.ByteLength())
	}
}

func TestTimeLocation(t *testing.T) {
	tests := []struct {
		name     string
		location string
		skip     bool // Skip if location not available
	}{
		{"UTC", "UTC", false},
		{"America/New_York", "America/New_York", false},
		{"Europe/London", "Europe/London", false},
		{"Asia/Tokyo", "Asia/Tokyo", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc, err := time.LoadLocation(tt.location)
			if err != nil && tt.skip {
				t.Skip("Location not available")
				return
			}
			require.NoError(t, err)

			timeType := TimeLocation()
			var buf bytes.Buffer

			value := time.Date(2024, 6, 15, 12, 30, 45, 123456789, loc)

			// Test Compile
			err = timeType.Compile(value, &buf)
			require.NoError(t, err)

			// Test Parse
			parsed, err := timeType.Parse(&buf)
			require.NoError(t, err)

			// Compare timestamps and locations
			assert.Equal(t, value.Unix(), parsed.Unix())
			assert.Equal(t, value.Location().String(), parsed.Location().String())
		})
	}
}

func TestCompileTime(t *testing.T) {
	now := time.Now().UTC()
	box := make([]byte, 8)

	err := CompileTime(now, box, binary.LittleEndian)
	require.NoError(t, err)

	// Verify bytes are not all zeros (unless it's epoch)
	if now.Unix() != 0 {
		allZero := true
		for _, b := range box {
			if b != 0 {
				allZero = false
				break
			}
		}
		assert.False(t, allZero, "Compiled time should not be all zeros")
	}
}

func TestParseTime(t *testing.T) {
	now := time.Now().UTC()
	box := make([]byte, 8)

	// Compile the time
	err := CompileTime(now, box, binary.LittleEndian)
	require.NoError(t, err)

	// Parse it back
	parsed, err := ParseTime(box, binary.LittleEndian)
	require.NoError(t, err)

	assert.Equal(t, now.Unix(), parsed.Unix())
	assert.Equal(t, now.UnixNano(), parsed.UnixNano())
}

func TestTimeUTC_RoundTrip(t *testing.T) {
	times := []time.Time{
		time.Unix(0, 0).UTC(),
		time.Unix(1234567890, 0).UTC(),
		time.Unix(0, 123456789).UTC(),
		time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.UTC),
	}

	timeType := TimeUTC()
	for _, tt := range times {
		t.Run("", func(t *testing.T) {
			var buf bytes.Buffer

			err := timeType.Compile(tt, &buf)
			require.NoError(t, err)

			parsed, err := timeType.Parse(&buf)
			require.NoError(t, err)

			assert.Equal(t, tt.UnixNano(), parsed.UnixNano())
		})
	}
}

func TestTimeLocation_RoundTrip(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Skip("America/New_York location not available")
	}

	value := time.Date(2024, 6, 15, 12, 30, 0, 0, loc)
	timeType := TimeLocation()
	var buf bytes.Buffer

	err = timeType.Compile(value, &buf)
	require.NoError(t, err)

	parsed, err := timeType.Parse(&buf)
	require.NoError(t, err)

	assert.Equal(t, value.Unix(), parsed.Unix())
	assert.Equal(t, value.Location().String(), parsed.Location().String())
}

func TestTimeUTC_ParseError_InsufficientBytes(t *testing.T) {
	timeType := TimeUTC()
	buf := bytes.NewBuffer([]byte{0x01, 0x02, 0x03}) // Only 3 bytes, need 8

	_, err := timeType.Parse(buf)
	assert.Error(t, err)
}

func TestTimeLocation_ParseError_InsufficientBytes(t *testing.T) {
	timeType := TimeLocation()
	buf := bytes.NewBuffer([]byte{0x01, 0x02}) // Insufficient bytes

	_, err := timeType.Parse(buf)
	assert.Error(t, err)
}
