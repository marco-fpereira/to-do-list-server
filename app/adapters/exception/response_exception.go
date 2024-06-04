package exception

import (
	"fmt"
)

type ResponseException struct {
	StatusCode int
	Message    string
	Fields     []string
}

func (r *ResponseException) Error() string {
	return fmt.Sprintf("err: %s | fields: %s", r.Message, r.Fields)
}

func BuildResponseException(
	statusCode int,
	message string,
) error {
	return &ResponseException{
		StatusCode: statusCode,
		Message:    message,
	}
}
