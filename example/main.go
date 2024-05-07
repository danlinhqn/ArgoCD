package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Thiết lập HTTP header để chỉ định kiểu nội dung là văn bản đơn giản
	w.Header().Set("Content-Type", "text/plain")

	// Ghi "Hello, World!" vào phản hồi
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Đăng ký một định tuyến đến đường dẫn "/hello" và xác định hàm xử lý của nó là helloWorldHandler
	http.HandleFunc("/hello", helloWorldHandler)

	// Bắt đầu máy chủ HTTP trên cổng 8080
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
