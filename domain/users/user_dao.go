package users

import (
	usersdb "github.com/annazhao/bookstore_users_api/datasources/mysql/users_db"
	"github.com/annazhao/bookstore_users_api/utils/dates"
	"github.com/annazhao/bookstore_users_api/utils/errors"
	"github.com/annazhao/bookstore_users_api/utils/mysqls"
)

// here we will have the access layer to our database

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
)

// Get method is used to retrieve the user by ID from database
func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	// QueryRow only get back 1 row from the result dataset
	// Query will get back *Rows, if using stmt.Query(user.ID), we need add defer result.Close()
	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysqls.ParseError(getErr)
	}
	return nil
}

// Save method is used to save the user into the database
func (user *User) Save() *errors.RestErr {

	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close() // this is very important

	user.DateCreated = dates.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysqls.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysqls.ParseError(err)
	}
	user.ID = userID
	return nil
}

// Update method is used to update the user in the database
func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysqls.ParseError(err)
	}
	return nil
}

// Delete method is used to update the user in the database
func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		return mysqls.ParseError(err)
	}
	return nil
}
