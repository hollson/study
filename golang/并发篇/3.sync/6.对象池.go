package main

import (
	"fmt"
	"sync"
)

func main() {
	// New:是一个func()interface{}的委托, 用于确定池对象类型
	intPool := &sync.Pool{New: func() interface{} { return 0 }}
	fmt.Println(intPool.Get().(int))

	intPool.Put(1)
	intPool.Put(2) //覆盖赋值是无效的
	//runtime.GC()
	fmt.Println(intPool.Get().(int))
}

