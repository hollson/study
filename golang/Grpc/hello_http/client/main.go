package main

import (
	"log"

	pb "github.com/jergoo/go-grpc-example/proto/hello_http"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const Address = "127.0.0.1:8888"

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	// 初始化客户端
	cli := pb.NewHelloHTTPClient(conn)

	// 调用方法
	reqBody := new(pb.HelloHTTPRequest)
	reqBody.Name = "gRPC"

	r, err := cli.SayHello(context.Background(), reqBody)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(r.Message)
}

// OR: curl -X POST -k http://localhost:8080/example/echo -d '{"name": "gRPC-HTTP is working!"}'
