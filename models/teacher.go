package models

import (
	"Zxun2/OneCV-Govtech/api"
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

// SuspendStudentPayload - incoming request
type SuspendStudentPayload struct {
	Student 			string `json:"student"`
}

// SuspendStudentResponse - outgoing response
type SuspendStudentResponse struct {
	api.Response
}

// RegistersStudentsPayload - incoming request
type RegistersStudentsPayload struct {
	Teacher 			string   `json:"teacher"`
	Students 			[]string `json:"students"`
}

// RegistersStudentsResponse - outgoing response
type RegistersStudentsResponse struct {
	api.Response
}

// RetrieveCommonStudentsPayload - incoming request
type RetrieveCommonStudentsPayload struct {
	Students 			[]string `json:"students"`
}

// RetrieveCommonStudentsResponse - outgoing response
type RetrieveCommonStudentsResponse struct {
	api.Response
	Students 			[]string `json:"students"`
}

// ListStudentReceivingNotificationPayload - incoming request
type ListStudentReceivingNotificationPayload struct {
	Teacher 			string `json:"teacher"`
	Notification 	string `json:"notification"`
}

// ListStudentReceivingNotificationResponse - outgoing response
type ListStudentReceivingNotificationResponse struct {
	api.Response
	Recipients 		[]string `json:"recipients"`
}