package services

import (
	"gorm.io/gorm"
)

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

type SubjectServiceDB struct {
	db *gorm.DB
}

func NewSubjectServiceDB(db *gorm.DB) *SubjectServiceDB {
	return &SubjectServiceDB{db: db}
}

func (s *SubjectServiceDB) GetAll() []Subject {
	var subjects []Subject
	s.db.Find(&subjects)
	return subjects
}

func (s *SubjectServiceDB) GetByID(id int) (Subject, bool) {
	var subject Subject
	result := s.db.Where("subject_id = ?", id).First(&subject)
	if result.Error != nil {
		return Subject{}, false
	}
	return subject, true
}

func (s *SubjectServiceDB) Create(sub Subject) Subject {
	tx := s.db.Begin()
	// Asegurarnos que el ID sea 0 para que GORM autogenere el ID
	sub.SubjectID = 0
	if err := tx.Create(&sub).Error; err != nil {
		tx.Rollback()
		return sub
	}
	tx.Commit()
	return sub
}

func (s *SubjectServiceDB) Update(id int, sub Subject) (Subject, error) {
	sub.SubjectID = id
	result := s.db.Model(&Subject{}).Where("subject_id = ?", id).Updates(sub)
	return sub, result.Error
}

func (s *SubjectServiceDB) Delete(id int) error {
	result := s.db.Exec("DELETE FROM subjects WHERE subject_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}
