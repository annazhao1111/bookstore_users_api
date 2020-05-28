package services

import (
	"github.com/annazhao/bookstore_users_api/domain/users"
	"github.com/annazhao/bookstore_users_api/utils/cryptos"
	"github.com/annazhao/bookstore_users_api/utils/dates"
	"github.com/annazhao/bookstore_users_api/utils/errors"
)

// UsersService is the type of usersService, as well as the type of usersServiceInterface
var UsersService usersService

type usersService struct {
}

// because type usersService has all the method that usersServiceInterface has,
// so usersService is also the type of usersServiceInterface
type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, user users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
}

// CreateUser function here is used to create a user record in database
// here is where the business logic happens and defines
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive // default status for new users is active
	user.DateCreated = dates.GetNowDBFormat()
	user.Password = cryptos.GetMd5(user.Password) // hashed password

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser function is used to get a user from database based on user id
func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateUser function is used to update a user in database
func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{ID: user.ID}
	if err := current.Get(); err != nil {
		return nil, err
	}

	// if we only want to update partial field from JSON request, we need to use PATCH method
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

// DeleteUser function is used to delete a user in database
func (s *usersService) DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete() // Delete() is a method
}

// Search function is used to find users in database based on status
func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	user := &users.User{}
	// the following code is the same as
	return user.FindByStatus(status)

	// userSlice, err := user.FindByStatus(status)
	// if err != nil {
	// 	return nil, err
	// }
	// return userSlice, nil
}
