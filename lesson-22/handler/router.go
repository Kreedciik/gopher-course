package handler

import (
	"lesson22/repository"
	"net/http"
)

type Handler struct {
	courseRepo  *repository.CourseRepository
	studentRepo *repository.StudentRepository
	tutorRepo   *repository.TutorRepository
	groupRepo   *repository.GroupRepository
}

func NewHandler(
	courseRepo *repository.CourseRepository,
	studentRepo *repository.StudentRepository,
	tutorRepo *repository.TutorRepository,
	groupRepo *repository.GroupRepository,
) *Handler {
	return &Handler{
		courseRepo:  courseRepo,
		studentRepo: studentRepo,
		tutorRepo:   tutorRepo,
		groupRepo:   groupRepo,
	}
}

func Run(handler *Handler) *http.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /student", handler.CreateStudent)
	mux.HandleFunc("GET /student/{id}", handler.GetStudent)
	mux.HandleFunc("PUT /student/{id}", handler.UpdateStudent)
	mux.HandleFunc("DELETE /student/{id}", handler.DeleteStudent)
	mux.HandleFunc("GET /student/getAll", handler.GetAllStudents)

	mux.HandleFunc("POST /course", handler.CreateCourse)
	mux.HandleFunc("GET /course/{id}", handler.GetCourse)
	mux.HandleFunc("GET /course/getAll", handler.GetAllCourses)
	mux.HandleFunc("PUT /course", handler.UpdateCourse)
	mux.HandleFunc("DELETE /course/{id}", handler.DeleteCourse)

	mux.HandleFunc("POST /tutor", handler.CreateTutor)
	mux.HandleFunc("GET /tutor/{id}", handler.GetTutor)
	mux.HandleFunc("PUT /tutor", handler.UpdateTutor)
	mux.HandleFunc("DELETE /tutor/{id}", handler.DeleteTutor)
	mux.HandleFunc("GET /tutor/getAll", handler.GetAllTutors)

	mux.HandleFunc("POST /group", handler.CreateGroup)
	mux.HandleFunc("GET /group/{id}", handler.GetGroup)
	mux.HandleFunc("PUT /group", handler.UpdateGroup)
	mux.HandleFunc("DELETE /group/{id}", handler.DeleteGroup)
	mux.HandleFunc("GET /group/getAll", handler.GetAllGroups)

	server := &http.Server{Addr: ":8080", Handler: mux}

	return server
}
