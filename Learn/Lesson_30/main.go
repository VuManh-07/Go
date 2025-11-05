// Channel Buffering
package main

import (
	"fmt"
	"time"
)

// func main() {
// 	var channel = make(chan int, 5)

// 	fmt.Printf("BufferChan has len = %d, cap = %d\n", len(channel), cap(channel))

// 	channel <- 1
// 	fmt.Printf("BufferChan has len = %d, cap = %d\n", len(channel), cap(channel))

// 	channel <- 2
// 	fmt.Printf("BufferChan has len = %d, cap = %d\n", len(channel), cap(channel))

// 	v_ch1 := <-channel
// 	fmt.Printf("BufferChan has len = %d, cap = %d, v: %v\n", len(channel), cap(channel), v_ch1)

// 	v_ch2 := <-channel
// 	fmt.Printf("BufferChan has len = %d, cap = %d, v: %v\n", len(channel), cap(channel), v_ch2)
// }

const (
	numberOfUrls    = 10
	numberOfWorkers = 2
)

func crawlUrl(queue <-chan int, name string) {
	for v := range queue {
		fmt.Printf("Worker %s is crawl URL %d\n", name, v)
		time.Sleep(time.Second)
	}

	fmt.Printf("Worker %s done and exit\n", name)
}

func startQueue() <-chan int {
	c := make(chan int, 5)

	go func() {
		for i := 1; i <= numberOfUrls; i++ {
			c <- i
			fmt.Printf("URL %d has been enqueued\n", i)
		}
		close(c)
	}()
	return c
}

func main() {
	queue := startQueue()

	for i := 1; i <= numberOfWorkers; i++ {
		go crawlUrl(queue, fmt.Sprintf("%d", i))
	}

	time.Sleep(time.Minute * 1)
}

// ung dung:
// + Web crawler / Scraper (lay du lieu tu trang web)
// + Gui email hang loat / Push notification
// + Xu ly file hoac hinh anh song song
// + pipeline xử lý dữ liệu
// + Nhiệm vụ nền trong microservice (Gửi log sang Kafka / Elastic; Lưu audit logs; Queue giao dịch blockchain; Cập nhật dữ liệu cache nền)
// + trong blockchain nhu analysis or price (so sanh va merge data)
