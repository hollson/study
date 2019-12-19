package main

import (
	"fmt"
	"sync"
)

// RWMutex比Mutex相对友好些，是经典的单写多读模型。在读锁占用的情况下，会阻止写，但不阻止读，也就是多个goroutine可同时获取读锁
// (调用RLock方法)；而写锁(调用Lock方法)会阻止任何其他goroutine(无论读和写)进来，整个锁由该goroutine独占。
// 从RWMutex的实现看，RWMutex类型其实组合了Mutex：

var (
	// 逻辑中使用的某个变量
	count int
	// 与变量对应的使用互斥锁
	rwmutex sync.RWMutex
)

func Get() int {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	return count
}

func Set(c int) {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	count = c
}

func main() {
	Set(1)
	fmt.Println(Get())
}
