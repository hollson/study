package main

import (
	"fmt"
	"time"
)

func main() {
	withRev()
	nobufFoo()
	nobufBar()
}

func withRev() {
	var ch = make(chan int)
	go func() {
		ch <- 1
	}()
	ret, ok := <-ch
	//ret= <-ch //接收一个或两个返回参数。
	fmt.Println(ret, ok)
}

// 非缓冲信道: 就是缓冲大小为0的通道
func nobufFoo() {
	ch := make(chan int)
	go func() {
		fmt.Println(<-ch)
	}()

	//main作为生产者
	ch <- 1
	time.Sleep(time.Second)
}

// range-close：以阻塞的方式遍历通道
func nobufBar() {
	ch := make(chan string)
	go func() {
		ch <- "A"
		ch <- "B"
		close(ch)
	}()

	//main作为消费者
	for c := range ch {
		fmt.Println(c)
	}
}
