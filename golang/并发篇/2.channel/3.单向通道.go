package main

import (
	"fmt"
)

// 生产者
func producer(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send:", i)
		c <- i
	}
	close(c)
}

// 消费者
func consumer(c <-chan int) {
	for v := range c {
		fmt.Println("receive:", v)
	}
}

// 单项通道用于约束参数，双向通道是通用通道
func main() {
	ch := make(chan int)    //1.阻塞式
	ch = make(chan int, 10) //2.非阻塞式
	go producer(ch)
	go consumer(ch)

	var quit string
	fmt.Scanln(&quit)
}
