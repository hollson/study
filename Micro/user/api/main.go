package main

import (
	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro"
	"user/api/handler"
	"user/api/client"

	user "user/api/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.github.hollson.api.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the User srv client
		micro.WrapHandler(client.UserWrapper(service)),
	)

	//curl -H "Content-Type:application/x-www-form-urlencoded" -X POST -d "name=f546e78a46284765" http://localhost:8080/user/info

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
