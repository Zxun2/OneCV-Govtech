package models

import "Zxun2/OneCV-Govtech/errors"

// Status is the model for the status columns
type Status string

const  (
	// SUSPENDED indicates that the student is not active
	SUSPENDED Status = "suspended"
	// ACTIVE Indicates that the student is active
	ACTIVE 		Status = "active"
)

const (
	studentModelName = "Student"
)

var (
	// ErrStudentNotFound is the error for when a student is not found
	ErrStudentNotFound = &errors.RecordNotFoundError{Model: studentModelName}
	// ErrStudentAlreadyExists is the error for when a student already exists
	ErrStudentAlreadyExists = &errors.RecordAlreadyExistsError{Model: studentModelName}
)

// Student is the model for the students table
type Student struct {
	ID     		uint8  		`gorm:"primaryKey" json:"id"`
	Email  		string 		`gorm:"not null;unique" json:"email"`
	Status 		Status 		`gorm:"default:active;not null" json:"status"`
	Teachers 	[]Teacher `gorm:"many2many:teacher_students;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"teacher"`
}