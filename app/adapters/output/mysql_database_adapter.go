package output

import (
	"context"
	"to-do-list-server/app/domain/model"
	"to-do-list-server/app/domain/port/output"
)

type mysqlDatabaseAdapter struct{}

func NewMysqlDatabaseAdapter() output.DatabasePort {
	return &mysqlDatabaseAdapter{}
}

func (m *mysqlDatabaseAdapter) GetUser(
	ctx context.Context,
	userId string,
) (*model.UserCredentialsDomain, error) {
	return nil, nil
}

func (m *mysqlDatabaseAdapter) GetUserByUsername(
	ctx context.Context,
	username string,
) (*model.UserCredentialsDomain, error) {
	return nil, nil
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
