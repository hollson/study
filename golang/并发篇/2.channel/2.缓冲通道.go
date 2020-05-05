package main

import "fmt"

// 缓冲通道是非阻塞通道
func main() {
	ch := make(chan string, 3)
	ch <- "A"
	ch <- "B"
	// 避免生产过量
	close(ch)

	for val := range ch {
		fmt.Println(val)
	}
}
