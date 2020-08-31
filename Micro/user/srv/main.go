package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	//"github.com/micro/protoc-gen-micro/plugin/micro"
	"user/srv/handler"
	user "user/srv/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.github.hollson.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Subscribe Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	//// Subscribe Struct as Subscriber
	//micro.RegisterSubscriber("com.github.hollson.srv.user", service.Server(), new(subscriber.User))
	//
	//// Subscribe Function as Subscriber
	//micro.RegisterSubscriber("com.github.hollson.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
