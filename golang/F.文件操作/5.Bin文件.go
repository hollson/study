package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type Website struct {
	Url int32
}

func main() {
	BinWrite()
	BinRead()
}

func BinWrite() {
	// 创建IO句柄
	file, err := os.Create("output.bin")
	if err != nil {
		fmt.Println("文件创建失败 ", err.Error())
		return
	}
	defer file.Close()

	for i := 0; i < 10; i++ {
		info := Website{
			int32(i),
		}

		// 写入缓冲
		var buf bytes.Buffer
		binary.Write(&buf, binary.LittleEndian, info)

		// 写入文件
		_, err = file.Write(buf.Bytes())
		if err != nil {
			fmt.Println("数据写入失败", err.Error())
			return
		}
	}
}

func BinRead() {
	file, err := os.Open("output.bin")
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}

	m := Website{}
	for i := 1; i <= 10; i++ {
		data := NextBytes(file, 4)

		// 读入缓冲
		buf := bytes.NewBuffer(data)
		err = binary.Read(buf, binary.LittleEndian, &m) //小端
		if err != nil {
			fmt.Println("二进制文件读取失败", err)
			return
		}

		fmt.Println("第", i, "个值为：", m)
	}
}

// 读取字节数据
func NextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println("解码失败", err)
	}
	return bytes
}
