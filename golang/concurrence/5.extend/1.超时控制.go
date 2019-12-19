package main

import (
	"errors"
	"fmt"
	"time"
)

// 超时只适合做幂等操作
func main() {
	//err := ConnectServer(time.Second * 3)
	//fmt.Println(err)

	err := youya(time.Second * 3)

	fmt.Println(err)

	var quit string
	fmt.Scanln(&quit)
}

// 案例1：超时后并没有终止工作的子协程
func ConnectServer(timeout time.Duration) (err error) {
	var done = make(chan bool)
	// 异步连接
	go func() {
		fmt.Println("连接到远程服务器中...")
		time.Sleep(time.Second * 5)
		fmt.Println("连接成功...")
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

// 订单处理(非幂等)
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

// 订单处理(非幂等)
func youya(timeout time.Duration) (err error) {
	var done = make(chan bool)
	var tip = 0
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("订单处理完成", tip)  //一点也不优雅
		done <- true
	}()

	select {
	case <-time.After(timeout):
		err = errors.New("订单处理超时...")
		tip = 1
	case <-done:
		err = nil
	}
	return
}
