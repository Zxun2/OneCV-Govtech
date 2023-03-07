package models

// Status is the model for the status columns
type Status string

const  (
	// SUSPENDED indicates that the student is not active
	SUSPENDED Status = "suspended"
	// ACTIVE Indicates that the student is active
	ACTIVE 		Status = "active"
)

// Student is the model for the students table
type Student struct {
	ID     uint8  `gorm:"primaryKey" json:"id"`
	Email  string `gorm:"not null" json:"email"`
	Status Status `gorm:"not null" json:"status"`
}
