package handler

import (
	"lesson22/repository"
	"net/http"
)

type Handler struct {
	courseRepo *repository.CourseRepository
}

func NewHandler(courseRepo *repository.CourseRepository) *Handler {
	return &Handler{
		courseRepo: courseRepo,
	}
}

func Run(handler *Handler) *http.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /course", handler.CreateCourse)
	mux.HandleFunc("GET /course/{id}", handler.GetCourse)
	mux.HandleFunc("PUT /course", handler.UpdateCourse)
	mux.HandleFunc("DELETE /course/{id}", handler.DeleteCourse)
	server := &http.Server{Addr: ":8080", Handler: mux}

	return server
}
