package main

import (
	"blogpost"
	"blogpost/pkg/handler"
	"blogpost/pkg/repository"
	"blogpost/pkg/repository/postgres"
	"blogpost/pkg/service"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	pgDB, err := postgres.NewPostgres(postgres.Config{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		Dbname:   os.Getenv("PG_DATABASE"),
		SSLMode:  os.Getenv("PG_SSL_MODE"),
	})

	if err != nil {
		panic(err)
	}

	repository := repository.NewRepository(pgDB)
	services := service.NewServices(repository)
	handler := handler.NewHandler(services)
	server := new(blogpost.Server)
	err = server.Run(":8000", handler.InitRoutes())

	if err != nil {
		panic(err)
	}
}
