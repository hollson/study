package main

import (
	"fmt"
	"time"
)

// select-default是非阻塞模式，应尽量避免

// 监听【收、发、关】
func main() {
	ListenClose()
	Generator()
}

// 监听close
func ListenClose() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 3)
		close(ch)
	}()

	select {
	case <-ch:
		fmt.Println("closed...")
	}
}

// 一个channel就是一个被激活的通道信号发生器
// channel发生器
func Generator() {
	timeBeat := func() chan int {
		out := make(chan int)
		go func() {
			for {
				time.Sleep(time.Second)
				out <- time.Now().Second()
			}
		}()
		return out
	}

	ch1, ch2 := timeBeat(), timeBeat()
	// 循环阻塞
	for {
		select {
		case val := <-ch1:
			fmt.Println("ch1:", val)
		case val := <-ch2:
			fmt.Println("ch2:", val)
		}
	}
}
