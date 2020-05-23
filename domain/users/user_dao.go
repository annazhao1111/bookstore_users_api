package users

import (
	"fmt"

	"github.com/annazhao/bookstore_users_api/utils/errors"
)

// here we will have the access layer to our database

// before using database, we can use a in-memory data structure to store users
var usersDB = make(map[int64]*User)

// Get method is used to retrieve the user by ID from database
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

// Save method is used to save the user into the database
func (user *User) Save() *errors.RestErr {
	// if the user id already exists in database, we cannot save it
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	usersDB[user.ID] = user
	return nil
}
