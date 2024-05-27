package exception

import (
	"fmt"
)

type SqlException struct {
	StatusCode int
	Err        error
}

func (s *SqlException) Error() string {
	return fmt.Sprintf("status %d: err %v", s.StatusCode, s.Err)
}

func BuildSqlException(err error) *SqlException {
	return &SqlException{
		StatusCode: 500,
		Err:        err,
	}
}
