package services

type Subject struct {
	SubjectID int    `json:"subject_id"`
	Name      string `json:"name"`
}

type SubjectService struct {
	subjects []Subject
	nextID   int
}

func NewSubjectService() *SubjectService {
	return &SubjectService{nextID: 1}
}

func (s *SubjectService) GetAll() []Subject {
	return s.subjects
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
	return Subject{}, ErrNotFound
}

func (s *SubjectService) Delete(id int) error {
	for i, subject := range s.subjects {
		if subject.SubjectID == id {
			s.subjects = append(s.subjects[:i], s.subjects[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
