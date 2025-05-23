package services

import (
	"errors"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type GradeService struct {
	grades         []Grade
	nextID         int
	studentService *StudentService
	subjectService *SubjectService
}

type GradeServiceDB struct {
	db *gorm.DB
}

func NewGradeService(stSrv *StudentService, sbSrv *SubjectService) *GradeService {
	return &GradeService{nextID: 1, studentService: stSrv, subjectService: sbSrv}
}

func NewGradeServiceDB(db *gorm.DB) *GradeServiceDB {
	return &GradeServiceDB{db: db}
}

func (g *GradeService) Create(gr Grade) (Grade, error) {
	if _, ok := g.studentService.GetByID(gr.StudentID); !ok {
		return Grade{}, errors.New("student_id does not exist")
	}
	if _, ok := g.subjectService.GetByID(gr.SubjectID); !ok {
		return Grade{}, errors.New("subject_id does not exist")
	}
	gr.GradeID = g.nextID
	g.nextID++
	g.grades = append(g.grades, gr)
	return gr, nil
}

func (g *GradeServiceDB) Create(gr Grade) (Grade, error) {
	// Verificar que el estudiante esté inscrito en la materia
	var count int64
	g.db.Model(&StudentSubject{}).Where("student_id = ? AND subject_id = ?",
		gr.StudentID, gr.SubjectID).Count(&count)
	if count == 0 {
		return Grade{}, errors.New("student is not enrolled in this subject")
	}

	tx := g.db.Begin()
	gr.GradeID = 0
	if err := tx.Create(&gr).Error; err != nil {
		tx.Rollback()
		return gr, err
	}
	tx.Commit()
	return gr, nil
}

func (g *GradeService) Update(id int, gr Grade) (Grade, error) {
	for i, grade := range g.grades {
		if grade.GradeID == id {
			if _, ok := g.studentService.GetByID(gr.StudentID); !ok {
				return Grade{}, errors.New("student_id does not exist")
			}
			if _, ok := g.subjectService.GetByID(gr.SubjectID); !ok {
				return Grade{}, errors.New("subject_id does not exist")
			}
			gr.GradeID = id
			g.grades[i] = gr
			return gr, nil
		}
	}
	return Grade{}, ErrNotFound
}

func (g *GradeServiceDB) Update(id int, gr Grade) (Grade, error) {
	// Validar llaves foráneas
	var student Student
	var subject Subject
	if err := g.db.First(&student, gr.StudentID).Error; err != nil {
		return Grade{}, err
	}
	if err := g.db.First(&subject, gr.SubjectID).Error; err != nil {
		return Grade{}, err
	}
	gr.GradeID = id
	result := g.db.Model(&Grade{}).Where("grade_id = ?", id).Updates(gr)
	return gr, result.Error
}

func (g *GradeService) Delete(id int) error {
	for i, grade := range g.grades {
		if grade.GradeID == id {
			g.grades = append(g.grades[:i], g.grades[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

func (g *GradeServiceDB) Delete(id int) error {
	result := g.db.Exec("DELETE FROM grades WHERE grade_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (g *GradeService) GetByIDAndStudent(gradeID, studentID int) (Grade, bool) {
	for _, grade := range g.grades {
		if grade.GradeID == gradeID && grade.StudentID == studentID {
			return grade, true
		}
	}
	return Grade{}, false
}

func (g *GradeServiceDB) GetByIDAndStudent(gradeID, studentID int) (Grade, bool) {
	var grade Grade
	result := g.db.Where("grade_id = ? AND student_id = ?", gradeID, studentID).First(&grade)
	return grade, result.Error == nil
}

func (g *GradeService) GetByStudent(studentID int) []Grade {
	var result []Grade
	for _, grade := range g.grades {
		if grade.StudentID == studentID {
			result = append(result, grade)
		}
	}
	return result
}

func (g *GradeServiceDB) GetByStudent(studentID int) []Grade {
	var grades []Grade
	g.db.Preload("Student").Preload("Subject").
		Where("student_id = ?", studentID).
		Find(&grades)
	return grades
}

func (g *GradeService) GetByID(id int) (Grade, bool) {
	for _, grade := range g.grades {
		if grade.GradeID == id {
			return grade, true
		}
	}
	return Grade{}, false
}

func (g *GradeServiceDB) GetByID(id int) (Grade, bool) {
	var grade Grade
	result := g.db.Joins("Student").Joins("Subject").
		Where("grades.grade_id = ?", id).
		First(&grade)
	if result.Error != nil {
		return Grade{}, false
	}
	return grade, true
}

type GradeServiceInterface interface {
	Create(gr Grade) (Grade, error)
	Update(id int, gr Grade) (Grade, error)
	Delete(id int) error
	GetByIDAndStudent(gradeID, studentID int) (Grade, bool)
	GetByStudent(studentID int) []Grade
	GetByID(id int) (Grade, bool)
	GetAll() []Grade
}

func (g *GradeServiceDB) GetAll() []Grade {
	var grades []Grade
	g.db.Joins("Student").Joins("Subject").Find(&grades)
	return grades
}
