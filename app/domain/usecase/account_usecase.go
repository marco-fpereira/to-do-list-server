package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"to-do-list-server/app/domain/model"
	"to-do-list-server/app/domain/port/input"
	"to-do-list-server/app/domain/port/output"
	"to-do-list-server/app/domain/validators"
	"to-do-list-server/app/exception"
)

type accountUseCase struct {
	database output.DatabasePort
}

func NewAccountUseCase(
	database output.DatabasePort,
) input.AccountPort {
	return &accountUseCase{
		database: database,
	}
}

func (a *accountUseCase) Signup(ctx context.Context, userCredentials *model.UserCredentialsDomain) error {
	user, err := a.database.GetUserByUsername(ctx, userCredentials.Username)
	if err != nil {
		return err
	}

	if validators.ValidateUserAlreadyExists(user) {
		return &exception.ResponseException{
			StatusCode: http.StatusConflict,
			Err:        errors.New("user already exists"),
		}
	}

	if !validators.ValidatePasswordMatchesRequirements(user.Password) {
		return &exception.ResponseException{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("password is not strong enough"),
		}
	}

	// todo: encrypt password before saving it
	err = a.database.CreateUser(ctx, userCredentials.Username, userCredentials.Password)

	if err != nil {
		return err
	}

	return nil
}

func (a *accountUseCase) Login(ctx context.Context, userCredentials *model.UserCredentialsDomain) (string, error) {
	// todo: implement login
	user, err := a.database.GetUserByUsername(ctx, userCredentials.Username)
	if err != nil {
		return "", err
	}
	fmt.Println(user)
	return "", nil
}
