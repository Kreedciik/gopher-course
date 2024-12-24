package main

import (
	"lesson22/handler"
	"lesson22/postgres"
	"lesson22/repository"
	"log"
)

func main() {
	pgDB, err := postgres.InitializePostgres()
	if err != nil {
		panic(err)
	}
	defer pgDB.Close()

	courseRepo := repository.CreateCourseRepository(pgDB)
	h := handler.NewHandler(&courseRepo)

	mux := handler.Run(h)
	err = mux.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
