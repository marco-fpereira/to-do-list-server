package converters

import (
	"to-do-list-server/app/config/grpc"
	"to-do-list-server/app/domain/model"
)

func ConvertToGrpcTaskDomain(t model.TaskDomain) *grpc.TaskDomain {
	return &grpc.TaskDomain{
		TaskId:          t.TaskId,
		TaskMessage:     t.TaskMessage,
		CreatedAt:       t.CreatedAt.String(),
		IsTaskCompleted: t.IsTaskCompleted,
	}
}
