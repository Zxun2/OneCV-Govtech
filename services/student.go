package services

import (
	"Zxun2/OneCV-Govtech/models"

	"gorm.io/gorm"
)

// SuspendStudent suspends a student given the email
func SuspendStudent(db *gorm.DB, email string) (error) {
	student := models.Student{}
	result :=  db.Model(&models.Student{}).Where("email = ?", email).Find(&student).Update("status", models.SUSPENDED)
	if result.RowsAffected == 0 {
		return models.ErrStudentNotFound
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CreateStudent create a Student
func CreateStudent(db *gorm.DB, email string) (*gorm.DB, error) {
	isExist := checkStudentExists(db, email)
	if isExist {
		return nil, models.ErrStudentAlreadyExists
	}

	student := models.Student{
		Email: email,
	}
	result := db.Create(&student)
	return result, result.Error
}

// DeleteStudent deletes a student
func DeleteStudent(db *gorm.DB, email string) (*gorm.DB, error) {
	student := models.Student{}
	result := db.Where("email = ?", email).Delete(&student)
	return result, result.Error
}

func checkStudentExists(db *gorm.DB, email string) bool {
	var count int64
	result := db.Model(&models.Student{}).Where("email = ?", email).Count(&count)
	return result.Error == nil && count > 0
}

func getStudentsByEmail(db *gorm.DB, emails []string) ([]models.Student, error) {
	students := []models.Student{}
	err := db.Model(&models.Student{}).Where("email IN (?)", emails).Find(&students).Error
	return students, err
}

// GetStudent gets a student by email
func GetStudent(db *gorm.DB, email string) (*models.Student, error) {
	student := models.Student{}
	result := db.Where("email = ?", email).First(&student)
	return &student, result.Error
}