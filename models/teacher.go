package models

import (
	"Zxun2/OneCV-Govtech/errors"
)


const (
	teacherModelName = "Teacher"
)

var (
	// ErrTeacherNotFound is the error for when a Teacher is not found
	ErrTeacherNotFound = &errors.RecordNotFoundError{Model: teacherModelName}
)

// Teacher is the model for the teachers table
type Teacher struct {
	ID       			uint8     `gorm:"primaryKey" json:"id"`
	Email    			string    `gorm:"not null" json:"email"`
	Students 			[]Student `gorm:"many2many:teacher_students" json:"students"`
}

