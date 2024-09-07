package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	inputAdapter "github.com/marco-fpereira/to-do-list-server/adapters/input"
	outputAdapter "github.com/marco-fpereira/to-do-list-server/adapters/output"
	"github.com/marco-fpereira/to-do-list-server/config"
	"github.com/marco-fpereira/to-do-list-server/config/env"
	pb "github.com/marco-fpereira/to-do-list-server/config/grpc"
	"github.com/marco-fpereira/to-do-list-server/config/logger"
	outputDomain "github.com/marco-fpereira/to-do-list-server/domain/port/output"
	domain "github.com/marco-fpereira/to-do-list-server/domain/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	grpc "google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	grpcPort        = env.GetEnv("GRPC_PORT", "50051")
	healthCheckPort = env.GetEnv("HEALTH_CHECK_PORT", "8888")
)

func main() {
	godotenv.Load()
	logger.InitLogger()
	ctx := context.Background()

	logger.Info(ctx, "Starting application")

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", grpcPort))
	if err != nil {
		logger.Fatal(ctx, "failed to listen", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	var db *gorm.DB
	db, err = config.DbConnect()
	if err != nil {
		logger.Fatal(ctx, "failed to open database connection", err)
	}
	logger.Info(ctx, "Database connection successfully created")

	bcryptCryptographyAdapter := outputAdapter.NewBCryptCryptographyAdapter()
	jwtAuthenticationAdapter := outputAdapter.NewJwtAuthenticationAdapter()
	mysqlDatabaseAdapter := outputAdapter.NewMysqlDatabaseAdapter(db)
	pb.RegisterTaskServer(grpcServer, buildTaskAdapter(jwtAuthenticationAdapter, mysqlDatabaseAdapter))
	pb.RegisterAccountServer(grpcServer,
		buildAccountAdapter(jwtAuthenticationAdapter, bcryptCryptographyAdapter, mysqlDatabaseAdapter),
	)
	go setupHealthCheck(ctx)
	logger.Info(ctx, fmt.Sprintf("application started on port %s", grpcPort))
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

func setupHealthCheck(ctx context.Context) {
	r := gin.Default()
	r.GET("/actuator/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})
	err := r.Run(fmt.Sprintf(":%s", healthCheckPort))
	if err != nil {
		logger.Fatal(ctx, "error starting health check server. Details: %v", err)
	}
}
