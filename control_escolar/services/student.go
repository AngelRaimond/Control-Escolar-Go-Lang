package services

type Student struct {
	StudentID int    `json:"student_id"`
	Name      string `json:"name"`
	Group     string `json:"group"`
	Email     string `json:"email"`
}

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
