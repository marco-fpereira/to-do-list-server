package input

import (
	"context"

	"github.com/marco-fpereira/to-do-list-server/domain/model"
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
