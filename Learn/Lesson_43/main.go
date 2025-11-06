// mutexe
package main

import (
	"fmt"
	"sync"
)

type container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *container) inc(name string) {
	c.mu.Lock()

	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	var wg sync.WaitGroup
	co := container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	doInc := func(name string, n int) {
		for range n {
			co.inc(name)
		}
	}

	wg.Go(func() {
		doInc("a", 1000)
	})

	wg.Go(func() {
		doInc("a", 1000)
	})

	wg.Go(func() {
		doInc("b", 1000)
	})

	wg.Wait()
	fmt.Println("done", co.counters)
}
