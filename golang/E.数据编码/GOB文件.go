package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

//Gob是以Golang为特定语言的二进制序列化/反序列化数据格式,在编码和解码过程中用到了Go的反射。
func main() {
	Encode()
	Decode()
}

func Encode() {
	info := map[string]string{"title": "hello world"}
	file, _ := os.OpenFile("output.gob", os.O_RDWR|os.O_CREATE, 0777)
	defer file.Close()

	enc := gob.NewEncoder(file)
	if err := enc.Encode(info); err != nil {
		fmt.Println(err)
	}
}

func Decode() {
	var info map[string]string
	file, _ := os.Open("output.gob")

	dec := gob.NewDecoder(file)
	dec.Decode(&info)
	fmt.Println(info)
}
