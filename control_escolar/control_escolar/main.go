package main

import (
	"control_escolar/controllers"
	"control_escolar/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(sqlite.Open("control_escolar.db?_busy_timeout=5000&_journal_mode=WAL"), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info), 
		TranslateError:         true,
		SkipDefaultTransaction: false,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Configurar el pool de conexiones
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database instance")
	}

	// Configurar límites de conexiones
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Migrar modelos
	db.AutoMigrate(&services.Student{}, &services.Subject{}, &services.Grade{}, &services.StudentSubject{})

	// Servicios
	studentService := services.NewStudentServiceDB(db)
	subjectService := services.NewSubjectServiceDB(db)
	gradeService := services.NewGradeServiceDB(db)

	// Controladores
	studentController := controllers.NewStudentController(studentService)
	subjectController := controllers.NewSubjectController(subjectService)
	gradeController := controllers.NewGradeController(gradeService, studentService)

	r := gin.Default()
	r.LoadHTMLGlob("public/*.html")

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
		api.GET("/grades/:grade_id", gradeController.GetGrade)
		api.GET("/grades/:grade_id/student/:student_id", gradeController.GetGradeByIDAndStudent)
		api.GET("/grades/student/:student_id", gradeController.GetGradesByStudent)

		// Nuevas rutas para inscripción
		api.POST("/students/:student_id/enroll/:subject_id", studentController.EnrollStudent)
		api.DELETE("/students/:student_id/enroll/:subject_id", studentController.UnenrollStudent)

		// Materias inscritas de un estudiante
		api.GET("/students/:student_id/subjects", studentController.GetEnrolledSubjects)
	}

	r.Static("/public", "./public")
	fmt.Println("Servidor escuchando en el puerto 3000")
	r.Run(":3000")
}
