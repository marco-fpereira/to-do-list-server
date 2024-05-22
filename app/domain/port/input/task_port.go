package input

import (
	"context"
	"to-do-list-server/app/domain/model"
)

type TaskPort interface {
	GetAllTasks(ctx context.Context, userId string) (*[]model.TaskDomain, error)
	CreateTask(ctx context.Context, taskMessage string) (*model.TaskDomain, error)
	UpdateTaskMessage(ctx context.Context, userId string, taskId string, taskMessage string) (*model.TaskDomain, error)
	UpdateTaskCompleteness(ctx context.Context, userId string, taskId string)
}
