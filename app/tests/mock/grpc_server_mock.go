package mock

import (
	"context"
	"database/sql/driver"
	"log"
	"math/rand"
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
	Username    = "test_user"
	Password    = "My_s@f3_password"
	TaskId      = "0fa0ac05-f20f-446f-beca-20fa636daf9c"
	TaskMessage = "Hello World"
)

func StartServer(ctx context.Context, t *testing.T) (pb.TaskClient, sqlmock.Sqlmock, func()) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	buffer := 101024 * 1024
	listener := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()

	dbMock, sqlMock := getSqlMock()

	database := output.NewMysqlDatabaseAdapter(dbMock)
	auth := output.NewJwtAuthenticationAdapter()
	task := usecase.NewTaskUseCase(auth, database)
	taskAdapter := adapterInput.NewTaskAdapter(task)

	pb.RegisterTaskServer(baseServer, taskAdapter)

	go func() {
		if err := baseServer.Serve(listener); err != nil {
			log.Fatalf("Error serving server %v", err)
		}
	}()

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("error connecting to server %v", err)
	}

	closer := func() {
		err := listener.Close()
		if err != nil {
			log.Printf("error closing listener %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewTaskClient(conn)
	return client, sqlMock, closer
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
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type AnyString struct{}

func (a AnyString) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

func GenerateRandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
