package main

import (
	"fmt"
	"net"

	pb "github.com/jergoo/go-grpc-example/proto/hello" // 引入编译生成的包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	fmt.Println(in.Name)
	return resp, nil
}


// 1. 声明监听器,即net.Listener
// 2. 声明服务，即http.Server
// 3. 注册服务路由，并映射响应的处理程序
// 4. 启动服务(指定监听器)
func main() {
	// 组件一： net.Listener 接口
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}

	// 组件二： http.Handler接口
	server := grpc.NewServer()

	// 注册server的路由处理程序
	pb.RegisterHelloServer(server, HelloService) //hello服务
	// pb.RegisterUserServer(s, HelloService)  //用户服务
	fmt.Println(Address)

	// 运行服务
	server.Serve(listen)
}


