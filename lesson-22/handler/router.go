package handler

import (
	"lesson22/repository"
	"net/http"
)

type Handler struct {
	courseRepo  *repository.CourseRepository
	studentRepo *repository.StudentRepository
}

func NewHandler(courseRepo *repository.CourseRepository, studentRepo *repository.StudentRepository) *Handler {
	return &Handler{
		courseRepo:  courseRepo,
		studentRepo: studentRepo,
	}
}

func Run(handler *Handler) *http.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /student", handler.CreateStudent)
	mux.HandleFunc("GET /student/{id}", handler.GetStudent)
	mux.HandleFunc("PUT /student/{id}", handler.UpdateStudent)
	mux.HandleFunc("PUT /student/{id}", handler.DeleteStudent)

	mux.HandleFunc("POST /course", handler.CreateCourse)
	mux.HandleFunc("GET /course/{id}", handler.GetCourse)
	mux.HandleFunc("PUT /course", handler.UpdateCourse)
	mux.HandleFunc("DELETE /course/{id}", handler.DeleteCourse)
	server := &http.Server{Addr: ":8080", Handler: mux}

	return server
}
