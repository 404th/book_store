package main

import (
	"fmt"
	"net"

	"github.com/404th/book_store/config"
	"github.com/404th/book_store/grpc"
	"github.com/404th/book_store/pkg/logger"
	"github.com/404th/book_store/storage/postgres"
)

func main() {
	// loads config file
	cfg := config.Load()

	// defines logger level
	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	pgStore, err := postgres.NewPostgres(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	), cfg)
	if err != nil {
		panic(err)
	}

	// gRPC
	grpcServer := grpc.SetUpServer(cfg, log, pgStore)

	// listening net
	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		panic(err)
	}

	log.Info("gRPC: get started...", logger.String("port", cfg.GRPCPort))

	// serving
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
