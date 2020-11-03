package services

import (
	"os/user"

	"github.com/Dzemite/bookstore_users-api/domain/users"
	"github.com/Dzemite/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*user.User, *errors.RestErr) {
	return nil, nil
}
