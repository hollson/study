package main

import (
	"log"
	"net"

	pb "github.com/jergoo/go-grpc-example/proto/hello_http"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const Address = "127.0.0.1:8888"

type helloHTTPService struct{}

// SayHello ...
func (h helloHTTPService) SayHello(ctx context.Context, in *pb.HelloHTTPRequest) (*pb.HelloHTTPResponse, error) {
	resp := new(pb.HelloHTTPResponse)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
}

var HelloHTTPService = helloHTTPService{}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloHTTPService
	pb.RegisterHelloHTTPServer(s, HelloHTTPService)

	log.Println("Listen on " + Address)

	s.Serve(listen)
}
