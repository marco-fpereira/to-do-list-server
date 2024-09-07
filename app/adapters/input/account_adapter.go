package input

import (
	"context"

	consts "github.com/marco-fpereira/to-do-list-server/adapters/consts"
	"github.com/marco-fpereira/to-do-list-server/adapters/converters"
	"github.com/marco-fpereira/to-do-list-server/adapters/exception/handler"
	"github.com/marco-fpereira/to-do-list-server/config/grpc"
	"github.com/marco-fpereira/to-do-list-server/config/logger"
	"github.com/marco-fpereira/to-do-list-server/domain/port/input"

	"go.uber.org/zap"
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
	tags := []zap.Field{
		zap.String("Username", userCredentialsRequest.Username),
		zap.String("RequestId", userCredentialsRequest.RequestId),
	}
	logger.Info(ctx, "Signing up user.", tags...)

	err := a.Account.Signup(
		ctx,
		converters.FromRequestToModelUserCredentialsDomain(userCredentialsRequest),
		userCredentialsRequest.Token,
	)
	if err != nil {
		logger.Error(ctx, "Error signing up user.", err, tags...)
		return nil, handler.HandleException(err)
	}
	logger.Info(ctx, "User successfully signed up.", tags...)

	return &grpc.Void{}, nil
}

func (a *accountAdapter) Login(
	ctx context.Context,
	userCredentialsRequest *grpc.UserCredentialsRequest,
) (*grpc.UserId, error) {
	ctx = context.WithValue(ctx, consts.REQUEST_ID, userCredentialsRequest.RequestId)
	tags := []zap.Field{
		zap.String("Username", userCredentialsRequest.Username),
		zap.String("RequestId", userCredentialsRequest.RequestId),
	}

	logger.Info(ctx, "Logging in user.", tags...)

	userId, err := a.Account.Login(
		ctx,
		converters.FromRequestToModelUserCredentialsDomain(userCredentialsRequest),
		userCredentialsRequest.Token,
	)
	if err != nil {
		logger.Error(ctx, "Error logging in user.", err, tags...)
		return nil, handler.HandleException(err)
	}

	tags = append(tags, zap.String("userId", userId))
	logger.Info(ctx, "User successfully logged in.", tags...)
	return &grpc.UserId{UserId: userId}, nil
}
