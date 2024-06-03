package input

import (
	"context"
	"to-do-list-server/app/domain/model"
)

type TaskPort interface {
	GetAllTasks(
		ctx context.Context,
		userId string,
		token string,
	) (*[]model.TaskDomain, error)

	CreateTask(
		ctx context.Context,
		userId string,
		taskMessage string,
		token string,
	) (*model.TaskDomain, error)

	UpdateTaskMessage(
		ctx context.Context,
		taskId string,
		taskMessage string,
		token string,
	) (*model.TaskDomain, error)

	UpdateTaskCompleteness(
		ctx context.Context,
		taskId string,
		token string,
	) error

	DeleteTask(
		ctx context.Context,
		taskId string,
		token string,
	) error
}
