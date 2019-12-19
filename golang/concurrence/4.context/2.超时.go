package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	withTimeout()
}

func withTimeout() {
	var wg sync.WaitGroup
	wg.Add(1)

	process := func(ctx context.Context) {
	OVER:
		for {
			time.Sleep(time.Second)
			select {
			case er := <-ctx.Done():
				fmt.Println(er) //取消的原因
				break OVER
			default:
				fmt.Println("时间心跳：", time.Now().Second())
			}
		}
		wg.Done()
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	go process(ctx)

	_ = cancel //cancel是个方法，这也可以手动提前结束
	wg.Wait()
}
