package users

import (
	"net/http"
	"strconv"

	"github.com/annazhao/bookstore_users_api/domain/users"
	"github.com/annazhao/bookstore_users_api/services"
	"github.com/annazhao/bookstore_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateUser function here is used to parse the JSON data from request body to a new User instance, and save it into the database
func CreateUser(c *gin.Context) {
	var user users.User

	// // first way to get the user is from JSON request body
	// // *************************
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// err = json.Unmarshal(bytes, &user)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// // *************************

	// second way to get the user from JSON request body
	// *************************
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	// *************************

	// create this user, and save it into the database
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	// return back a JSON result
	c.JSON(http.StatusCreated, result)

}

// GetUser is to find user from database based on the user_id value in url parameter
func GetUser(c *gin.Context) {
	// c.Param("user_id") is to get the parameter value in url /users/:user_id
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")
// }
