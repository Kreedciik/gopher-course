package main

import (
	"blogpost"
	"blogpost/pkg/handler"
	"blogpost/pkg/repository"
	"blogpost/pkg/repository/cache"
	"blogpost/pkg/repository/mongodb"
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

	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		panic("You must set your 'MONGODB_URI' environment variable")
	}

	pgDB, err := postgres.NewPostgres(postgres.Config{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		Dbname:   os.Getenv("PG_DATABASE"),
		SSLMode:  os.Getenv("PG_SSL_MODE"),
	})

	rdb := cache.NewRedisCache(cache.RedisConfig{
		Address: os.Getenv("REDIS_ADDRESS"),
	})

	mongoClient := mongodb.NewMongoDB(mongoUri)

	if err != nil {
		panic(err)
	}

	repository := repository.NewRepository(pgDB, mongoClient.Database("some_database"))
	services := service.NewServices(repository, rdb)
	handler := handler.NewHandler(services)
	server := new(blogpost.Server)
	err = server.Run(":8000", handler.InitRoutes())

	if err != nil {
		panic(err)
	}
}
