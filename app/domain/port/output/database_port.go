package output

import (
	"context"
	"to-do-list-server/app/domain/model"
)

type DatabasePort interface {
	GetUser(ctx context.Context, userId string) (*model.UserCredentialsDomain, error)
	GetUserByUsername(ctx context.Context, username string) (*model.UserCredentialsDomain, error)
	CreateUser(ctx context.Context, username string, password string) error
	GetAllTasks(ctx context.Context, userId string) (*[]model.TaskDomain, error)
	CreateTask(ctx context.Context, taskMessage string) (*model.TaskDomain, error)
	UpdateTaskMessage(ctx context.Context, userId string, taskId string, taskMessage string) (*model.TaskDomain, error)
	UpdateTaskCompleteness(ctx context.Context, userId string, taskId string) error
}
