package main

import (
	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro"
	"helloapi/handler"
	"helloapi/client"

	helloapi "helloapi/proto/helloapi"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.helloapi"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Helloapi srv client
		micro.WrapHandler(client.HelloapiWrapper(service)),
	)
// curl -H "Content-Type:application/x-www-form-urlencoded" -X POST -d "name=f546e78a46284765" http://localhost:8080/Helloapi/call

// Subscribe Handler
	helloapi.RegisterHelloapiHandler(service.Server(), new(handler.Helloapi))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
