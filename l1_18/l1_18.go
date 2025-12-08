package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	v  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}

func main() {
	var wg sync.WaitGroup
	c := &Counter{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()
			fmt.Println("Run:", taskID)
			c.Inc()
		}(i)
	}

	wg.Wait()
	fmt.Println("Результат:", c.Value())
}
