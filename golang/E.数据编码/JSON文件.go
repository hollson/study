package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 字符串，结构体，切片，字典分别因设为JS的字符串、对象、数组和对象
func main() {
	data := JsonMarshal()
	JsonUnmarshal(data)

	JsonEncode()
	JsonDecode()
}

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"name"`
}

// 编排
func JsonMarshal() []byte {
	var book = Book{1001, "Go教程"}
	data, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}
	json.MarshalIndent(data, "", "\t")
	fmt.Println(string(data))
	return data
}

// 反编
func JsonUnmarshal(data []byte) {
	book := new(Book)
	if err := json.Unmarshal(data, book); err != nil {
		panic(err)
	}
	fmt.Println("Unmarshal:", book)
}

// 编码
func JsonEncode() {
	// A1: 创建文件句柄
	file, err := os.OpenFile("./output.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	// A2: 关闭文件句柄
	defer file.Close()

	// B1:创建编码器
	enc := json.NewEncoder(file)

	// B2:执行编码操作
	var book = Book{1001, "Go教程"}
	if err := enc.Encode(book); err != nil {
		panic(err)
	}

	JsonDecode()
}

// 解码
func JsonDecode() {
	// A1: 创建IO句柄
	file, err := os.Open("./output.json")
	if err != nil {
		panic(err)
	}

	// A2: 关闭文件句柄
	defer file.Close()

	// B1: 创建解码器
	dec := json.NewDecoder(file)
	var book = new(Book)

	// B2: 执行解码操作
	if err := dec.Decode(book); err != nil {
		panic(err)
	}
	fmt.Println("Decode:", book)
}
