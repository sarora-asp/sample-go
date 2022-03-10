package model

import (
	userModel "sample/twirp/model/user"
)

type Repository interface {
	userModel.UserRepository
}
