package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/protoc-gen-micro/plugin/micro"

	"github.com/micro/go-micro"
	"hello-api/handler"
	"hello-api/client"

	hello "hello-api/proto/hello"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.github.hollosn.api.hello"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Hello srv client
		micro.WrapHandler(client.HelloWrapper(service)),
	)

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
