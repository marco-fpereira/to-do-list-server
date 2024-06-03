package exception

import "fmt"

type JwtException struct {
	StatusCode int
	Err        string
}

func (j *JwtException) Error() string {
	return fmt.Sprintf("err: %v", j.Err)
}

func BuildJwtException(err error) error {
	return &JwtException{
		StatusCode: 401,
		Err:        err.Error(),
	}
}
