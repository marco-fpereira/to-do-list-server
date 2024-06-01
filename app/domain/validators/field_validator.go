package validators

import (
	"to-do-list-server/app/domain/model"
	"to-do-list-server/app/domain/model/exception"

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
			Fields:     []string{"taskId"},
		}
	}
	return nil
}
