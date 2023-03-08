package models

// Teacher is the model for the teachers table
type Teacher struct {
	ID       uint8     `gorm:"primaryKey" json:"id"`
	Email    string    `gorm:"not null" json:"email"`
	Students []Student `gorm:"many2many:teacher_students" json:"students"`
}

// SuspendStudentPayload - incoming request
type SuspendStudentPayload struct {
	Student string `json:"student"`
}

// SuspendStudentResponse - outgoing response
type SuspendStudentResponse struct {
	Response
}

// RegistersStudentsPayload - incoming request
type RegistersStudentsPayload struct {
	Teacher 	string   `json:"teacher"`
	Students 	[]string `json:"students"`
}

// RegistersStudentsResponse - outgoing response
type RegistersStudentsResponse struct {
	Response
}

// RetrieveCommonStudentsPayload - incoming request
type RetrieveCommonStudentsPayload struct {
	Students []string `json:"students"`
}

// RetrieveCommonStudentsResponse - outgoing response
type RetrieveCommonStudentsResponse struct {
	Response
	Students []string `json:"students"`
}

// ListStudentReceivingNotificationPayload - incoming request
type ListStudentReceivingNotificationPayload struct {
	Teacher 	string `json:"teacher"`
	Notification string `json:"notification"`
}

// ListStudentReceivingNotificationResponse - outgoing response
type ListStudentReceivingNotificationResponse struct {
	Response
	Recipients []string `json:"recipients"`
}