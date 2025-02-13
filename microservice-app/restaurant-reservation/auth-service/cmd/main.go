package main

import (
	"auth/config"
	"auth/pkg/controller"
	"auth/pkg/repository"
	"auth/pkg/repository/postgres"
	"auth/pkg/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(".env"); err != nil {
		slog.Error(fmt.Sprintf("could not read .env file %s", err.Error()))
	}

	if err := config.InitConfig(); err != nil {
		slog.Error(err.Error())
	}

	postgresDB, err := postgres.NewPostgres(postgres.PostgresConfig{
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetString("postgres.port"),
		User:     viper.GetString("postgres.user"),
		Password: os.Getenv("PG_PASSWORD"),
		Dbname:   viper.GetString("postgres.dbName"),
		SSLMode:  viper.GetString("postgres.sslMode"),
	})

	if err != nil {
		slog.Error(fmt.Sprintf("postgres: %s", err.Error()))
	}
	repositories := repository.NewRepository(postgresDB)
	services := service.NewService(repositories)
	contr := controller.NewHandler(services)

	l, err := net.Listen("tcp", ":8545") // Standard Ethereum JSON-RPC port
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()
	server := grpc.NewServer()
	contr.InitServers(server)

	if err := server.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
