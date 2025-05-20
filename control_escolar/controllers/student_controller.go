package controllers

import (
	"net/http"
	"strconv"

	"control_escolar/services"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	service *services.StudentService
}

func NewStudentController(s *services.StudentService) *StudentController {
	return &StudentController{service: s}
}

func (ctrl *StudentController) GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, ctrl.service.GetAll())
}

func (ctrl *StudentController) GetStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("student_id"))
	student, ok := ctrl.service.GetByID(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (ctrl *StudentController) CreateStudent(c *gin.Context) {
	var s services.Student
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	student := ctrl.service.Create(s)
	c.JSON(http.StatusOK, student)
}

func (ctrl *StudentController) UpdateStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("student_id"))
	var s services.Student
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	student, err := ctrl.service.Update(id, s)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (ctrl *StudentController) DeleteStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("student_id"))
	if err := ctrl.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
