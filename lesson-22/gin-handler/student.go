package ginhandler

import (
	"lesson22/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (g *GinHandler) CreateStudent(c *gin.Context) {
	var student model.CreateStudentRequest
	err := c.ShouldBindJSON(&student)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = g.studentRepo.CreateStudent(student)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully created!"})
}
func (g *GinHandler) GetStudent(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	student, err := g.studentRepo.GetStudent(id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			http.Error(c.Writer, err.Error(), http.StatusNotFound)
		} else {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}
func (g *GinHandler) GetAllStudents(c *gin.Context) {
	students, err := g.studentRepo.GetAllStudents()
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": students})
}
func (g *GinHandler) UpdateStudent(c *gin.Context) {
	var student model.Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = g.studentRepo.UpdateStudent(student)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully updated"})
}
func (g *GinHandler) DeleteStudent(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	err := g.studentRepo.DeleteStudent(id)

	if err != nil {
		if err.Error() == "not listed" {
			http.Error(c.Writer, err.Error(), http.StatusNotFound)
		} else {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}
