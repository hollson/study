package main

import (
	"fmt"
	"sync"
	"time"
)

// 场景：大爷大妈们超市领大米(先排号，再叫号领礼品)
func main() {
	cond := sync.NewCond(new(sync.Mutex))

	// 创建30个goroutine，同时进入等待状态
	fmt.Println("=================领号牌================")
	for i := 0; i < 30; i++ {
		go func(x int) {
			cond.L.Lock()
			defer cond.L.Unlock()

			fmt.Printf("%-2d 已领到号\n", x)
			cond.Wait() // 暂时阻塞,等待通知吧...

			fmt.Printf("%-2d 成功领取领一袋大米\n", x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("=================第一位================")
	cond.Signal()

	time.Sleep(time.Second)
	fmt.Println("=================下一位================")
	cond.Signal()

	time.Sleep(time.Second * 3)
	fmt.Println("================剩余全部===============")
	cond.Broadcast()

	var quit string
	fmt.Scanln(&quit)
}
