package main

import (
	"fmt"
	"net"

	inputAdapter "to-do-list-server/app/adapters/input"
	outputAdapter "to-do-list-server/app/adapters/output"
	"to-do-list-server/app/config"
	pb "to-do-list-server/app/config/grpc"
	outputDomain "to-do-list-server/app/domain/port/output"
	domain "to-do-list-server/app/domain/usecase"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	config.InitLog()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	var db *gorm.DB
	db, err = config.DbConnect()
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}

	jwtAuthenticationAdapter := outputAdapter.NewJwtAuthenticationAdapter()
	mysqlDatabaseAdapter := outputAdapter.NewMysqlDatabaseAdapter(db)
	pb.RegisterTaskServer(grpcServer, buildTaskAdapter(jwtAuthenticationAdapter, mysqlDatabaseAdapter))
	pb.RegisterAccountServer(grpcServer, buildAccountAdapter(jwtAuthenticationAdapter, mysqlDatabaseAdapter))

	grpcServer.Serve(listener)
}

func buildAccountAdapter(
	jwtAuthenticationAdapter outputDomain.AuthenticationPort,
	mysqlDatabaseAdapter outputDomain.DatabasePort,
) pb.AccountServer {
	accountPort := domain.NewAccountUseCase(jwtAuthenticationAdapter, mysqlDatabaseAdapter)
	return inputAdapter.NewAccountAdapter(accountPort)
}

func buildTaskAdapter(
	jwtAuthenticationAdapter outputDomain.AuthenticationPort,
	mysqlDatabaseAdapter outputDomain.DatabasePort,
) pb.TaskServer {
	taskPort := domain.NewTaskUseCase(jwtAuthenticationAdapter, mysqlDatabaseAdapter)
	return inputAdapter.NewTaskAdapter(taskPort)
}
