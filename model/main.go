package model

import (
	userModel "sample/twirp/model/user"
	userpb "sample/twirp/rpc/user"

	"github.com/jmoiron/sqlx"
)

type DBHelper interface {
	InsertOne(db sqlx.DB, user *userpb.User) int64
	indUserByEmail(db sqlx.DB, email string) *userModel.User
	FindUserById(db sqlx.DB, id int) *userModel.User
}
