package users

import (
	"encoding/json"
)

// here is how your domain (user) will be presented to the client

// PublicUser struct is facing to public request
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// PrivateUser struct is facing internal request
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// Marshal is used to decide what user information should be returned based on different type of request
func (user *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	// return PrivateUser{
	// 	ID:          user.ID,
	// 	FirstName:   user.FirstName,
	// 	LastName:    user.LastName,
	// 	Email:       user.Email,
	// 	DateCreated: user.DateCreated,
	// 	Status:      user.Status,
	// }

	// if User and PrivateUser have the same json structure, we can use another way to return PrivateUser:
	userJSON, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}

// Marshal is used to returen a slice of different type of users
func (users Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshal(isPublic)
	}
	return result
}
