package handler

import (
	"context"

	user "user/srv/proto/user"
)

type User struct{}

func (e *User) GetUserByID(ctx context.Context, req *user.Request, rsp *user.Response) error {
	userID := req.Uid
	if userID == "" {
		userID = "unknown"
	}
	userInfo := &user.UserInfo{
		Uid:  userID,
		Name: "hollson",
		Tel:  "18200000001",
	}
	rsp.Code = 200
	rsp.Error = nil
	rsp.Data = userInfo
	return nil
}

