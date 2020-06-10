package main

import (
	"fmt"
	"time"

	pb "github.com/jergoo/go-grpc-example/proto/hello" // 引入proto包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	for i := 0; i < 100; i++ {
		Hello(fmt.Sprintf("世界 %d",i))
		time.Sleep(time.Millisecond*20)
	}
}

func Hello(name string){
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())  // 不安全连接
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: name}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Message)
}
