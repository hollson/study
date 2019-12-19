package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/util/log"

	"helloapi/client"
	"github.com/micro/go-micro/errors"
	//helloapi "helloapi/proto/helloapi"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/api/proto"
	//helloapi "path/to/service/proto/helloapi"
)

type Helloapi struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Helloapi.Call is called by the API as /helloapi/call with post body {"name": "foo"}
func (e *Helloapi) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Helloapi.Call request")

	// extract the client from the context
	helloapiClient, ok := client.HelloapiFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.helloapi.helloapi.call", "helloapi client not found")
	}

	// make request
	response, err := helloapiClient.Call(ctx, &go_api.Request{
		Body: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.helloapi.helloapi.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
