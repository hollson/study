package main

import (
	"fmt"
	"sync"
	"time"
)

// 单例锁在被第一次被执行后就永久锁定了(可不就相当于受精卵么)
// 整个once对象只被Do一次
func main() {
	var once sync.Once
	once.Do(func() { fmt.Println("hello 1") })
	once.Do(func() { fmt.Println("hello 2") })
	go func() {
		once.Do(func() {
			fmt.Println("hello 3")
		})
	}()

	fmt.Println("over")
	time.Sleep(time.Second)
}

//======================================================

// 单例模式
type singleRef map[string]string

var (
	once     sync.Once
	instance singleRef //私有变量
)

func Single() singleRef {
	once.Do(func() {
		instance = make(singleRef)
	})
	return instance
}
