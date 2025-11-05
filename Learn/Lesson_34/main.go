// Timeouts

package main

import (
	"fmt"
	"time"
)

func fetchAPI(done chan<- bool) {
	time.Sleep(time.Second * 2)
	done <- true
}

func main() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go fetchAPI(done1)
	go fetchAPI(done2)

	select {
	case <-done1:
		fmt.Println("done")
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout api")
	}

	select {
	case <-done2:
		fmt.Println("done")
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout api")
	}
}
