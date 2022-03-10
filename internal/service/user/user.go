package usersvc

import (
	"context"
	"fmt"
	utils "sample/twirp/internal/utils"
	userpb "sample/twirp/rpc/user"

	"google.golang.org/protobuf/types/known/timestamppb"

	"sample/twirp/model"

	"github.com/jmoiron/sqlx"
	"github.com/twitchtv/twirp"
)

type userServiceProvider struct {
	db   sqlx.DB
	repo model.Repository
}

func New(db sqlx.DB, repo model.Repository) *userServiceProvider {
	return &userServiceProvider{
		db:   db,
		repo: repo,
	}
}

// 	CreateUser(context.Context, *User) (*Response, error)
func (u userServiceProvider) CreateUser(_ context.Context, user *userpb.User) (*userpb.Response, error) {
	pass, err := utils.HashPassword(user.Password)

	if err != nil {
		return &userpb.Response{
			Code:    400,
			Success: false,
			Msg:     "Unable to create user!",
			User:    nil,
		}, nil
	}
	user.Password = pass
	fmt.Println("IN SVC", user)
	num := u.repo.InsertOne(u.db, user)
	fmt.Println("NUM", num)
	if num < 1 {
		return &userpb.Response{
			Code:    400,
			Success: false,
			Msg:     "Unable to insert user!",
			User:    nil,
		}, nil
	}
	return &userpb.Response{
		Code:    201,
		Success: true,
		Msg:     "Inserted successfully!",
		User:    nil,
	}, nil
}

// 	GetUser(context.Context, *Request) (*Response, error)
func (u *userServiceProvider) GetUser(_ context.Context, _ *userpb.Request) (*userpb.Response, error) {
	user := u.repo.FindUserById(u.db, 3)
	x := &userpb.User{
		Id:        int32(user.Id),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
	return &userpb.Response{
		Code:    200,
		Success: true,
		Msg:     "this is not working",
		User:    x,
	}, nil
}

// Login(context.Context, *LoginReq) (*Response, error)
func (u *userServiceProvider) Login(_ context.Context, req *userpb.LoginReq) (*userpb.Response, error) {
	if req.Email == "" || len(req.Email) < 3 {
		return nil, twirp.NewError(twirp.InvalidArgument, "Email is not valid")
	}
	if req.Password == "" || len(req.Password) < 3 {
		return nil, twirp.NewError(twirp.InvalidArgument, "Password is not valid")
	}

	usr := u.repo.FindUserByEmail(u.db, req.Email)
	isValidPwd := utils.CheckPasswordHash(req.Password, usr.Password)
	if !isValidPwd {
		return nil, twirp.NewError(twirp.InvalidArgument, "Invalid username or password")
	}

	result := &userpb.User{
		Id:   int32(usr.Id),
		Name: usr.Name,
	}

	return &userpb.Response{
		Code:    200,
		Success: true,
		Msg:     "this is not working",
		User:    result,
	}, nil
}
