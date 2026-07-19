package parco

import (
	"errors"
	"io"
	"slices"
	"sync"
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

// compileWriter buffers a whole message during compilation so that fixed-size
// types can write into scratch space directly instead of paying a pooled
// buffer round-trip plus a small Write per field. It is fetched from a pool
// once per Compile call and flushed to the underlying writer at the end.
type compileWriter struct {
	w   io.Writer
	buf []byte
}

// scratch grows the buffer by n bytes and returns that space for the caller
// to fill in place.
func (cw *compileWriter) scratch(n int) []byte {
	le := len(cw.buf)
	cw.buf = slices.Grow(cw.buf, n)[:le+n]
	return cw.buf[le:]
}

func (cw *compileWriter) Write(p []byte) (int, error) {
	cw.buf = append(cw.buf, p...)
	return len(p), nil
}

func (cw *compileWriter) flush() error {
	if len(cw.buf) == 0 {
		return nil
	}
	written, err := cw.w.Write(cw.buf)
	if err != nil {
		return err
	}
	if written != len(cw.buf) {
		return ErrCannotWrite
	}
	cw.buf = cw.buf[:0]
	return nil
}

var compileWriterPool = sync.Pool{New: func() any { return &compileWriter{} }}

func getCompileWriter(w io.Writer) *compileWriter {
	//nolint:errcheck // Type assertion is safe - we control pool contents
	cw := compileWriterPool.Get().(*compileWriter)
	cw.w = w
	return cw
}

func putCompileWriter(cw *compileWriter) {
	cw.w = nil
	cw.buf = cw.buf[:0]
	compileWriterPool.Put(cw)
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
