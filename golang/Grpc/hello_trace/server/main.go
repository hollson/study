package main

import (
	"fmt"
	"net"
	"net/http"

	pb "github.com/jergoo/go-grpc-example/proto/hello" // 引入编译生成的包

	"log"

	"golang.org/x/net/context"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const Address = "127.0.0.1:8888"

// 我实现了Proto服务端协议
type helloService struct{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	return resp, nil
}

// 开启Trance
func init() {
	grpc.EnableTracing = true
}

func main() {

	// 监听器，我实现了net.Listener接口
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// 服务器，我实现了http.Handler接口
	server := grpc.NewServer()

	// 注册服务(绑定服务器与处理程序)
	var HelloService = helloService{}
	pb.RegisterHelloServer(server, HelloService)

	// 开启trace
	go startTrace()

	fmt.Println("Listen on " + Address)
	server.Serve(listener)
}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		log.Println("存储Trance数据：",req.RemoteAddr,req.RequestURI)
		return true, true
	}
	go http.ListenAndServe(":50051", nil)
	fmt.Println("Trace listen on 50051")
}

// 访问trance页面：
// localhost:50051/debug/events
// localhost:50051/debug/requests
