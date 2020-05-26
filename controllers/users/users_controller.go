package users

import (
	"net/http"
	"strconv"

	"github.com/annazhao/bookstore_users_api/domain/users"
	"github.com/annazhao/bookstore_users_api/services"
	"github.com/annazhao/bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userID, nil
}

// Create function here is used to parse the JSON data from request body to a new User instance, and save it into the database
func Create(c *gin.Context) {
	var user users.User
	// get the user from JSON request body
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// create this user, and save it into the database
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	// return back a JSON result
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))

}

// Get is to find user from database based on the user_id value in url parameter
func Get(c *gin.Context) {
	// c.Param("user_id") is to get the parameter value in url /users/:user_id
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	// to see whether it's a public request or private, we can see from "X-Public" in request header
	c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public") == "true"))
}

// Update is to update user in database
func Update(c *gin.Context) {
	// first get user_id from url
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.ID = userID

	// if it's PUT method, isPartial will be false; if it's PATCH, isPartial will be true
	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

// Delete is to delete user in database
func Delete(c *gin.Context) {
	// first get user_id from url
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}

	if err := services.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search is used to find users in database
func Search(c *gin.Context) {
	// in url: /internal/users/search?status=active, if we want to get "active" status, we need to use c.Query()
	status := c.Query("status")
	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	// services.Search returns a slice of user
	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}
