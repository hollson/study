package main

import (
	"fmt"

	pb "github.com/jergoo/go-grpc-example/proto/hello" // 引入proto包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const	Address = "127.0.0.1:50052"

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Message)
}
