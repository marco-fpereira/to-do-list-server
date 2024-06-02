package exception

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ResponseException struct {
	StatusCode int
	Err        error
	Fields     []string
}

func (r *ResponseException) Error() string {
	return fmt.Sprintf("err: %v | fields: %s", r.Err, r.Fields)
}

func BuildResponseException(err error) error {
	var responseException *ResponseException
	if errors.As(err, &responseException) {
		return status.New(
			codes.Code(responseException.StatusCode),
			responseException.Error(),
		).Err()
	}

	var sqlException *SqlException
	if errors.As(err, &sqlException) {
		return status.New(
			codes.Code(sqlException.StatusCode),
			sqlException.Error(),
		).Err()
	}

	return status.New(
		codes.Internal,
		"internal server error",
	).Err()
}
