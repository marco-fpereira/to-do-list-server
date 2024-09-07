package usecase

import (
	"context"

	"github.com/marco-fpereira/to-do-list-server/domain/model"
	"github.com/marco-fpereira/to-do-list-server/domain/port/input"
	"github.com/marco-fpereira/to-do-list-server/domain/port/output"
	"github.com/marco-fpereira/to-do-list-server/domain/validators"

	"github.com/marco-fpereira/to-do-list-server/adapters/exception"
)

type accountUseCase struct {
	auth     output.AuthenticationPort
	crypt    output.CryptographyPort
	database output.DatabasePort
}

func NewAccountUseCase(
	auth output.AuthenticationPort,
	crypt output.CryptographyPort,
	database output.DatabasePort,
) input.AccountPort {
	return &accountUseCase{
		auth:     auth,
		crypt:    crypt,
		database: database,
	}
}

func (a *accountUseCase) Signup(
	ctx context.Context,
	userCredentials *model.UserCredentialsDomain,
	token string,
) error {
	if isValid, err := a.auth.ValidateToken(token); !isValid {
		return err
	}

	user, err := a.database.GetUserByUsername(ctx, userCredentials.Username)
	if err != nil {
		return err
	}

	if validators.ValidateUserAlreadyExists(user) {
		return &exception.ResponseException{
			StatusCode: 409,
			Message:    "user already exists",
			Fields:     []string{"Username"},
		}
	}

	if !validators.ValidatePasswordMatchesRequirements(userCredentials.Password) {
		return &exception.ResponseException{
			StatusCode: 400,
			Message:    "password is not strong enough",
		}
	}

	encryptedPassword, err := a.crypt.EncryptKey(userCredentials.Password)
	if err != nil {
		return err
	}

	err = a.database.CreateUser(ctx, userCredentials.Username, encryptedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (a *accountUseCase) Login(
	ctx context.Context,
	userCredentials *model.UserCredentialsDomain,
	token string,
) (string, error) {
	if isValid, err := a.auth.ValidateToken(token); !isValid {
		return "", err
	}

	user, err := a.database.GetUserByUsername(ctx, userCredentials.Username)
	if err != nil {
		return "", err
	}

	if validClaim, err := a.auth.ValidateClaim(token, "userId", user.UserId); !validClaim {
		return "", err
	}

	isPasswordCorrect := a.crypt.VerifyEncryptedKey(userCredentials.Password, user.Password)

	if !isPasswordCorrect {
		return "", &exception.ResponseException{
			StatusCode: 400,
			Message:    "username or password is incorrect",
		}
	}
	return user.UserId, nil
}
