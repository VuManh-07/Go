// Worker Pools
package main

import (
	"fmt"
	"time"
)

const (
	numberOfJobs   = 1000
	numberOfWorker = 5
)

func worker(i int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("Word", i, "start job", j)
		time.Sleep(time.Second)
		fmt.Println("Word", i, "finished job", j)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, numberOfWorker)
	result := make(chan int, numberOfJobs)

	for w := 1; w <= numberOfWorker; w++ {
		go worker(w, jobs, result)
	}

	for k := 1; k <= numberOfJobs; k++ {
		jobs <- k
	}
	close(jobs)

	for a := 1; a <= numberOfJobs; a++ {
		<-result
	}
}
