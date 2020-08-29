package main

import (
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "github.com/jergoo/go-grpc-example/proto/hello_http"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// grpc服务地址
	endpoint := "127.0.0.1:8080"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// HTTP转grpc
	err := gw.RegisterHelloHTTPHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		log.Println("Register handler err:%v\n", err)
	}

	log.Println("HTTP Listen on 8081")
	http.ListenAndServe(":8081", mux)
}
