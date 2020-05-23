package services

import (
	"github.com/annazhao/bookstore_users_api/domain/users"
	"github.com/annazhao/bookstore_users_api/utils/errors"
)

// CreateUser function here is used to create a user record in database
// here is where the business logic happens and defines
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser function is used to get a user from database based on user id
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	if userID <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}

	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
