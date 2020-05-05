package main

import (
	"fmt"
	"sync"
)

// WaitGroup：等待组(信号量),起到同步等待的作用
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)  //如果注释掉，会出现短路模式，即不等待
		go func(n int) {
			fmt.Printf("第%d条下载完毕.\n", n)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
