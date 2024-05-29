package tests

import (
	"context"
	"log"
	"net"
	"testing"
	"time"
	adapterInput "to-do-list-server/app/adapters/input"
	"to-do-list-server/app/adapters/output"
	pb "to-do-list-server/app/config/grpc"
	"to-do-list-server/app/domain/usecase"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	UserId      = "7f937a3d-03fe-445b-8df8-7d8a43c25a52"
	TaskId      = "0fa0ac05-f20f-446f-beca-20fa636daf9c"
	TaskMessage = "Hello World"

// username = "test_user"
// password = "My_s@f3_password"
)

func StartServer(ctx context.Context, t *testing.T) (pb.TaskClient, func()) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	buffer := 101024 * 1024
	listener := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()

	dbMock, _ := getSqlMock()

	database := output.NewMysqlDatabaseAdapter(dbMock)
	task := usecase.NewTaskUseCase(database)
	taskAdapter := adapterInput.NewTaskAdapter(task)

	pb.RegisterTaskServer(baseServer, taskAdapter)

	go func() {
		if err := baseServer.Serve(listener); err != nil {
			log.Printf("Error serving server %v", err)
		}
	}()

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Printf("error connecting to server %v", err)
	}

	closer := func() {
		err := listener.Close()
		if err != nil {
			log.Printf("error closing listener %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewTaskClient(conn)
	return client, closer
}

func getSqlMock() (*gorm.DB, sqlmock.Sqlmock) {
	mockDb, sqlMock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	dialector := mysql.New(mysql.Config{
		DSN:        "sqlmock_db_0",
		Conn:       mockDb,
		DriverName: "mysql",
	})
	mockSqlCalls(sqlMock)

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("Error initializing database. Details %v", err)
	}

	return db, sqlMock
}

func mockSqlCalls(sqlMock sqlmock.Sqlmock) {
	sqlMock.ExpectQuery("SELECT VERSION()").
		WillReturnRows(sqlmock.NewRows([]string{"version()"}).AddRow("8.0.23"))

	sqlMock.ExpectQuery("SELECT * FROM `TASK` WHERE `TASK`.`user_id` = ?").WithArgs(UserId).WillReturnRows(
		sqlMock.NewRows([]string{"TaskId", "TaskMessage", "CreatedAt", "IsTaskCompleted", "UserId"}).
			AddRow(TaskId, TaskMessage, time.Now(), true, UserId),
	)
}
