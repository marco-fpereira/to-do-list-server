package validators

import (
	"errors"
	"to-do-list-server/app/domain/model"
	"to-do-list-server/app/exception"
	"unicode"

	"github.com/google/uuid"
)

func ValidateStringMaxLength(
	field string,
	value string,
	maxLength int32,
) error {
	if len(value) > int(maxLength) {
		return &exception.ResponseException{
			StatusCode: 400,
			Err:        errors.New("field surpassed its max length"),
			Fields:     []string{field},
		}
	}
	return nil
}

func ValidateUUID(
	field string,
	value string,
) error {
	if _, err := uuid.Parse(value); err != nil {
		return &exception.ResponseException{
			StatusCode: 400,
			Err:        errors.New("field is not in uuid format"),
			Fields:     []string{field},
		}
	}
	return nil
}

func ValidateUserExists(
	user *model.UserCredentialsDomain,
) error {
	if user == nil {
		return &exception.ResponseException{
			StatusCode: 404,
			Err:        errors.New("user does not exist"),
			Fields:     []string{"userId"},
		}
	}
	return nil
}

func ValidateTaskExists(
	task *model.TaskDomain,
) error {
	if task == nil {
		return &exception.ResponseException{
			StatusCode: 404,
			Err:        errors.New("task does not exist"),
			Fields:     []string{"taskId"},
		}
	}
	return nil
}

func ValidateUserAlreadyExists(user *model.UserCredentialsDomain) bool {
	return user != nil
}

func ValidatePasswordMatchesRequirements(password string) bool {
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
