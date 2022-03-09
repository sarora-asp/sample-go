package service

import (
	"context"
	userpb "sample/twirp/rpc/user"
)

type Service interface {
	CreateUser(_ context.Context, user *userpb.User) (*userpb.Response, error)
	GetUser(_ context.Context, _ *userpb.Request) (*userpb.Response, error)
	Login(_ context.Context, req *userpb.LoginReq) (*userpb.Response, error)
}
