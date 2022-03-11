package usersvc

import (
	"context"
	"fmt"
	helper "sample/twirp/internal/helper"
	mocks "sample/twirp/mocks"
	usermodel "sample/twirp/model/user"
	userpb "sample/twirp/rpc/user"
	"testing"
	"time"

	mock "github.com/stretchr/testify/mock"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockHelper := new(mocks.Helper)

	u := userpb.User{
		Id:       1,
		Name:     "Sample User",
		Email:    "sample@user",
		Password: "whatever",
	}
	mockRepo.On("InsertOne", mock.AnythingOfType("sqlx.DB"), mock.AnythingOfType("*user.User")).Return(func(db sqlx.DB, u *userpb.User) int64 {
		return 1
	})

	var db sqlx.DB
	userService := New(db, mockRepo, mockHelper)

	result, err := userService.CreateUser(context.TODO(), &u)
	fmt.Println("RESP", result, err)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, int32(201), result.Code)

}

func TestGetUser(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockHelper := new(mocks.Helper)

	mockRepo.On("FindUserById", mock.AnythingOfType("sqlx.DB"), mock.AnythingOfType("int")).Return(func(db sqlx.DB, id int) *usermodel.User {
		return &usermodel.User{
			Id:        id,
			Name:      "Test User",
			Email:     "sample@user.com",
			Password:  "helloworld",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	})

	var db sqlx.DB
	userService := New(db, mockRepo, mockHelper)

	result, err := userService.GetUser(context.TODO(), &userpb.Request{})
	fmt.Println("RESP", result, err)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, int32(3), result.User.Id)

}

func TestAPICall(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockHelper := new(mocks.Helper)

	mockHelper.On("GetAPIResponse").Return(func() helper.APIResponse {
		return helper.APIResponse{
			Data: []struct {
				ID        int    `json:"id"`
				Email     string `json:"email"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Avatar    string `json:"avatar"`
			}{
				{
					ID:        1,
					Email:     "hello@world",
					FirstName: "Sample name",
					LastName:  "Whatever",
				},
				{
					ID:        2,
					Email:     "hello@world",
					FirstName: "Sample name",
					LastName:  "Whatever",
				},
			},
		}
	})

	var db sqlx.DB
	userService := New(db, mockRepo, mockHelper)

	result, err := userService.ApiCall(context.TODO(), &userpb.Empty{})
	fmt.Println("API CALL", result, err)

	mockHelper.AssertExpectations(t)

	assert.Equal(t, int32(200), result.Code)

}
