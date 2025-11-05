// Timers - cho phep chay 1 hanh dong or nhan tin hieu sau 1 khoang time nhat dinh
package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("timer 1 done")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 done")
	}()

	stopped := timer2.Stop()
	if stopped {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(100 * time.Millisecond)
}
