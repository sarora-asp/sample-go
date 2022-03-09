package usersvc

import (
	mocks "sample/twirp/mocks/model"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestCreateUser(t *testing.T) {
	mockHelper := new(mocks.DBHelper)

	mockHelper.On("InsertOne").Return(1)
	var db sqlx.DB
	userService := New(db)

}
