package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func main() {
	//实例化对象
	info := Website{"C语言中文网", "http://c.biancheng.net/golang/", []string{"Go语言入门教程", "Golang入门教程"}}
	f, err := os.Create("./info.xml")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer f.Close()
	//序列化到文件中
	encoder := xml.NewEncoder(f)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误：", err.Error())
		return
	} else {
		fmt.Println("编码成功")
	}
}


func main2() {
	//打开xml文件
	file, err := os.Open("./info.xml")
	if err != nil {
		fmt.Printf("文件打开失败：%v", err)
		return
	}
	defer file.Close()

	info := Website{}
	//创建 xml 解码器
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Printf("解码失败：%v", err)
		return
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}