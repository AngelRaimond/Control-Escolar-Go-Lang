package controllers

import (
	"net/http"
	"strconv"

	"control_escolar/services"

	"github.com/gin-gonic/gin"
)

type SubjectController struct {
	service services.SubjectServiceInterface
}

func NewSubjectController(s services.SubjectServiceInterface) *SubjectController {
	return &SubjectController{service: s}
}

func (ctrl *SubjectController) GetSubjects(c *gin.Context) {
	c.JSON(http.StatusOK, ctrl.service.GetAll())
}

func (ctrl *SubjectController) GetSubject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("subject_id"))
	subject, ok := ctrl.service.GetByID(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subject not found"})
		return
	}
	c.JSON(http.StatusOK, subject)
}

func (ctrl *SubjectController) CreateSubject(c *gin.Context) {
	var s services.Subject
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	subject := ctrl.service.Create(s)
	c.JSON(http.StatusOK, subject)
}

func (ctrl *SubjectController) UpdateSubject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("subject_id"))
	var s services.Subject
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	subject, err := ctrl.service.Update(id, s)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subject)
}

func (ctrl *SubjectController) DeleteSubject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("subject_id"))
	if err := ctrl.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
