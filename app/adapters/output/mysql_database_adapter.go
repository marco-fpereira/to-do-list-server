package output

import (
	"context"
	"to-do-list-server/app/adapters/converters"
	"to-do-list-server/app/adapters/output/dto"
	"to-do-list-server/app/domain/model"
	"to-do-list-server/app/domain/model/exception"
	"to-do-list-server/app/domain/port/output"

	"gorm.io/gorm"
)

type mysqlDatabaseAdapter struct {
	db *gorm.DB
}

func NewMysqlDatabaseAdapter(
	DB *gorm.DB,
) output.DatabasePort {
	return &mysqlDatabaseAdapter{
		db: DB,
	}
}

func (m *mysqlDatabaseAdapter) GetUser(
	ctx context.Context,
	userId string,
) (*model.UserCredentialsDomain, error) {
	var userCredentialsDTO *dto.UserCredentialsDTO

	result := m.db.First(&userCredentialsDTO, userId)
	if result.Error != nil {
		return nil, exception.BuildSqlException(result.Error)
	}

	user := converters.FromDtoToModelUserCredentialsDomain(userCredentialsDTO)

	return user, nil
}

func (m *mysqlDatabaseAdapter) GetUserByUsername(
	ctx context.Context,
	username string,
) (*model.UserCredentialsDomain, error) {
	var userCredentialsDTO *dto.UserCredentialsDTO

	result := m.db.Where(&dto.UserCredentialsDTO{Username: username}).First(&userCredentialsDTO)
	if result.Error != nil {
		return nil, exception.BuildSqlException(result.Error)
	}

	user := converters.FromDtoToModelUserCredentialsDomain(userCredentialsDTO)

	return user, nil
}

func (m *mysqlDatabaseAdapter) CreateUser(
	ctx context.Context,
	username string,
	password string,
) error {
	return nil
}

func (m *mysqlDatabaseAdapter) GetAllTasks(
	ctx context.Context,
	userId string,
) (*[]model.TaskDomain, error) {
	return nil, nil
}

func (m *mysqlDatabaseAdapter) CreateTask(
	ctx context.Context,
	taskMessage string,
) (*model.TaskDomain, error) {
	return nil, nil
}

func (m *mysqlDatabaseAdapter) UpdateTaskMessage(
	ctx context.Context,
	userId string,
	taskId string,
	taskMessage string,
) (*model.TaskDomain, error) {
	return nil, nil
}

func (m *mysqlDatabaseAdapter) UpdateTaskCompleteness(
	ctx context.Context,
	userId string,
	taskId string,
) error {
	return nil
}
