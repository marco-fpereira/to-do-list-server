package input

import (
	"context"
	consts "to-do-list-server/app/adapters/consts"
	"to-do-list-server/app/adapters/converters"
	"to-do-list-server/app/adapters/exception"
	"to-do-list-server/app/config/grpc"
	"to-do-list-server/app/domain/port/input"
)

type accountAdapter struct {
	grpc.UnimplementedAccountServer
	Account input.AccountPort
}

func NewAccountAdapter(accountPort input.AccountPort) grpc.AccountServer {
	return &accountAdapter{
		Account: accountPort,
	}
}

func (a *accountAdapter) Signup(
	ctx context.Context,
	userCredentialsRequest *grpc.UserCredentialsRequest,
) (*grpc.Void, error) {
	ctx = context.WithValue(ctx, consts.REQUEST_ID, userCredentialsRequest.RequestId)
	err := a.Account.Signup(
		ctx,
		converters.FromRequestToModelUserCredentialsDomain(userCredentialsRequest),
		userCredentialsRequest.Token,
	)
	if err != nil {
		return nil, exception.BuildResponseException(err)
	}

	return &grpc.Void{}, nil
}

func (a *accountAdapter) Login(
	ctx context.Context,
	userCredentialsRequest *grpc.UserCredentialsRequest,
) (*grpc.UserId, error) {
	ctx = context.WithValue(ctx, consts.REQUEST_ID, userCredentialsRequest.RequestId)
	userId, err := a.Account.Login(
		ctx,
		converters.FromRequestToModelUserCredentialsDomain(userCredentialsRequest),
		userCredentialsRequest.Token,
	)
	if err != nil {
		return nil, exception.BuildResponseException(err)
	}

	return &grpc.UserId{UserId: userId}, nil
}
