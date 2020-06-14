package main

import (
	"fmt"
	"log"

	pb "github.com/jergoo/go-grpc-example/proto/hello" // 引入proto包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const Address = "127.0.0.1:8888"

func main() {
	// 连接对象
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

// Examples:
//	Dial("tcp", "golang.org:http")
//	Dial("tcp", "192.0.2.1:http")
//	Dial("tcp", "198.51.100.1:80")
//	Dial("udp", "[2001:db8::1]:domain")
//	Dial("udp", "[fe80::1%lo0]:53")
//	Dial("tcp", ":80")
//
// For IP networks, the network must be "ip", "ip4" or "ip6" followed
// by a colon and a literal protocol number or a protocol name, and
// the address has the form "host". The host must be a literal IP
// address or a literal IPv6 address with zone.
// It depends on each operating system how the operating system
// behaves with a non-well known protocol number such as "0" or "255".
//
// Examples:
//	Dial("ip4:1", "192.0.2.1")
//	Dial("ip6:ipv6-icmp", "2001:db8::1")
//	Dial("ip6:58", "fe80::1%lo0")

// Proto客户端四要素：client，request，handler，response
func Foo() {
	// RPC拨号
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	cli := pb.NewHelloClient(conn)
cli.SayHello(context.TODO(),nil)
}
