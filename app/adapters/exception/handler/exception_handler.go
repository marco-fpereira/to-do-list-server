package handler

import (
	"errors"
	"to-do-list-server/app/adapters/exception"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleException(err error) error {
	var responseException *exception.ResponseException
	var sqlException *exception.SqlException
	var jwtException *exception.JwtException

	switch {
	case errors.As(err, &responseException):
		return status.New(
			codes.Code(responseException.StatusCode),
			responseException.Error(),
		).Err()
	case errors.As(err, &sqlException):
		return status.New(
			codes.Code(sqlException.StatusCode),
			sqlException.Error(),
		).Err()
	case errors.As(err, &jwtException):
		return status.New(
			codes.Code(jwtException.StatusCode),
			jwtException.Error(),
		).Err()
	default:
		return status.New(
			codes.Internal,
			"internal server error",
		).Err()
	}
}
