package services

// Student model
type Student struct {
	StudentID int       `gorm:"primaryKey;autoIncrement" json:"student_id"`
	Name      string    `json:"name"`
	Group     string    `json:"group"`
	Email     string    `json:"email"`
	Subjects  []Subject `gorm:"many2many:student_subjects;joinForeignKey:StudentID;joinReferences:SubjectID" json:"subjects,omitempty"`
	Grades    []Grade   `gorm:"foreignKey:StudentID" json:"grades,omitempty"`
}

// Subject model
type Subject struct {
	SubjectID int       `gorm:"primaryKey;autoIncrement" json:"subject_id"`
	Name      string    `json:"name"`
	Students  []Student `gorm:"many2many:student_subjects;joinForeignKey:SubjectID;joinReferences:StudentID" json:"students,omitempty"`
	Grades    []Grade   `gorm:"foreignKey:SubjectID" json:"grades,omitempty"`
}

// Grade model
type Grade struct {
	GradeID   int     `gorm:"primaryKey;autoIncrement" json:"grade_id"`
	StudentID int     `gorm:"not null;index" json:"student_id"`
	Student   Student `gorm:"foreignKey:StudentID" json:"-"`
	SubjectID int     `gorm:"not null;index" json:"subject_id"`
	Subject   Subject `gorm:"foreignKey:SubjectID" json:"-"`
	Grade     float64 `json:"grade"`
}

// StudentSubject model para la relación many-to-many
type StudentSubject struct {
	StudentID int `gorm:"primaryKey"`
	SubjectID int `gorm:"primaryKey"`
}
