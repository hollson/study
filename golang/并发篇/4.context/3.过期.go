package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	dead := time.Now().Add(3 * time.Second) //截止时间
	ctx, cancel := context.WithDeadline(context.Background(), dead)
	defer cancel()

	// 会员有效期检测
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("欢迎光临,尽情畅玩吧...")
	case <-ctx.Done():
		fmt.Println("会员失效，请充值...", ctx.Err())
	}
	fmt.Println("over")
}
