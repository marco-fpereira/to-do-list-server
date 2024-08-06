package main

import (
	"context"
	"fmt"
	"net"

	inputAdapter "to-do-list-server/app/adapters/input"
	outputAdapter "to-do-list-server/app/adapters/output"
	"to-do-list-server/app/config"
	pb "to-do-list-server/app/config/grpc"
	"to-do-list-server/app/config/logger"
	outputDomain "to-do-list-server/app/domain/port/output"
	domain "to-do-list-server/app/domain/usecase"

	"github.com/joho/godotenv"
	grpc "google.golang.org/grpc"
	"gorm.io/gorm"
)

var grpcPort = 50051

func main() {
	godotenv.Load()
	logger.InitLogger()
	ctx := context.Background()

	logger.Info(ctx, "Starting application")

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", grpcPort))
	if err != nil {
		logger.Fatal(context.Background(), "failed to listen", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	var db *gorm.DB
	db, err = config.DbConnect()
	if err != nil {
		logger.Fatal(context.Background(), "failed to open database connection", err)
	}
	logger.Info(ctx, "Database connection successfully created")

	bcryptCryptographyAdapter := outputAdapter.NewBCryptCryptographyAdapter()
	jwtAuthenticationAdapter := outputAdapter.NewJwtAuthenticationAdapter()
	mysqlDatabaseAdapter := outputAdapter.NewMysqlDatabaseAdapter(db)
	pb.RegisterTaskServer(grpcServer, buildTaskAdapter(jwtAuthenticationAdapter, mysqlDatabaseAdapter))
	pb.RegisterAccountServer(grpcServer,
		buildAccountAdapter(jwtAuthenticationAdapter, bcryptCryptographyAdapter, mysqlDatabaseAdapter),
	)
	logger.Info(context.Background(), fmt.Sprintf("application started on port %d", grpcPort))
	grpcServer.Serve(listener)
}

func buildAccountAdapter(
	jwtAuthenticationAdapter outputDomain.AuthenticationPort,
	bcryptCryptographyPort outputDomain.CryptographyPort,
	mysqlDatabaseAdapter outputDomain.DatabasePort,
) pb.AccountServer {
	accountPort := domain.NewAccountUseCase(
		jwtAuthenticationAdapter,
		bcryptCryptographyPort,
		mysqlDatabaseAdapter,
	)
	return inputAdapter.NewAccountAdapter(accountPort)
}

func buildTaskAdapter(
	jwtAuthenticationAdapter outputDomain.AuthenticationPort,
	mysqlDatabaseAdapter outputDomain.DatabasePort,
) pb.TaskServer {
	taskPort := domain.NewTaskUseCase(jwtAuthenticationAdapter, mysqlDatabaseAdapter)
	return inputAdapter.NewTaskAdapter(taskPort)
}
