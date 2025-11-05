// Non-Blocking Channel Operations
package main

import (
	"fmt"
	"time"
)

// func main() {
// 	messages := make(chan string)
// 	signals := make(chan bool)

// 	select {
// 	case msg := <-messages:
// 		fmt.Println("received message ", msg)
// 	default:
// 		fmt.Println("no message received")
// 	}

// 	msg := "hi"
// 	select {
// 	case messages <- msg:
// 		fmt.Println("Sent message: ", msg)
// 	default:
// 		fmt.Println("no message sent")
// 	}

// 	select {
// 	case msg := <-messages:
// 		fmt.Println("received message: ", msg)
// 	case sig := <-signals:
// 		fmt.Println("received signal: ", sig)
// 	default:
// 		fmt.Println("no active")
// 	}
// }

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "res"
	}()

	for {
		select {
		case msg := <-ch:
			fmt.Println("msg: ", msg)
			return // ?? retrun de thoat khoi vong lap
		case <-time.After(500 * time.Millisecond):
			fmt.Println("No res yet, doing something else...")
		}
	}
}
