package converters

import (
	"to-do-list-server/app/config/grpc"
	"to-do-list-server/app/domain/model"
)

func FromRequestToModelUserCredentialsDomain(
	ucr *grpc.UserCredentialsRequest,
) *model.UserCredentialsDomain {
	return &model.UserCredentialsDomain{
		Username: ucr.Username,
		Password: ucr.Password,
	}
}
