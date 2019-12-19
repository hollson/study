package main

import (
	"fmt"
	"sync"
)

//https://www.jianshu.com/p/b09853ecd39d

// 读写删改遍历
func main() {
	var m sync.Map
	m.Store("a", 10)                    //写
	m.Store("a", 11)                    //改(覆盖)
	m.Store("b", 20)                    //
	fmt.Println(m.LoadOrStore("d", 30)) //写或读
	fmt.Println(m.LoadOrStore("d", 31)) //(不覆盖)

	fmt.Println(m.Load("a")) //读
	m.Delete("b")            //删
	m.Range(func(key, value interface{}) bool { //遍历
		fmt.Println(key, value)
		return true
	})
}
