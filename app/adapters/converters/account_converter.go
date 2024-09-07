package converters

import (
	"github.com/marco-fpereira/to-do-list-server/config/grpc"
	"github.com/marco-fpereira/to-do-list-server/domain/model"
)

func FromRequestToModelUserCredentialsDomain(
	ucr *grpc.UserCredentialsRequest,
) *model.UserCredentialsDomain {
	return &model.UserCredentialsDomain{
		Username: ucr.Username,
		Password: ucr.Password,
	}
}
