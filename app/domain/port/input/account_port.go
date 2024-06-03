package input

import (
	"context"
	"to-do-list-server/app/domain/model"
)

type AccountPort interface {
	Signup(
		ctx context.Context,
		userCredentials *model.UserCredentialsDomain,
		token string,
	) (err error)

	Login(
		ctx context.Context,
		userCredentials *model.UserCredentialsDomain,
		token string,
	) (userId string, err error)
}
