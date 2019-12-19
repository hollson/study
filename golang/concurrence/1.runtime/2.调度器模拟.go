package main

import (
	"fmt"
	"runtime"
	"time"
)

// 模拟时间片
func main() {
	//修改参数值，观察打印结果的离散度
	runtime.GOMAXPROCS(4)
	foo := func(n int) {
		for {
			fmt.Print(n)
		}
	}
	go foo(0)
	go foo(1)
	time.Sleep(time.Second * 5)
	fmt.Println("~~~Done~~~")
}