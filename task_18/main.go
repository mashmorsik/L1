package main

import (
	"fmt"
	"sync"
)

/* Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика. */

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) GetCount() int {
	return c.count
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup

	n := 15
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter.GetCount())
}
