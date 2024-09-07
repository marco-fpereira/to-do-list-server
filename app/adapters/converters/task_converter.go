package converters

import (
	"github.com/marco-fpereira/to-do-list-server/config/grpc"
	"github.com/marco-fpereira/to-do-list-server/domain/model"
)

func ConvertToGrpcTaskDomain(t model.TaskDomain) *grpc.TaskDomain {
	return &grpc.TaskDomain{
		TaskId:          t.TaskId,
		TaskMessage:     t.TaskMessage,
		CreatedAt:       t.CreatedAt.String(),
		IsTaskCompleted: t.IsTaskCompleted,
		UserId:          t.UserId,
	}
}
