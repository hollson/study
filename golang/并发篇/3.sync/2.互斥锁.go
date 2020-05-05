package main

import (
	"fmt"
	"sync"
	"time"
)

// Mutex：是最简单的一种锁类型，也比较暴力，当一个goroutine获得了Mutex后，
// 其他goroutine就只能乖乖等到这个goroutine释放该Mutex。

func main() {
	mutex_foo()
	mutex_bar()
}

func mutex_foo() {
	var (
		share int            //公共变量
		wg    sync.WaitGroup //等待组
		m     sync.Mutex     //互斥锁
	)

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			defer m.Unlock()
			share++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(share)
}

func mutex_bar() {
	var (
		share int        //公共变量
		m     sync.Mutex //互斥锁
	)

	go func() {
		time.Sleep(time.Second)
		fmt.Println(share)
	}()

	m.Lock()
	share = 1
	time.Sleep(time.Second * 5)
	m.Unlock()
}
