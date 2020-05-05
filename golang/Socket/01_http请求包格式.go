package main

import (
	"fmt"
	"net"
)

// 肌肉记忆
// 1.创建net.listen服务,network类型为tcp，端口为：8000
// 2.根据创建的监听对象，创建conn链接对象
// 3.基于Conn对象，读取字节流数据，使用字节数据接收，即n,err:=conn.Read(buf)
// 4.buf是一个字节切片，创建方式如：buf:=make(byte[],1<<20),即1M大小
func main() {
	// 创建listen-tcp监听服务，使用:8080端口
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	//基于当前监听服务，创建链接对象
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//基于当前链接，接收数据
	buf := make([]byte, 1<<20) //1M
	n,err:= conn.Read(buf)
	if err!=nil{
		panic(err)
	}

	fmt.Println(string(buf[:n]))
}
