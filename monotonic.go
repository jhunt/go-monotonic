package monotonic

import (
	"sync"
)

type Clock struct {
	mu sync.Mutex
	v  int64
}

func NewClock(initial int64) *Clock {
	return &Clock{v: initial}
}

func (c *Clock) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}

func (c *Clock) Add(delta int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v += delta
}

func (c *Clock) Increment() {
	c.Add(1)
}

func (c *Clock) Decrement() {
	c.Add(-1)
}
