package sync

import (
	"sync"
)

type counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *counter {
	return &counter{}
}

func (c *counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *counter) Value() int {
	return c.value
}
