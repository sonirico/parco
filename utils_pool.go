package parco

import "sync"

const (
	// MaxPoolMapSize limits the number of dynamically created pools to prevent memory leaks
	MaxPoolMapSize = 100
	// DefaultPoolAnySize is the size of buffers returned by GetAny() (8KB)
	DefaultPoolAnySize = 1 << 13
)

type (
	Pool struct {
		pool1   sync.Pool
		pool2   sync.Pool
		pool4   sync.Pool
		pool8   sync.Pool
		pool256 sync.Pool

		poolAny    sync.Pool
		poolAnyMap map[int]*sync.Pool
	}

	Pooler interface {
		Get(size int) *[]byte
		Put(*[]byte)
	}
)

func newSyncPool(byteSize int) *sync.Pool {
	return &sync.Pool{New: func() any {
		b := make([]byte, byteSize)
		return &b
	}}
}

func NewPool() *Pool {
	return &Pool{
		pool1:      *newSyncPool(1),
		pool2:      *newSyncPool(2),
		pool4:      *newSyncPool(4),
		pool8:      *newSyncPool(8),
		pool256:    *newSyncPool(256),
		poolAny:    *newSyncPool(DefaultPoolAnySize),
		poolAnyMap: make(map[int]*sync.Pool),
	}
}

// Get returns a byte slice of at least the requested size from the pool.
// The returned slice may be larger than requested.
// Always use a defer to Put() the slice back when done.
func (p *Pool) Get(size int) *[]byte {
	switch {
	case size <= 1:
		return p.Get1()
	case size <= 2:
		return p.Get2()
	case size <= 4:
		return p.Get4()
	case size <= 8:
		return p.Get8()
	case size <= 256:
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
	//nolint:errcheck // Type assertion is safe - we control pool contents
	return p.pool1.Get().(*[]byte)
}

func (p *Pool) Put1(b *[]byte) {
	p.pool1.Put(b)
}

func (p *Pool) Get2() *[]byte {
	//nolint:errcheck // Type assertion is safe - we control pool contents
	return p.pool2.Get().(*[]byte)
}

func (p *Pool) Put2(b *[]byte) {
	p.pool2.Put(b)
}

func (p *Pool) Get4() *[]byte {
	//nolint:errcheck // Type assertion is safe - we control pool contents
	return p.pool4.Get().(*[]byte)
}

func (p *Pool) Put4(b *[]byte) {
	p.pool4.Put(b)
}

func (p *Pool) Get8() *[]byte {
	//nolint:errcheck // Type assertion is safe - we control pool contents
	return p.pool8.Get().(*[]byte)
}

func (p *Pool) Put8(b *[]byte) {
	p.pool8.Put(b)
}

func (p *Pool) Get256() *[]byte {
	//nolint:errcheck // Type assertion is safe - we control pool contents
	return p.pool256.Get().(*[]byte)
}

func (p *Pool) Put256(b *[]byte) {
	p.pool256.Put(b)
}

func (p *Pool) GetAny() *[]byte {
	//nolint:errcheck // Type assertion is safe - we control pool contents
	return p.poolAny.Get().(*[]byte)
}

func (p *Pool) PutAny(b *[]byte) {
	p.poolAny.Put(b)
}

// GetAnyMap returns a byte slice of exactly the requested size from a dynamically created pool.
// This method creates new pools on-demand but limits the total number to MaxPoolMapSize.
// If the limit is reached, it falls back to GetAny().
// Note: Only use this if you need exact sizes and will call PutAnyMap() to return the buffer.
func (p *Pool) GetAnyMap(size int) *[]byte {
	pool, ok := p.poolAnyMap[size]
	if !ok {
		// Prevent unbounded growth of poolAnyMap
		if len(p.poolAnyMap) >= MaxPoolMapSize {
			// Fallback to GetAny when we've reached the limit
			return p.GetAny()
		}
		pool = newSyncPool(size)
		p.poolAnyMap[size] = pool
	}
	//nolint:errcheck // Type assertion is safe - we control pool contents
	return pool.Get().(*[]byte)
}

// PutAnyMap returns a byte slice to the dynamically created pool.
// If no pool exists for this size (e.g., buffer came from GetAny fallback),
// the buffer is silently discarded to prevent panics.
func (p *Pool) PutAnyMap(b *[]byte) {
	if b == nil {
		return
	}
	pool, ok := p.poolAnyMap[len(*b)]
	if !ok {
		// Buffer doesn't belong to any pool in the map, likely from GetAny() fallback
		// Just discard it - GC will clean it up
		return
	}
	pool.Put(b)
}

var (
	SinglePool = NewPool()
)
