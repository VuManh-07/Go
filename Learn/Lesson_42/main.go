// atomic counter
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var au atomic.Uint64

	var wg sync.WaitGroup

	for range 50 {
		wg.Go(func() {
			for range 1000 {
				au.Add(1)
			}
		})
	}

	wg.Wait()
	fmt.Println("og: ", au.Load())
}
