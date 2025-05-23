package controllers

import (
	"net/http"
	"strconv"

	"control_escolar/services"

	"github.com/gin-gonic/gin"
)

type GradeController struct {
	service        services.GradeServiceInterface
	studentService services.StudentServiceInterface
}

func NewGradeController(s services.GradeServiceInterface, studentService services.StudentServiceInterface) *GradeController {
	return &GradeController{
		service:        s,
		studentService: studentService,
	}
}

func (ctrl *GradeController) CreateGrade(c *gin.Context) {
	var g services.Grade
	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	grade, err := ctrl.service.Create(g)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grade)
}

func (ctrl *GradeController) UpdateGrade(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("grade_id"))
	var g services.Grade
	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	grade, err := ctrl.service.Update(id, g)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grade)
}

func (ctrl *GradeController) DeleteGrade(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("grade_id"))
	if err := ctrl.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (ctrl *GradeController) GetGradeByIDAndStudent(c *gin.Context) {
	gradeID, _ := strconv.Atoi(c.Param("grade_id"))
	studentID, _ := strconv.Atoi(c.Param("student_id"))
	grade, ok := ctrl.service.GetByIDAndStudent(gradeID, studentID)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grade not found"})
		return
	}
	c.JSON(http.StatusOK, grade)
}

func (ctrl *GradeController) GetGradesByStudent(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("student_id"))
	if err != nil || studentID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	// Verificar primero si el estudiante existe
	_, exists := ctrl.studentService.GetByID(studentID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	grades := ctrl.service.GetByStudent(studentID)
	if len(grades) == 0 {
		c.JSON(http.StatusOK, []services.Grade{}) // Retornar array vacío en lugar de null
		return
	}
	c.JSON(http.StatusOK, grades)
}

func (ctrl *GradeController) GetGrade(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("grade_id"))
	grade, ok := ctrl.service.GetByID(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grade not found"})
		return
	}
	c.JSON(http.StatusOK, grade)
}
