package exportlease

import (
	"fmt"
	"sync"
)

type Pool struct {
	mu    sync.Mutex
	open  int
	limit int
}

func NewPool(limit int) *Pool {
	return &Pool{limit: limit}
}

func (p *Pool) Acquire() (*Lease, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.open >= p.limit {
		return nil, fmt.Errorf("export resource limit reached: %d", p.limit)
	}
	p.open++
	return &Lease{pool: p}, nil
}

func (p *Pool) Open() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.open
}

type Lease struct {
	pool *Pool
	once sync.Once
}

func (l *Lease) Close() {
	l.once.Do(func() {
		l.pool.mu.Lock()
		l.pool.open--
		l.pool.mu.Unlock()
	})
}
