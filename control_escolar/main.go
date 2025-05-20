package main

import (
	"control_escolar/controllers"
	"control_escolar/middleware"
	"control_escolar/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	r.LoadHTMLGlob("public/*.html")

	// Servicios
	studentService := services.NewStudentService()
	subjectService := services.NewSubjectService()
	gradeService := services.NewGradeService(studentService, subjectService)

	// Controladores
	studentController := controllers.NewStudentController(studentService)
	subjectController := controllers.NewSubjectController(subjectService)
	gradeController := controllers.NewGradeController(gradeService)

	// UI
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	api := r.Group("/api")
	{
		// Students
		api.GET("/students", studentController.GetStudents)
		api.GET("/students/:student_id", studentController.GetStudent)
		api.POST("/students", studentController.CreateStudent)
		api.PUT("/students/:student_id", studentController.UpdateStudent)
		api.DELETE("/students/:student_id", studentController.DeleteStudent)

		// Subjects
		api.GET("/subjects", subjectController.GetSubjects)
		api.GET("/subjects/:subject_id", subjectController.GetSubject)
		api.POST("/subjects", subjectController.CreateSubject)
		api.PUT("/subjects/:subject_id", subjectController.UpdateSubject)
		api.DELETE("/subjects/:subject_id", subjectController.DeleteSubject)

		// Grades
		api.POST("/grades", gradeController.CreateGrade)
		api.PUT("/grades/:grade_id", gradeController.UpdateGrade)
		api.DELETE("/grades/:grade_id", gradeController.DeleteGrade)
		api.GET("/grades/:grade_id/student/:student_id", gradeController.GetGradeByIDAndStudent)
		api.GET("/grades/student/:student_id", gradeController.GetGradesByStudent)
	}

	r.Static("/public", "./public")
	fmt.Println("Servidor escuchando en el puerto 3000")
	r.Run(":3000")
}
