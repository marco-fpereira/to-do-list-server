package input

import (
	"context"

	consts "github.com/marco-fpereira/to-do-list-server/adapters/consts"
	"github.com/marco-fpereira/to-do-list-server/adapters/converters"
	"github.com/marco-fpereira/to-do-list-server/adapters/exception/handler"
	"github.com/marco-fpereira/to-do-list-server/config/grpc"
	"github.com/marco-fpereira/to-do-list-server/config/logger"
	"github.com/marco-fpereira/to-do-list-server/domain/port/input"

	"go.uber.org/zap"
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
	ctx = context.WithValue(ctx, consts.REQUEST_ID, getAllTasksRequest.RequestId)
	tags := []zap.Field{
		zap.String("UserId", getAllTasksRequest.UserId),
		zap.String("RequestId", getAllTasksRequest.RequestId),
	}

	logger.Info(ctx, "Getting all tasks for the given UserId.", tags...)

	taskModelList, err := t.Task.GetAllTasks(
		ctx,
		getAllTasksRequest.UserId,
		getAllTasksRequest.Token,
	)
	if err != nil {
		logger.Error(ctx, "Error getting all tasks.", err)
		return nil, handler.HandleException(err)
	}

	taskDomainList := grpc.TaskDomainList{}

	for _, taskModel := range *taskModelList {
		taskDomainList.TaskDomain = append(
			taskDomainList.TaskDomain,
			converters.ConvertToGrpcTaskDomain(taskModel),
		)
	}
	tags = append(tags, zap.Int("NumberOfTasks", len(taskDomainList.TaskDomain)))
	logger.Info(ctx, "Tasks successfully retrieved.", tags...)

	return &taskDomainList, nil
}

func (t *taskAdapter) CreateTask(
	ctx context.Context,
	createTaskRequest *grpc.CreateTaskRequest,
) (*grpc.TaskDomain, error) {
	ctx = context.WithValue(ctx, consts.REQUEST_ID, createTaskRequest.RequestId)
	tags := []zap.Field{
		zap.String("UserId", createTaskRequest.UserId),
		zap.String("RequestId", createTaskRequest.RequestId),
	}

	logger.Info(ctx, "Creating new task.", tags...)

	taskDomain, err := t.Task.CreateTask(
		ctx,
		createTaskRequest.UserId,
		createTaskRequest.TaskMessage,
		createTaskRequest.Token,
	)
	if err != nil {
		logger.Error(ctx, "Error creating new task.", err, tags...)
		return nil, handler.HandleException(err)
	}

	tags = append(tags, zap.String("TaskId", taskDomain.TaskId))
	logger.Info(ctx, "Task successfully created.", tags...)
	return converters.ConvertToGrpcTaskDomain(*taskDomain), nil
}

func (t *taskAdapter) UpdateTaskMessage(
	ctx context.Context,
	updateTaskMessageRequest *grpc.UpdateTaskMessageRequest,
) (*grpc.TaskDomain, error) {
	ctx = context.WithValue(ctx, consts.REQUEST_ID, updateTaskMessageRequest.RequestId)
	tags := []zap.Field{
		zap.String("TaskID", updateTaskMessageRequest.TaskId),
		zap.String("RequestId", updateTaskMessageRequest.RequestId),
	}

	logger.Info(ctx, "Updating task message.", tags...)

	taskDomain, err := t.Task.UpdateTaskMessage(
		ctx,
		updateTaskMessageRequest.TaskId,
		updateTaskMessageRequest.TaskMessage,
		updateTaskMessageRequest.Token,
	)
	if err != nil {
		logger.Error(ctx, "Error updating task message.", err, tags...)
		return nil, handler.HandleException(err)
	}
	logger.Info(ctx, "Task message successfully updated.", tags...)
	return converters.ConvertToGrpcTaskDomain(*taskDomain), nil
}

func (t *taskAdapter) UpdateTaskCompleteness(
	ctx context.Context,
	updateTaskCompletenessRequest *grpc.UpdateTaskCompletenessRequest,
) (*grpc.Void, error) {
	ctx = context.WithValue(ctx, consts.REQUEST_ID, updateTaskCompletenessRequest.RequestId)
	tags := []zap.Field{
		zap.String("TaskID", updateTaskCompletenessRequest.TaskId),
		zap.String("RequestId", updateTaskCompletenessRequest.RequestId),
	}

	logger.Info(ctx, "Updating task completeness.", tags...)

	err := t.Task.UpdateTaskCompleteness(
		ctx,
		updateTaskCompletenessRequest.TaskId,
		updateTaskCompletenessRequest.Token,
	)
	if err != nil {
		logger.Error(ctx, "Error updating task completeness.", err, tags...)
		return nil, handler.HandleException(err)
	}
	logger.Info(ctx, "Task completeness successfully updated.", tags...)
	return &grpc.Void{}, nil
}

func (t *taskAdapter) DeleteMessage(
	ctx context.Context,
	deleteMessageRequest *grpc.DeleteMessageRequest,
) (*grpc.Void, error) {
	ctx = context.WithValue(ctx, consts.REQUEST_ID, deleteMessageRequest.RequestId)
	tags := []zap.Field{
		zap.String("TaskID", deleteMessageRequest.TaskId),
		zap.String("RequestId", deleteMessageRequest.RequestId),
	}

	logger.Info(ctx, "Deleting task.", tags...)
	err := t.Task.DeleteTask(
		ctx,
		deleteMessageRequest.TaskId,
		deleteMessageRequest.Token,
	)
	if err != nil {
		logger.Error(ctx, "Error deleting task.", err, tags...)
		return nil, handler.HandleException(err)
	}
	logger.Info(ctx, "Task successfully deleted.", tags...)
	return &grpc.Void{}, nil
}
