//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2020-01-21
// @ Version: 1.0.0
//
// Here's the code description...
//-------------------------------------------------------------------------------------

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//O_RDONLY：只读模式
//O_WRONLY：只写模式
//O_RDWR：  读写模式

//O_APPEND：尾部追加
//O_CREATE：如果不存在将创建一个新文件；
//O_EXCL：和 O_CREATE 配合使用，文件必须不存在，否则返回一个错误；
//O_SYNC：同步写入
//O_TRUNC：清空文件

// Unix:   只有换行， 即"\n"；
// Windows 回车+换行，即“\r\n”；

func main() {
	//Create()
	Read()
	Copy()
}

// 写入数据
func Create() {
	// 创建IO句柄
	path := "./output.txt"
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	//关闭文件句柄
	defer file.Close()

	// 1. 创建读写器
	write := bufio.NewWriter(file)

	// 2. 执行读写操作
	for i := 0; i < 5; i++ {
		write.WriteString("hello world\n")
	}

	// 3. 保存数据
	write.Flush()
}

// 追加数据
func Append() {
	// 创建IO句柄
	path := "./output.txt"
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	// 关闭文件句柄
	defer file.Close()

	// 1. 创建读写器
	write := bufio.NewWriter(file)

	// 2. 执行读写操作(附加)
	for i := 0; i < 5; i++ {
		write.WriteString("append hello world \r")
	}

	// 3. 保存数据
	write.Flush()
}

// 3. 读取数据
func Read() {
	// 创建IO句柄
	path := "./output.txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	// 关闭文件句柄
	defer file.Close()

	// 1. 创建读写器
	reader := bufio.NewReader(file)

	// 2. (循环)执行读写操作
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}
}

// 4. 拷贝数据
func Copy() {
	path1 := "./output.txt"
	path2 := "./backup.txt"
	data, err := ioutil.ReadFile(path1)
	if err != nil {
		fmt.Printf("文件打开失败,%v\n", err)
		return
	}

	err = ioutil.WriteFile(path2, data, 0666)
	if err != nil {
		fmt.Printf("文件打开失败,%v\n", err)
	}
}
