package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"to-do-list-server/app/domain/model"
	"to-do-list-server/app/domain/model/exception"
	"to-do-list-server/app/domain/port/input"
	"to-do-list-server/app/domain/port/output"
	"unicode"
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
	user, err := a.database.GetUser(ctx, userCredentials.Username)
	if err != nil {
		return err
	}

	if userAlreadyExists(user) {
		return &exception.ResponseException{
			StatusCode: http.StatusConflict,
			Err:        errors.New("user already exists"),
		}
	}

	if !passwordMatchesRequirements(user.Password) {
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

func userAlreadyExists(user *model.UserCredentialsDomain) bool {
	return user != nil
}

func passwordMatchesRequirements(password string) bool {
	matches := true

	switch {
	case !validateField(unicode.IsNumber, password):
		matches = false
	case !validateField(unicode.IsUpper, password):
		matches = false
	case !validateField(unicode.IsLower, password):
		matches = false
	case !validateField(unicode.IsSymbol, password) && !validateField(unicode.IsPunct, password):
		matches = false
	case len(password) < 8:
		matches = false
	}

	return matches
}

func validateField(fn func(r rune) bool, value string) bool {
	valid := false
	for _, c := range value {
		if fn(c) {
			valid = true
		}
	}
	return valid
}
