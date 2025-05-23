package services

import (
	"gorm.io/gorm"
)

type StudentService struct {
	students []Student
	nextID   int
}

func NewStudentService() *StudentService {
	return &StudentService{nextID: 1}
}

func (s *StudentService) GetAll() []Student {
	return s.students
}

func (s *StudentService) GetByID(id int) (Student, bool) {
	for _, st := range s.students {
		if st.StudentID == id {
			return st, true
		}
	}
	return Student{}, false
}

func (s *StudentService) Create(st Student) Student {
	st.StudentID = s.nextID
	s.nextID++
	s.students = append(s.students, st)
	return st
}

func (s *StudentService) Update(id int, st Student) (Student, error) {
	for i, student := range s.students {
		if student.StudentID == id {
			st.StudentID = id
			s.students[i] = st
			return st, nil
		}
	}
	return Student{}, ErrNotFound
}

func (s *StudentService) Delete(id int) error {
	for i, student := range s.students {
		if student.StudentID == id {
			s.students = append(s.students[:i], s.students[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

type StudentServiceDB struct {
	db *gorm.DB
}

func NewStudentServiceDB(db *gorm.DB) *StudentServiceDB {
	return &StudentServiceDB{db: db}
}

func (s *StudentServiceDB) GetAll() []Student {
	var students []Student
	s.db.Find(&students)
	return students
}

func (s *StudentServiceDB) Create(st Student) Student {
	tx := s.db.Begin()
	// Asegurarnos que el ID sea 0 para que GORM autogenere el ID
	st.StudentID = 0
	if err := tx.Create(&st).Error; err != nil {
		tx.Rollback()
		return st
	}
	tx.Commit()
	return st
}

func (s *StudentServiceDB) GetByID(id int) (Student, bool) {
	var student Student
	result := s.db.Preload("Subjects").Preload("Grades").Where("student_id = ?", id).First(&student)
	return student, result.Error == nil
}

func (s *StudentServiceDB) Update(id int, st Student) (Student, error) {
	st.StudentID = id
	result := s.db.Model(&Student{}).Where("student_id = ?", id).Updates(st)
	return st, result.Error
}

func (s *StudentServiceDB) Delete(id int) error {
	result := s.db.Exec("DELETE FROM students WHERE student_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (s *StudentServiceDB) EnrollInSubject(studentID, subjectID int) error {
	return s.db.Exec("INSERT INTO student_subjects (student_id, subject_id) VALUES (?, ?)",
		studentID, subjectID).Error
}

func (s *StudentServiceDB) UnenrollFromSubject(studentID, subjectID int) error {
	return s.db.Exec("DELETE FROM student_subjects WHERE student_id = ? AND subject_id = ?",
		studentID, subjectID).Error
}

type StudentServiceInterface interface {
	GetAll() []Student
	GetByID(id int) (Student, bool)
	Create(st Student) Student
	Update(id int, st Student) (Student, error)
	Delete(id int) error
	EnrollInSubject(studentID, subjectID int) error
	UnenrollFromSubject(studentID, subjectID int) error
}

type SubjectServiceInterface interface {
	GetAll() []Subject
	GetByID(id int) (Subject, bool)
	Create(sub Subject) Subject
	Update(id int, sub Subject) (Subject, error)
	Delete(id int) error
}
