package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	err := ConnectServer(time.Second * 3)
	fmt.Println(err)

	// 尽量避免非幂等端超时操作
	err = OrderHandler(time.Second * 3)
	fmt.Println(err)

	var quit string
	fmt.Scanln(&quit)
}

// 幂等案例 ：超时后并没有终止工作的子协程
func ConnectServer(timeout time.Duration) (err error) {
	var done = make(chan bool)
	// 异步连接
	go func() {
		fmt.Println("连接到远程服务器中...")
		time.Sleep(time.Second * 5)
		done <- true
	}()

	// 阻塞：等待ok或超时信号
	select {
	case <-time.After(timeout):
		err = errors.New("连接超时...")
	case <-done:
		err = nil
	}
	return
}

// 非幂等案例：订单处理，虽然客户端收到订单超时提示，但是服务端还是完成了流程操作
func OrderHandler(timeout time.Duration) (err error) {
	var done = make(chan bool)
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("订单处理完成")
		done <- true
	}()

	select {
	case <-time.After(timeout):
		err = errors.New("订单处理超时...")
	case <-done:
		err = nil
	}
	return
}
