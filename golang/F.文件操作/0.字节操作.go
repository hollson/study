//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2020-01-22
// @ Version: 1.0.0
//
// Here's the code description...
//-------------------------------------------------------------------------------------

package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// 将目标数据迁移至目标容器(字节数组)
	r := bytes.NewReader([]byte("hello"))

//bytes.NewBufferString("hello")

	tar := make([]byte, 1)
	for {
		// 将源数据拷贝到目标容器,以(字节)游标的方式读取源数据
		_, err := r.Read(tar)
		if err ==io.EOF {
			break
		}
		fmt.Println(string(tar))
	}
}
