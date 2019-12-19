package handler

import (
	"github.com/micro/go-micro/errors"
	"hello-api/client"
)


import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/util/log"
	"hello-api/client"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/go-micro/api/proto"
	hello "path/to/service/proto/hello"
)
type Hello struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Hello.Call is called by the API as /hello/call with post body {"name": "foo"}
func (e *Hello) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Hello.Call request")

	// extract the client from the context
	helloClient, ok := client.HelloFromContext(ctx)
	if !ok {
		return errors.InternalServerError("com.github.hollosn.api.hello.hello.call", "hello client not found")
	}

	// make request
	response, err := helloClient.Call(ctx, &hello.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("com.github.hollosn.api.hello.hello.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
