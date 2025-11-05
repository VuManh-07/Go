// Channel Directions

package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// tang tinh an toan compile-time
// tang kha nang doc code
// giup tach biet trach nghiem: producer, consumer, logger, pipeline... moi phan chi lam dung 1 nghiem vu
// thuong dung trong pipeline parttern: generate() -> filter() -> transform() -> consume()
//   generate() – tạo dữ liệu đầu vào (source).
//   filter() – lọc dữ liệu theo điều kiện.
//   transform() – chuyển đổi dữ liệu (ví dụ nhân đôi, đổi kiểu, xử lý nội dung).
//   consume() – tiêu thụ kết quả (in ra, lưu vào DB, gửi network...).
//   Mỗi mũi tên → là một channel.
//   Mỗi “hàm” là một goroutine stage.
