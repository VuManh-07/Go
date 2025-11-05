// Goroutines (G–M–P)
// +------------------------------------------------+
// | Go Runtime Scheduler (quản lý toàn cục)        |
// |                                                |
// |  +--------+   +--------+   +--------+          |
// |  |   P1   |   |   P2   |   |   P3   |  ...     |
// |  |        |   |        |   |        |          |
// |  | runq → Gs | runq → Gs | runq → Gs |          |
// |  +--------+   +--------+   +--------+          |
// |     |            |            |                |
// |     v            v            v                |
// |   M1(thread)   M2(thread)   M3(thread)         |
// +------------------------------------------------+
// G: goroutine, M: machine, P: processor
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go func(msg string) {
		fmt.Println(msg)
	}("going1")

	go f("goroutines")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
