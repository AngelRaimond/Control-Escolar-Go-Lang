package main

import (
	"errors"
)

// Student model
type Student struct {
	StudentID int    `json:"student_id"`
	Name      string `json:"name"`
	Group     string `json:"group"`
	Email     string `json:"email"`
}

// Subject model
type Subject struct {
	SubjectID int    `json:"subject_id"`
	Name      string `json:"name"`
}

// Grade model
type Grade struct {
	GradeID   int     `json:"grade_id"`
	StudentID int     `json:"student_id"`
	SubjectID int     `json:"subject_id"`
	Grade     float64 `json:"grade"`
}

// --- Student Service ---
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
	return Student{}, errors.New("student not found")
}
func (s *StudentService) Delete(id int) error {
	for i, student := range s.students {
		if student.StudentID == id {
			s.students = append(s.students[:i], s.students[i+1:]...)
			return nil
		}
	}
	return errors.New("student not found")
}

// --- Subject Service ---
type SubjectService struct {
	subjects []Subject
	nextID   int
}

func NewSubjectService() *SubjectService {
	return &SubjectService{nextID: 1}
}
func (s *SubjectService) GetByID(id int) (Subject, bool) {
	for _, sub := range s.subjects {
		if sub.SubjectID == id {
			return sub, true
		}
	}
	return Subject{}, false
}
func (s *SubjectService) Create(sub Subject) Subject {
	sub.SubjectID = s.nextID
	s.nextID++
	s.subjects = append(s.subjects, sub)
	return sub
}
func (s *SubjectService) Update(id int, sub Subject) (Subject, error) {
	for i, subject := range s.subjects {
		if subject.SubjectID == id {
			sub.SubjectID = id
			s.subjects[i] = sub
			return sub, nil
		}
	}
	return Subject{}, errors.New("subject not found")
}
func (s *SubjectService) Delete(id int) error {
	for i, subject := range s.subjects {
		if subject.SubjectID == id {
			s.subjects = append(s.subjects[:i], s.subjects[i+1:]...)
			return nil
		}
	}
	return errors.New("subject not found")
}

// --- Grade Service ---
type GradeService struct {
	grades         []Grade
	nextID         int
	studentService *StudentService
	subjectService *SubjectService
}

func NewGradeService(stSrv *StudentService, sbSrv *SubjectService) *GradeService {
	return &GradeService{nextID: 1, studentService: stSrv, subjectService: sbSrv}
}
func (g *GradeService) Create(gr Grade) (Grade, error) {
	// Validate foreign keys
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
func (g *GradeService) Update(id int, gr Grade) (Grade, error) {
	for i, grade := range g.grades {
		if grade.GradeID == id {
			// Validate foreign keys
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
	return Grade{}, errors.New("grade not found")
}
func (g *GradeService) Delete(id int) error {
	for i, grade := range g.grades {
		if grade.GradeID == id {
			g.grades = append(g.grades[:i], g.grades[i+1:]...)
			return nil
		}
	}
	return errors.New("grade not found")
}
func (g *GradeService) GetByIDAndStudent(gradeID, studentID int) (Grade, bool) {
	for _, grade := range g.grades {
		if grade.GradeID == gradeID && grade.StudentID == studentID {
			return grade, true
		}
	}
	return Grade{}, false
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
