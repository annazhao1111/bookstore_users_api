package mysqls

import (
	"strings"

	"github.com/annazhao/bookstore_users_api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const ErrorNoRows = "no rows in result set"

// ParseError is used to handle errors related to mysql database process
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
