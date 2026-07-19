package parco

import (
	"errors"
	"io"
	"slices"
)

// readFull reads exactly len(data) bytes from r, tolerating short reads.
// A stream ending before len(data) bytes surfaces as ErrCannotRead.
func readFull(r io.Reader, data []byte) error {
	if _, err := io.ReadFull(r, data); err != nil {
		if errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) {
			return ErrCannotRead
		}
		return err
	}
	return nil
}

// readChunked reads exactly size bytes from r, growing the buffer chunk by
// chunk so that a corrupted or malicious size header cannot force a large
// upfront allocation: memory grows only as data actually arrives.
func readChunked(r io.Reader, size, chunk int) ([]byte, error) {
	data := make([]byte, 0, min(size, chunk))
	for len(data) < size {
		n := min(size-len(data), chunk)
		data = slices.Grow(data, n)[:len(data)+n]
		if err := readFull(r, data[len(data)-n:]); err != nil {
			return nil, err
		}
	}
	return data, nil
}

type discard struct {
	counter int
}

func (d *discard) Write(b []byte) (int, error) {
	d.counter += len(b)
	return len(b), nil
}

func (d *discard) Reset() {
	d.counter = 0
}

func (d *discard) Size() int {
	return d.counter
}
