package parco

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
