package main

import (
	"context"
	"fmt"
)

func main() {
	process := func(ctx context.Context) {
		trace_id, ok := ctx.Value("trace_id").(int)
		if ok {
			fmt.Println("trace_id =", trace_id)
		}

		session, _ := ctx.Value("session").(string)
		fmt.Println("session =", session)
	}

	// 上下文传值
	ctx := context.WithValue(context.Background(), "trace_id", 1001)
	ctx = context.WithValue(ctx, "session", "86011WQGJ")
	process(ctx)
}
