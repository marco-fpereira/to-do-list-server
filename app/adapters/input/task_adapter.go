package input

import (
	"context"
	consts "to-do-list-server/app/adapters/consts"
	"to-do-list-server/app/adapters/converters"
	"to-do-list-server/app/config/grpc"
	"to-do-list-server/app/domain/port/input"
	"to-do-list-server/app/exception"

	log "github.com/sirupsen/logrus"
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
	tags := log.Fields{
		"UserId":    getAllTasksRequest.UserId,
		"RequestId": getAllTasksRequest.RequestId,
	}

	log.WithFields(tags).Info("Getting all tasks for the given UserId.")

	taskModelList, err := t.Task.GetAllTasks(ctx, getAllTasksRequest.UserId)
	if err != nil {
		tags["error"] = err
		log.WithFields(tags).Error("Error getting all tasks.")
		return nil, exception.BuildResponseException(err)
	}

	taskDomainList := grpc.TaskDomainList{}

	for _, taskModel := range *taskModelList {
		taskDomainList.TaskDomain = append(
			taskDomainList.TaskDomain,
			converters.ConvertToGrpcTaskDomain(taskModel),
		)
	}
	tags["NumberOfTasks"] = len(taskDomainList.TaskDomain)
	log.WithFields(tags).Info("Tasks successfully retrieved.")

	return &taskDomainList, nil
}

func (t *taskAdapter) CreateTask(
	ctx context.Context,
	createTaskRequest *grpc.CreateTaskRequest,
) (*grpc.TaskDomain, error) {
	ctx = context.WithValue(ctx, consts.REQUEST_ID, createTaskRequest.RequestId)
	tags := log.Fields{
		"UserId":    createTaskRequest.UserId,
		"RequestId": createTaskRequest.RequestId,
	}
	log.WithFields(tags).Info("Creating new task.")
	taskDomain, err := t.Task.CreateTask(ctx, createTaskRequest.UserId, createTaskRequest.TaskMessage)
	if err != nil {
		tags["error"] = err
		log.WithFields(tags).Error("Error creating new task.")
		return nil, exception.BuildResponseException(err)
	}

	tags["TaskId"] = taskDomain.TaskId
	log.WithFields(tags).Info("Task successfully created.")
	return converters.ConvertToGrpcTaskDomain(*taskDomain), nil
}

func (t *taskAdapter) UpdateTaskMessage(
	ctx context.Context,
	updateTaskMessageRequest *grpc.UpdateTaskMessageRequest,
) (*grpc.TaskDomain, error) {
	ctx = context.WithValue(ctx, consts.REQUEST_ID, updateTaskMessageRequest.RequestId)
	tags := log.Fields{
		"TaskId":    updateTaskMessageRequest.TaskId,
		"RequestId": updateTaskMessageRequest.RequestId,
	}
	log.WithFields(tags).Info("Updating task message.")
	taskDomain, err := t.Task.UpdateTaskMessage(
		ctx,
		updateTaskMessageRequest.TaskId,
		updateTaskMessageRequest.TaskMessage,
	)
	if err != nil {
		tags["error"] = err
		log.WithFields(tags).Error("Error updating task message.")
		return nil, exception.BuildResponseException(err)
	}
	log.WithFields(tags).Info("Task message successfully updated.")
	return converters.ConvertToGrpcTaskDomain(*taskDomain), nil
}

func (t *taskAdapter) UpdateTaskCompleteness(
	ctx context.Context,
	updateTaskCompletenessRequest *grpc.UpdateTaskCompletenessRequest,
) (*grpc.Void, error) {
	ctx = context.WithValue(ctx, consts.REQUEST_ID, updateTaskCompletenessRequest.RequestId)
	tags := log.Fields{
		"TaskId":    updateTaskCompletenessRequest.TaskId,
		"RequestId": updateTaskCompletenessRequest.RequestId,
	}
	log.WithFields(tags).Info("Updating task completeness.")
	err := t.Task.UpdateTaskCompleteness(
		ctx,
		updateTaskCompletenessRequest.TaskId,
	)
	if err != nil {
		tags["error"] = err
		log.WithFields(tags).Error("Error updating task completeness.")
		return nil, exception.BuildResponseException(err)
	}
	log.WithFields(tags).Info("Task completeness successfully updated.")
	return &grpc.Void{}, nil
}

func (t *taskAdapter) DeleteMessage(
	ctx context.Context,
	deleteMessageRequest *grpc.DeleteMessageRequest,
) (*grpc.Void, error) {
	ctx = context.WithValue(ctx, consts.REQUEST_ID, deleteMessageRequest.RequestId)
	tags := log.Fields{
		"TaskId":    deleteMessageRequest.TaskId,
		"RequestId": deleteMessageRequest.RequestId,
	}
	log.WithFields(tags).Info("Deleting task.")
	err := t.Task.DeleteTask(
		ctx,
		deleteMessageRequest.TaskId,
	)
	if err != nil {
		tags["error"] = err
		log.WithFields(tags).Error("Error deleting task.")
		return nil, exception.BuildResponseException(err)
	}
	log.WithFields(tags).Info("Task successfully deleted.")
	return &grpc.Void{}, nil
}
