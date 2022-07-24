package parco

import "sync"

type (
	Pool struct {
		pool1   sync.Pool
		pool2   sync.Pool
		pool4   sync.Pool
		pool8   sync.Pool
		pool256 sync.Pool
		poolAny sync.Pool
	}

	Pooler interface {
		Get(size int) *[]byte
		Put(*[]byte)
	}
)

func newSyncPool(byteSize int) sync.Pool {
	return sync.Pool{New: func() any {
		b := make([]byte, byteSize)
		return &b
	}}
}

func NewPool() *Pool {
	return &Pool{
		pool1:   newSyncPool(1),
		pool2:   newSyncPool(2),
		pool4:   newSyncPool(4),
		pool8:   newSyncPool(8),
		pool256: newSyncPool(256),
		poolAny: newSyncPool(32),
	}
}

func (p *Pool) Get(size int) *[]byte {
	switch size {
	case 1:
		return p.Get1()
	case 2:
		return p.Get2()
	case 4:
		return p.Get4()
	case 8:
		return p.Get8()
	case 256:
		return p.Get256()
	default:
		return p.GetAny()
	}
}

func (p *Pool) Put(b *[]byte) {
	switch len(*b) {
	case 1:
		p.Put1(b)
	case 2:
		p.Put2(b)
	case 4:
		p.Put4(b)
	case 8:
		p.Put8(b)
	case 256:
		p.Put256(b)
	default:
		p.PutAny(b)
	}
}

func (p *Pool) Get1() *[]byte {
	return p.pool1.Get().(*[]byte)
}

func (p *Pool) Put1(b *[]byte) {
	p.pool1.Put(b)
}

func (p *Pool) Get2() *[]byte {
	return p.pool2.Get().(*[]byte)
}

func (p *Pool) Put2(b *[]byte) {
	p.pool2.Put(b)
}

func (p *Pool) Get4() *[]byte {
	return p.pool4.Get().(*[]byte)
}

func (p *Pool) Put4(b *[]byte) {
	p.pool4.Put(b)
}

func (p *Pool) Get8() *[]byte {
	return p.pool8.Get().(*[]byte)
}

func (p *Pool) Put8(b *[]byte) {
	p.pool4.Put(b)
}

func (p *Pool) Get256() *[]byte {
	return p.pool256.Get().(*[]byte)
}

func (p *Pool) Put256(b *[]byte) {
	p.pool256.Put(b)
}

func (p *Pool) GetAny() *[]byte {
	return p.poolAny.Get().(*[]byte)
}

func (p *Pool) PutAny(b *[]byte) {
	p.poolAny.Put(b)
}

var (
	SinglePool = NewPool()
)
