package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio包封装了io.Reader和io.Writer对象，与io包很相似，能对高频IO操作起到缓冲作用。
func main() {
	BufWrite()
	BufRead()
}
func BufWrite() {
	// 创建文件句柄
	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	// 关闭文件句柄
	defer file.Close()

	// 1.创建读写器
	writer := bufio.NewWriterSize(file, 1<<10*4) //4096

	// 2. 执行读写操作
	data := []byte("hello world")
	if _, err := writer.Write(data); err != nil {
		panic(err)
	}

	// 3. 刷入数据
	if err := writer.Flush(); err != nil {
		panic(err)
	}
}

func BufRead() {
	// 打开文件句柄
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("文件打开失败：", err)
		return
	}
	// 关闭文件句柄
	defer file.Close()

	// 1. 创建读写器
	reader := bufio.NewReader(file)

	// 2. 执行读写操作
	buf := make([]byte, 1024)
	count, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("字节数：%d，内容：%s\n", count, string(buf))
}
