package utils

type BufferCursor struct {
	cursor int
	data   []byte
}

func (b *BufferCursor) Read(box []byte) (int, error) {
	to := b.cursor + len(box)
	copy(box, b.data[b.cursor:to])
	b.cursor = to
	return len(box), nil
}

func NewBufferCursor(data []byte, cursor int) BufferCursor {
	return BufferCursor{
		cursor: cursor,
		data:   data,
	}
}
