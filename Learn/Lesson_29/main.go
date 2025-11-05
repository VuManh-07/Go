// channels
// Channel là kênh giao tiếp trung gian giữa các Goroutines có thể gởi và nhận dữ liệu cho nhau một cách an toàn thông qua cơ chế lock-free.

package main

import "fmt"

func main() {
	channelName := make(chan int)

	go func() {
		for i := range 10 {
			channelName <- i
		}
		close(channelName)
	}()

	for value := range channelName {
		fmt.Printf("Value: %d\n", value)
	}

}
