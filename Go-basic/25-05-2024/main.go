package main

import (
	"context"
	"fmt"
	"time"
)

// 1. Tạo ra một context với timeout = 3 giây
// 2. Tạo một hàm sử dụng context trên, hàm này sẽ trả về kết quả của một logic gì đó + lỗi nếu có
// 3. Bên trong hàm sẽ xử lý như sau:
// 3.1 Trong hàm này mình sẽ lắng nghe context timeout, khi timeout tới thì trả về kết quả luôn
// 3.2 Mình sẽ chạy một go routine tính toán gì đó, nếu như chưa kết thúc timeout mà go routine này tính xong két quả thì trả về func main()
func main() {
	// Tạo một context với timeout là 3 giây
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Gọi hàm sử dụng context với timeout
	result, err := myFunction(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func myFunction(ctx context.Context) (string, error) {
	// Tạo một kênh để nhận kết quả tính toán từ goroutine
	resultChan := make(chan string)

	// Tạo một go routine tính toán
	go func() {
		// Simulate some time-consuming computation
		time.Sleep(2 * time.Second)

		// Trả về kết quả qua kênh
		resultChan <- "Some bla bla result"
	}()

	select {
	case <-ctx.Done():
		// Trong trường hợp context hết hạn
		return "Timeout reached", ctx.Err()
	case result := <-resultChan:
		// Trong trường hợp tính toán hoàn thành trước khi hết hạn
		return result, nil
	}
}
