package usecase

import (
	"context"
	"to-do-list-server/app/domain/model"
	"to-do-list-server/app/domain/port/input"
	"to-do-list-server/app/domain/port/output"
	"to-do-list-server/app/domain/validators"
)

const MAX_MESSAGE_LENGTH = 512

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
	if err := validators.ValidateUUID("userId", userId); err != nil {
		return nil, err
	}

	tasks, err := t.database.GetAllTasks(ctx, userId)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *taskUsecase) CreateTask(
	ctx context.Context,
	userId string,
	taskMessage string,
) (*model.TaskDomain, error) {
	if err := validators.ValidateUUID("userId", userId); err != nil {
		return nil, err
	}

	if err := validators.ValidateStringMaxLength("taskMessage", taskMessage, MAX_MESSAGE_LENGTH); err != nil {
		return nil, err
	}

	user, err := t.database.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	if err := validators.ValidateUserExists(user); err != nil {
		return nil, err
	}

	task, err := t.database.CreateTask(ctx, userId, taskMessage, false)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *taskUsecase) UpdateTaskMessage(
	ctx context.Context,
	taskId string,
	taskMessage string,
) (*model.TaskDomain, error) {
	if err := validators.ValidateUUID("taskId", taskId); err != nil {
		return nil, err
	}

	if err := validators.ValidateStringMaxLength("taskMessage", taskMessage, MAX_MESSAGE_LENGTH); err != nil {
		return nil, err
	}

	task, err := t.database.GetTask(ctx, taskId)
	if err != nil {
		return nil, err
	}
	if err := validators.ValidateTaskExists(task); err != nil {
		return nil, err
	}

	err = t.database.UpdateTaskMessage(ctx, taskId, taskMessage)
	if err != nil {
		return nil, err
	}
	task.TaskMessage = taskMessage
	return task, nil
}

func (t *taskUsecase) UpdateTaskCompleteness(ctx context.Context, taskId string) error {
	if err := validators.ValidateUUID("taskId", taskId); err != nil {
		return err
	}

	task, err := t.database.GetTask(ctx, taskId)
	if err != nil {
		return err
	}
	if err := validators.ValidateTaskExists(task); err != nil {
		return err
	}

	err = t.database.UpdateTaskCompleteness(ctx, taskId, !task.IsTaskCompleted)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskUsecase) DeleteTask(ctx context.Context, taskId string) error {
	task, err := t.database.GetTask(ctx, taskId)
	if err != nil {
		return err
	}
	if err = validators.ValidateTaskExists(task); err != nil {
		return err
	}

	err = t.database.DeleteTask(ctx, taskId)
	if err != nil {
		return err
	}
	return nil
}
