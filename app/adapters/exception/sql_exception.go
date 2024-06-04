package exception

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type SqlException struct {
	StatusCode int
	Message    string
}

func (s *SqlException) Error() string {
	return fmt.Sprintf("err: %s", s.Message)
}

func BuildSqlException(
	err error,
) *SqlException {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &SqlException{
			StatusCode: 404,
			Message:    err.Error(),
		}
	}
	return &SqlException{
		StatusCode: 500,
		Message:    err.Error(),
	}
}
