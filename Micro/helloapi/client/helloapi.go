package client

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	helloapi "helloapi/proto/helloapi"
)

type helloapiKey struct {}

// FromContext retrieves the client from the Context
func HelloapiFromContext(ctx context.Context) (helloapi.HelloapiService, bool) {
	fmt.Println("llllllll")
	c, ok := ctx.Value(helloapiKey{}).(helloapi.HelloapiService)
	return c, ok
}
//curl http://localhost:8080/helloapi/call?name=hollson
// Client returns a wrapper for the HelloapiClient
func HelloapiWrapper(service micro.Service) server.HandlerWrapper {
	client := helloapi.NewHelloapiService("go.micro.srv.template", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, helloapiKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
