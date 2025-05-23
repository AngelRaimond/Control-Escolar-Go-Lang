package controllers

import (
	"net/http"
	"strconv"

	"control_escolar/services"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	service services.StudentServiceInterface
}

func NewStudentController(s services.StudentServiceInterface) *StudentController {
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

func (ctrl *StudentController) EnrollStudent(c *gin.Context) {
	studentID, _ := strconv.Atoi(c.Param("student_id"))
	subjectID, _ := strconv.Atoi(c.Param("subject_id"))

	err := ctrl.service.EnrollInSubject(studentID, subjectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (ctrl *StudentController) UnenrollStudent(c *gin.Context) {
	studentID, _ := strconv.Atoi(c.Param("student_id"))
	subjectID, _ := strconv.Atoi(c.Param("subject_id"))

	err := ctrl.service.UnenrollFromSubject(studentID, subjectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (ctrl *StudentController) GetEnrolledSubjects(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("student_id"))
	student, ok := ctrl.service.GetByID(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student.Subjects)
}
