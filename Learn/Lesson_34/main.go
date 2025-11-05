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

/*
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetchWithRetries(url string) {
	// Deadline tổng: 2 phút
	overallCtx, overallCancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer overallCancel()

	// Ticker để khởi một lần gọi mới mỗi 10s (sau lần đầu)
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	client := &http.Client{
		// Không set Timeout ở đây để ta dùng per-request context;
		// nhưng bạn có thể set tổng timeout nếu muốn: Timeout: 10 * time.Second,
	}

	tryOnce := func(ctx context.Context) (bool, error) {
		// Per-request timeout: 10s hoặc ít hơn nếu overallCtx hết hạn trước
		reqCtx, reqCancel := context.WithTimeout(ctx, 10*time.Second)
		defer reqCancel()

		req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, url, nil)
		if err != nil {
			return false, err
		}

		resp, err := client.Do(req)
		if err != nil {
			// lỗi (có thể do timeout hoặc network) -> không thành công, sẽ retry
			return false, err
		}
		defer resp.Body.Close()

		// Nếu server trả mã lỗi (non-2xx) bạn có thể quyết định retry hay fail luôn.
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			// đọc body ngắn để log (không bắt buộc)
			b, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
			fmt.Printf("HTTP %d, body: %s\n", resp.StatusCode, string(b))
			// Xem chính sách: có retry khi 5xx? ở đây ta retry cho mọi lỗi.
			return false, fmt.Errorf("http status %d", resp.StatusCode)
		}

		// Thành công: đọc body (hoặc xử lý tùy cách của bạn)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}

		fmt.Println("API success — response:")
		fmt.Println(string(body))
		return true, nil
	}

	// Thử ngay lần đầu (không cần chờ ticker)
	select {
	case <-overallCtx.Done():
		fmt.Println("API failed: overall timeout before first try")
		return
	default:
		ok, err := tryOnce(overallCtx)
		if ok {
			return
		} else {
			fmt.Printf("Attempt failed (first): %v — will retry every 10s until 2m\n", err)
		}
	}

	// Các lần sau: mỗi tick là một lần thử mới, dừng khi overallCtx hết hạn hoặc khi thành công
	for {
		select {
		case <-overallCtx.Done():
			fmt.Println("API failed: overall timeout (2m) reached")
			return
		case <-ticker.C:
			ok, err := tryOnce(overallCtx)
			if ok {
				return
			} else {
				fmt.Printf("Attempt failed: %v — retrying...\n", err)
			}
		}
	}
}

func main() {
	url := "https://httpbin.org/delay/3" // ví dụ: endpoint chậm
	fetchWithRetries(url)
}
*/
