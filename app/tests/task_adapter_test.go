package tests

import (
	"context"
	"database/sql/driver"
	"fmt"
	"os"
	"testing"
	"time"
	pb "to-do-list-server/app/config/grpc"
	mock "to-do-list-server/app/tests/mock"

	goSqlMock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func TestGetAllTasks_200(t *testing.T) {
	setEnvVars()
	defer deleteEnvVars()
	assert := assert.New(t)
	ctx := context.Background()
	client, sqlMock, closer := mock.StartServer(ctx, t)
	defer closer()

	sqlMock.ExpectQuery(
		"SELECT * FROM `TASK` WHERE UserId = ?",
	).WithArgs(mock.UserId).WillReturnRows(
		sqlMock.NewRows([]string{"TaskId", "TaskMessage", "CreatedAt", "IsTaskCompleted", "UserId"}).
			AddRows(
				[]driver.Value{mock.TaskId, mock.TaskMessage, time.Now(), true, mock.UserId},
				[]driver.Value{mock.TaskId, mock.TaskMessage, time.Now(), true, mock.UserId},
			),
	)

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
	assert.Equal(2, len(taskDomainSlice.TaskDomain))
	assert.Equal(mock.TaskId, taskDomainSlice.TaskDomain[0].TaskId)
}

func TestCreateTask_200(t *testing.T) {
	setEnvVars()
	defer deleteEnvVars()
	assert := assert.New(t)
	ctx := context.Background()
	client, sqlMock, closer := mock.StartServer(ctx, t)
	defer closer()

	sqlMock.ExpectQuery(
		"SELECT * FROM `ACCOUNT` WHERE `ACCOUNT`.`UserId` = ? ORDER BY `ACCOUNT`.`UserId` LIMIT ?",
	).WithArgs(mock.UserId, 1).WillReturnRows(
		sqlMock.NewRows([]string{"UserId", "Username", "Password"}).
			AddRow(mock.UserId, mock.Username, mock.Password),
	)
	sqlMock.ExpectBegin()

	sqlMock.ExpectExec(
		"INSERT INTO `TASK` (`TaskId`,`TaskMessage`,`CreatedAt`,`IsTaskCompleted`,`UserId`) VALUES (?,?,?,?,?)",
	).WithArgs(mock.AnyString{}, mock.TaskMessage, mock.AnyTime{}, false, mock.UserId).WillReturnResult(
		goSqlMock.NewResult(1, 1),
	)

	sqlMock.ExpectCommit()

	request := &pb.CreateTaskRequest{
		UserId:      mock.UserId,
		TaskMessage: mock.TaskMessage,
		RequestId:   uuid.New().String(),
	}

	taskDomain, err := client.CreateTask(ctx, request)
	if err != nil {
		t.Fatalf("error testing CreateTask: %v", err)
	} else if taskDomain == nil {
		t.Fatalf("unexpected nil response testing CreateTask: %v", err)
	}
	assert.Equal(mock.UserId, taskDomain.UserId)
}

func TestUpdateTaskMessage_200(t *testing.T) {
	setEnvVars()
	defer deleteEnvVars()
	assert := assert.New(t)
	ctx := context.Background()
	client, sqlMock, closer := mock.StartServer(ctx, t)
	defer closer()

	sqlMock.ExpectQuery(
		"SELECT * FROM `TASK` WHERE `TASK`.`TaskId` = ? ORDER BY `TASK`.`TaskId` LIMIT ?",
	).WithArgs(mock.TaskId, 1).WillReturnRows(
		sqlMock.NewRows([]string{"TaskId", "TaskMessage", "CreatedAt", "IsTaskCompleted", "UserId"}).
			AddRow(mock.TaskId, mock.TaskMessage, time.Now(), true, mock.UserId),
	)

	sqlMock.ExpectBegin()

	sqlMock.ExpectExec(
		"UPDATE `TASK` SET `TaskMessage`=? WHERE TaskId = ?",
	).WithArgs(
		fmt.Sprintf("NEW %s", mock.TaskMessage),
		mock.TaskId,
	).WillReturnResult(
		goSqlMock.NewResult(1, 1),
	)

	sqlMock.ExpectCommit()

	request := &pb.UpdateTaskMessageRequest{
		TaskId:      mock.TaskId,
		TaskMessage: fmt.Sprintf("NEW %s", mock.TaskMessage),
		RequestId:   uuid.New().String(),
	}

	taskDomain, err := client.UpdateTaskMessage(ctx, request)
	if err != nil {
		t.Fatalf("error testing UpdateTaskMessage: %v", err)
	} else if taskDomain == nil {
		t.Fatalf("unexpected nil response testing UpdateTaskMessage: %v", err)
	}
	assert.Equal(mock.UserId, taskDomain.UserId)
	assert.Equal(fmt.Sprintf("NEW %s", mock.TaskMessage), taskDomain.TaskMessage)
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
