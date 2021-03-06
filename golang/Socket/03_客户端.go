package main

import (
	"fmt"
	"net"
)

// 代码模型
////一个Listen服务可以监听多个链接，一个链接可以有多个循环读写
//listen:=net.Listen("tcp",":8080") //创建监听
//for{  							  //创建多个链接
//	conn:=listen.Acception
//	for{						  //循环读写
//		//处理函数，一般作为协程，即go handler()
//		buf:=make(byte[],1024)
//		conn.read(buf)
//	}
//}

//可以访问tcp或http服务
func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("dial err = ", err)
		return
	}

	defer conn.Close()

	requestBuf := "GET /go HTTP/1.1\r\nAccept: image/gif, image/jpeg, image/pjpeg, application/x-ms-application, application/xaml+xml, application/x-ms-xbap, */*\r\nAccept-Language: zh-Hans-CN,zh-Hans;q=0.8,en-US;q=0.5,en;q=0.3\r\nUser-Agent: Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)\r\nAccept-Encoding: gzip, deflate\r\nHost: 127.0.0.1:8000\r\nConnection: Keep-Alive\r\n\r\n"

	//先发请求包，服务器才会回响应包
	conn.Write([]byte(requestBuf))

	//接收服务器回复的响应包
	buf := make([]byte, 1024*4)
	n, err1 := conn.Read(buf)
	if n == 0 {
		fmt.Println("Read err = ", err1)
		return
	}

	//打印响应报文
	fmt.Printf("#%v#", string(buf[:n]))

}
