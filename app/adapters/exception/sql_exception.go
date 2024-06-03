package exception

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type SqlException struct {
	StatusCode int
	Err        error
}

func (s *SqlException) Error() string {
	return fmt.Sprintf("err: %v", s.Err)
}

func BuildSqlException(
	err error,
) *SqlException {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &SqlException{
			StatusCode: 404,
			Err:        err,
		}
	}
	return &SqlException{
		StatusCode: 500,
		Err:        err,
	}
}
