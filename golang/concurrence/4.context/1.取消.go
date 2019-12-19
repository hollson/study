package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 传统方式：读变量
	cancelByVar()
	fmt.Println("==================")

	// 传统方式：通信
	cancelByChannel()
	fmt.Println("==================")

	// 推荐：上下文通道
	cancelByContext()
	fmt.Println("==================")

	// 扩展案例
	withCancelEx()
}

// 方式一：变量控制取消操作
func cancelByVar() {
	var cancel = false
	go func() {
		for {
			if cancel {
				break
			}
			fmt.Println("时间心跳：", time.Now().Second())
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 3)
	cancel = true
}

// 方式二：channel控制取消操作
func cancelByChannel() {
	var cancel = make(chan struct{})
	go func(c chan struct{}) {
	OVER:
		for {
			select {
			case <-c:
				break OVER
			default:
			}

			fmt.Println("时间心跳：", time.Now().Second())
			time.Sleep(time.Second)
		}
	}(cancel)

	time.Sleep(time.Second * 3)
	cancel <- struct{}{}
}

// 方式三：context手动控制取消
func cancelByContext() {
	process := func(ctx context.Context) {
	OVER:
		for {
			fmt.Println("时间心跳：", time.Now().Second())
			time.Sleep(time.Second)
			select {
			case <-ctx.Done(): // 向上报告
				break OVER
			default:
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	go process(ctx)

	time.Sleep(time.Second * 3)
	cancel() //三秒后通知取消
}

// WithCancel案例
func withCancelEx() {
	fetch := func(ctx context.Context) <-chan int {
		trans := make(chan int) //单通道必须经过中转
		go func() {
		OVER:
			for {
				select {
				case <-ctx.Done():
					break OVER
					// (模拟)发送当前时间秒
				case trans <- time.Now().Second():
					time.Sleep(time.Second)
				}
			}
		}()
		return trans
	}

	ctx, cancel := context.WithCancel(context.Background())
	for n := range fetch(ctx) {
		fmt.Println("时间心跳：", n)
		if n%5 == 0 {
			break
		}
	}
	cancel()
}
