package Counter

import "sync"

type Counter struct {
	mu    sync.Mutex //mutual exclusion lock.
	count int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

//any goroutine calling the function will aquire a lock, will be the only thread acessing the method,
//other goroutines will have to wait for the current goroutine(who has the lock) to finish it's task and unlock
//before entering the method.

func (c *Counter) Value() int {
	return c.count
}

func NewCounter() *Counter {
	return &Counter{}
}
