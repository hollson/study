package main

import (
	"math"
	"sync"
)

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	println(id, x)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum(id)
		}(i)
	}
	wg.Wait()
}

// shell运行程序，并调整GOMAXPROCS，观察结构变化

//GOMAXPROCS=2  time -p go run main.go
//0 9223372030412324865
//1 9223372030412324865
//real         1.72 // 程序开始到结束时间差 (非 CPU 时间)
//user         3.20 // ⽤用户态所使⽤用 CPU 时间片 (多核累加)
//sys          0.09 // 内核态所使⽤用 CPU 时间片