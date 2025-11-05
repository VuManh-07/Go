// Ticker - sau x giay lai gui y time vao channel cua ticker

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Done")
				return
			case t := <-ticker.C:
				fmt.Println("time at: ", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Stopped")
}
