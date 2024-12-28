package ginhandler

import (
	"lesson22/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	studentRepo *repository.StudentRepository
}

func NewGinHandler(
	studentRepo *repository.StudentRepository,
) *GinHandler {
	return &GinHandler{
		studentRepo: studentRepo,
	}
}

func RunWithGin(handler *GinHandler) *http.Server {
	r := gin.Default()

	studentApi := r.Group("/student")
	{
		studentApi.POST("/", handler.CreateStudent)
		studentApi.GET(":id", handler.GetStudent)
		studentApi.PUT(":id", handler.UpdateStudent)
		studentApi.DELETE(":id", handler.DeleteStudent)
		studentApi.GET("getAll", handler.GetAllStudents)
	}

	server := &http.Server{Addr: ":8080", Handler: r}

	return server
}
