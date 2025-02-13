package main

import (
	"farmish/app"
	"farmish/config"
	"farmish/pkg/handler"
	"farmish/pkg/repository"
	"farmish/pkg/repository/postgres"
	"farmish/pkg/service"
	"fmt"
	"log/slog"
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
	h := handler.NewHandler(services)
	server := app.Server{AuthRPC: h}
	err = server.Run(viper.GetString("server.port"))

	if err != nil {
		slog.Error(fmt.Sprintf("error when initialize rpc server: %s", err.Error()))
	}
}
