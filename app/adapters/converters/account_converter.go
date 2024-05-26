package converters

import (
	"to-do-list-server/app/config/grpc"
	"to-do-list-server/app/domain/model"
)

func ConvertToModelUserCredentialsDomain(
	a *grpc.UserCredentialsRequest,
) *model.UserCredentialsDomain {
	return &model.UserCredentialsDomain{
		Username: a.Username,
		Password: a.Password,
	}
}
