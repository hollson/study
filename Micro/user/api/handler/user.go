package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/util/log"
	//com_github_hollson_api_user "user/api/proto/user"

	"user/api/client"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/go-micro/api/proto"
	//user "user/api/proto/user"
)

type User struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}
//micro new helloapi --type=api --gopath=false
// User.Info is called by the API as /user/call with post body {"name": "foo"}
func (e *User) Info(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received User.Info request")

	// extract the client from the context
	userClient, ok := client.UserFromContext(ctx)
	if !ok {
		return errors.InternalServerError("io.github.entere.api.user.info", "user client not found")
	}

	// make request
	response, err := userClient.GetUserByID(ctx, &userSrv.Request{
		UserId: extractValue(req.Post["user_id"]),
	})
	if err != nil {
		return errors.InternalServerError("io.github.entere.api.user.user.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
//
//// User.Info is called by the API as /user/call with post body {"name": "foo"}
//func (e *User) Info(ctx context.Context, req *api.Request, rsp *api.Response) error {
//	log.Log("Received User.Info request")
//
//	// extract the client from the context
//	userClient, ok := client.UserFromContext(ctx)
//	if !ok {
//		return errors.InternalServerError("io.github.entere.api.user.info", "user client not found")
//	}
//
//	// make request
//	response, err := userClient.Info(ctx, &api.Request{
//		Body: "hello " + extractValue(req.Post["name"]),
//	})
//	if err != nil {
//		return errors.InternalServerError("io.github.entere.api.user.user.call", err.Error())
//	}
//
//	b, _ := json.Marshal(response)
//
//	rsp.StatusCode = 200
//	rsp.Body = string(b)
//
//	return nil
//}
