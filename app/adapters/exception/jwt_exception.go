package exception

import (
	"fmt"
)

type JwtException struct {
	StatusCode int
	Message    string
}

func (j *JwtException) Error() string {
	return fmt.Sprintf("err: %v", j.Message)
}

func BuildJwtException(
	statusCode int,
	message string,
) error {
	return &JwtException{
		StatusCode: statusCode,
		Message:    message,
	}
}
