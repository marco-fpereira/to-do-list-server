package input

import (
	"context"
	"to-do-list-server/app/adapters/converters"
	"to-do-list-server/app/config/grpc"
	"to-do-list-server/app/domain/model/exception"
	"to-do-list-server/app/domain/port/input"
)

type Key string

const (
	RequestId Key = "RequestId"
)

type taskAdapter struct {
	grpc.UnimplementedTaskServer
	Task input.TaskPort
}

func NewTaskAdapter(
	taskPort input.TaskPort,
) grpc.TaskServer {
	return &taskAdapter{
		Task: taskPort,
	}
}

func (t *taskAdapter) GetAllTasks(
	ctx context.Context,
	getAllTasksRequest *grpc.GetAllTasksRequest,
) (*grpc.TaskDomainList, error) {
	ctx = context.WithValue(ctx, RequestId, getAllTasksRequest.RequestId)
	taskModelList, err := t.Task.GetAllTasks(ctx, getAllTasksRequest.UserId)
	if err != nil {
		return nil, exception.BuildError(err)
	}

	taskDomainList := grpc.TaskDomainList{}

	for _, taskModel := range *taskModelList {
		taskDomainList.TaskDomain = append(
			taskDomainList.TaskDomain,
			converters.ConvertToGrpcTaskDomain(taskModel),
		)
	}

	return &taskDomainList, nil
}

func (t *taskAdapter) CreateTask(
	ctx context.Context,
	createTaskRequest *grpc.CreateTaskRequest,
) (*grpc.TaskDomain, error) {
	ctx = context.WithValue(ctx, RequestId, createTaskRequest.RequestId)
	taskDomain, err := t.Task.CreateTask(ctx, createTaskRequest.TaskMessage)
	if err != nil {
		return nil, exception.BuildError(err)
	}

	return converters.ConvertToGrpcTaskDomain(*taskDomain), nil
}

func (t *taskAdapter) UpdateTaskMessage(
	ctx context.Context,
	updateTaskMessageRequest *grpc.UpdateTaskMessageRequest,
) (*grpc.TaskDomain, error) {
	ctx = context.WithValue(ctx, RequestId, updateTaskMessageRequest.RequestId)
	taskDomain, err := t.Task.UpdateTaskMessage(
		ctx,
		updateTaskMessageRequest.UserId,
		updateTaskMessageRequest.TaskId,
		updateTaskMessageRequest.TaskMessage,
	)
	if err != nil {
		return nil, exception.BuildError(err)
	}
	return converters.ConvertToGrpcTaskDomain(*taskDomain), nil
}

func (t *taskAdapter) UpdateTaskCompleteness(
	ctx context.Context,
	updateTaskCompletenessRequest *grpc.UpdateTaskCompletenessRequest,
) (*grpc.Void, error) {
	ctx = context.WithValue(ctx, RequestId, updateTaskCompletenessRequest.RequestId)
	err := t.Task.UpdateTaskCompleteness(ctx, updateTaskCompletenessRequest.UserId, updateTaskCompletenessRequest.TaskId)
	if err != nil {
		return nil, exception.BuildError(err)
	}
	return &grpc.Void{}, nil
}
