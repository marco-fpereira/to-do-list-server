package converters

import (
	"to-do-list-server/app/adapters/output/dto"
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

func FromDtoToModelUserCredentialsDomain(
	ucd *dto.UserCredentialsDTO,
) *model.UserCredentialsDomain {
	return &model.UserCredentialsDomain{
		UserId:   ucd.UserId,
		Username: ucd.Username,
		Password: ucd.Password,
	}
}
