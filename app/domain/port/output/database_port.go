package output

import (
	"context"

	"github.com/marco-fpereira/to-do-list-server/domain/model"
)

type DatabasePort interface {
	GetUser(
		ctx context.Context,
		userId string,
	) (*model.UserCredentialsDomain, error)

	GetUserByUsername(
		ctx context.Context,
		username string,
	) (*model.UserCredentialsDomain, error)

	CreateUser(
		ctx context.Context,
		username string,
		password string,
	) error

	GetAllTasks(
		ctx context.Context,
		userId string,
	) (*[]model.TaskDomain, error)

	GetTask(
		ctx context.Context,
		taskId string,
	) (*model.TaskDomain, error)

	CreateTask(
		ctx context.Context,
		userId string,
		taskMessage string,
		isTaskCompleted bool,
	) (*model.TaskDomain, error)

	UpdateTaskMessage(
		ctx context.Context,
		taskId string,
		taskMessage string,
	) error

	UpdateTaskCompleteness(
		ctx context.Context,
		taskId string,
		newTaskCompleteness bool,
	) error

	DeleteTask(
		ctx context.Context,
		taskId string,
	) error
}
