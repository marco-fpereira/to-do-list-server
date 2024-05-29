package tests

import (
	"context"
	"os"
	"testing"
	pb "to-do-list-server/app/config/grpc"
	mock "to-do-list-server/app/tests/mock"

	"github.com/google/uuid"
)

func TestGetAllTasks_200(t *testing.T) {
	setEnvVars()
	defer deleteEnvVars()
	ctx := context.Background()
	client, closer := mock.StartServer(ctx, t)

	defer closer()
	request := &pb.GetAllTasksRequest{
		UserId:    mock.UserId,
		RequestId: uuid.New().String(),
	}

	taskDomainSlice, err := client.GetAllTasks(ctx, request)
	if err != nil {
		t.Fatalf("error testing GetAllTasks: %v", err)
	} else if taskDomainSlice == nil {
		t.Fatalf("unexpected nil response testing GetAllTasks: %v", err)
	}

	if len(taskDomainSlice.TaskDomain) != 1 || taskDomainSlice.TaskDomain[0].TaskId != mock.TaskId {
		t.Fatalf("unexpected list size response testing GetAllTasks: %d", len(taskDomainSlice.TaskDomain))
	} else {
		t.Log("success")
	}
}

func setEnvVars() {
	os.Setenv("HOST", "localhost")
	os.Setenv("DBPORT", "3306")
	os.Setenv("DBUSER", "root")
	os.Setenv("DBPASS", "root")
	os.Setenv("DBNAME", "TODOLIST")
}

func deleteEnvVars() {
	os.Unsetenv("HOST")
	os.Unsetenv("DBPORT")
	os.Unsetenv("DBUSER")
	os.Unsetenv("DBPASS")
	os.Unsetenv("DBNAME")
}
