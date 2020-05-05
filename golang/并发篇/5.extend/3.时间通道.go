package main

import (
	"fmt"
	"time"
)

func main() {
	//定时器(一次计时)：sleep、after、timer(推荐)
	After()
	Timer()

	// 计时器(周期计时)
	Ticker()
}

// 定时器
func After() {
	fmt.Println("三秒倒计时开始 ：")
	at := <-time.After(time.Second * 3)
	fmt.Printf("[%s]引爆炸弹~~~ 💣 ~~~", at.Format("20060102150405"))
}

// 定时器 （Timer是可控的一种定时器）
func Timer() {
	fmt.Println("三秒后引爆炸弹：")
	stoped := false
	timer := time.NewTimer(time.Second * 1)
	timer.Reset(3 * time.Second) //重置(修改)为3秒
	//stoped = timer.Stop() //拆除定时器

	if !stoped {
		<-timer.C
		fmt.Println("炸弹💣引爆了~~~")
	} else {
		fmt.Println("炸弹拆除了~~~")
	}
}

// 计时器(周期计时)
func Ticker() {
	// 方式二: 返回时间(单)通道 <-chan Time
	tickers := time.Tick(1 * time.Second)
	for t := range tickers {
		fmt.Println("time beat: ", t.Second())
	}
}
