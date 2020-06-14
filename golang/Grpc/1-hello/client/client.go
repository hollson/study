package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-example/proto/hello"

)

const Address = "127.0.0.1:8888"

func main() {
	// Grpc拨号
	conn, err := grpc.Dial(Address, grpc.WithInsecure())  // 无证书
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// 创建客户端
	cli := pb.NewHelloClient(conn)

	// 执行handler
	rsp, err := cli.SayHello(context.TODO(), &pb.HelloRequest{Name: "世界"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(rsp.Message)
}
