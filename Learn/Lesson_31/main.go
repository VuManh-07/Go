// Channel Synchronization (cho cho cac goroutines chay xong)
package main

import (
	"fmt"
)

// func worker(done chan bool) {
// 	fmt.Println("Working ...")
// 	time.Sleep(time.Second)
// 	fmt.Println("done")

// 	done <- true
// }

// func main() {
// 	done := make(chan bool, 1)
// 	go worker(done)

// 	<-done
// }

func worker(id int, done chan struct{}) {
	fmt.Printf("Worker %d done\n", id)

	done <- struct{}{} // gui tin hieu rong thay vi gia tri de tiet kiem bo nho
}

func main() {
	done := make(chan struct{})

	for i := range 3 {
		go worker(i, done)
	}

	for i := 1; i <= 3; i++ {
		<-done
	}

	fmt.Println("All done")
}
