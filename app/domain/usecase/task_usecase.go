package usecase

import (
	"context"
	"to-do-list-server/app/domain/model"
	"to-do-list-server/app/domain/port/input"
	"to-do-list-server/app/domain/port/output"
)

type taskUsecase struct {
	database output.DatabasePort
}

func NewTaskUseCase(
	database output.DatabasePort,
) input.TaskPort {
	return &taskUsecase{
		database: database,
	}
}

func (t *taskUsecase) GetAllTasks(ctx context.Context, userId string) (*[]model.TaskDomain, error) {
	tasks, err := t.database.GetAllTasks(ctx, userId)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *taskUsecase) CreateTask(ctx context.Context, taskMessage string) (*model.TaskDomain, error) {
	return nil, nil
}

func (t *taskUsecase) UpdateTaskMessage(ctx context.Context, userId string, taskId string, taskMessage string) (*model.TaskDomain, error) {
	return nil, nil
}

func (t *taskUsecase) UpdateTaskCompleteness(ctx context.Context, userId string, taskId string) error {
	return nil
}
