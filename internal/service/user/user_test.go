package usersvc

import (
	"fmt"
	mocks "sample/twirp/mocks/model"
	userpb "sample/twirp/rpc/user"
	"testing"

	mock "github.com/stretchr/testify/mock"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	mockRepo := new(mocks.Repository)
	u := userpb.User{
		Id:       1,
		Name:     "Sample User",
		Email:    "sample@user",
		Password: "whatever",
	}
	mockRepo.On("InsertOne", mock.AnythingOfType("sqlx.DB"), mock.AnythingOfType("*userpb.User")).Return(1)
	fmt.Println("HElper is about")

	var db sqlx.DB
	userService := New(db, mockRepo)
	result, err := userService.CreateUser(nil, &u)
	fmt.Println(err)
	mockRepo.AssertExpectations(t)

	assert.Equal(t, 1, result.User.Id)

}
