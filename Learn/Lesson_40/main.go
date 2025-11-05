// WaitGroups
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	// defer wg.Done()
	fmt.Println("Start ...", id)
	time.Sleep(time.Second)
	fmt.Println("finished ...", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		// wg.Add(1)
		// go worker(i, &wg)
		wg.Go(func() {
			worker(i)
		})
	}

	wg.Wait()
	fmt.Println("Done")
}
