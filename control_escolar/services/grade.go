package services

import "errors"

type Grade struct {
	GradeID   int     `json:"grade_id"`
	StudentID int     `json:"student_id"`
	SubjectID int     `json:"subject_id"`
	Grade     float64 `json:"grade"`
}

type GradeService struct {
	grades         []Grade
	nextID         int
	studentService *StudentService
	subjectService *SubjectService
}

var ErrNotFound = errors.New("not found")

func NewGradeService(stSrv *StudentService, sbSrv *SubjectService) *GradeService {
	return &GradeService{nextID: 1, studentService: stSrv, subjectService: sbSrv}
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

func (g *GradeService) Delete(id int) error {
	for i, grade := range g.grades {
		if grade.GradeID == id {
			g.grades = append(g.grades[:i], g.grades[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
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
