package exception

import (
	"errors"
	"fmt"
)

type ResponseException struct {
	StatusCode int
	Err        error
	Fields     []string
}

func (r *ResponseException) Error() string {
	return fmt.Sprintf("status: %d | err: %v | fields: %s", r.StatusCode, r.Err, r.Fields)
}

func BuildResponseException(err error) *ResponseException {
	var responseException *ResponseException
	if errors.As(err, &responseException) {
		return responseException
	}
	return &ResponseException{
		StatusCode: 500,
		Err:        err,
	}
}
