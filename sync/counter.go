package sync_go

import "sync"

type Counter struct {
	mutex sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}
