package input

import (
	"context"
	"to-do-list-server/app/adapters/converters"
	"to-do-list-server/app/config/grpc"
	"to-do-list-server/app/domain/port/input"
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
	taskModelList, err := t.Task.GetAllTasks(ctx, getAllTasksRequest.UserId)
	if err != nil {
		return nil, err
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

func (t *taskAdapter) CreateTask(context.Context, *grpc.CreateTaskRequest) (*grpc.TaskDomain, error) {
	return nil, nil
}

func (t *taskAdapter) UpdateTaskMessage(context.Context, *grpc.UpdateTaskMessageRequest) (*grpc.TaskDomain, error) {
	return nil, nil
}

func (t *taskAdapter) UpdateTaskCompleteness(context.Context, *grpc.UpdateTaskCompletenessRequest) (*grpc.Void, error) {
	return nil, nil
}
