package input

import (
	"context"
	consts "to-do-list-server/app/adapters/consts"
	"to-do-list-server/app/adapters/converters"
	"to-do-list-server/app/adapters/exception/handler"
	"to-do-list-server/app/config/grpc"
	"to-do-list-server/app/domain/port/input"

	log "github.com/sirupsen/logrus"
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
	tags := log.Fields{
		"Username":  userCredentialsRequest.Username,
		"RequestId": userCredentialsRequest.RequestId,
	}
	log.WithFields(tags).Info("Signing up user.")

	err := a.Account.Signup(
		ctx,
		converters.FromRequestToModelUserCredentialsDomain(userCredentialsRequest),
		userCredentialsRequest.Token,
	)
	if err != nil {
		tags["error"] = err
		log.WithFields(tags).Error("Error signing up user.")
		return nil, handler.HandleException(err)
	}
	log.WithFields(tags).Info("User successfully signed up.")

	return &grpc.Void{}, nil
}

func (a *accountAdapter) Login(
	ctx context.Context,
	userCredentialsRequest *grpc.UserCredentialsRequest,
) (*grpc.UserId, error) {
	ctx = context.WithValue(ctx, consts.REQUEST_ID, userCredentialsRequest.RequestId)
	tags := log.Fields{
		"Username":  userCredentialsRequest.Username,
		"RequestId": userCredentialsRequest.RequestId,
	}
	log.WithFields(tags).Info("Logging in user.")

	userId, err := a.Account.Login(
		ctx,
		converters.FromRequestToModelUserCredentialsDomain(userCredentialsRequest),
		userCredentialsRequest.Token,
	)
	if err != nil {
		tags["error"] = err
		log.WithFields(tags).Error("Error logging in user.")
		return nil, handler.HandleException(err)
	}

	tags["userId"] = userId
	log.WithFields(tags).Info("User successfully logged in.")
	return &grpc.UserId{UserId: userId}, nil
}
