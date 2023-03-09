package services

import (
	"Zxun2/OneCV-Govtech/models"

	"gorm.io/gorm"
)

// SuspendStudent suspends a student given the email
func SuspendStudent(db *gorm.DB, email string) (error) {
	student := models.Student{}
	result :=  db.Model(&models.Student{Email: email}).Find(&student).Update("status", models.SUSPENDED)
	if result.RowsAffected == 0 {
		return models.ErrStudentNotFound
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func checkStudentExists(db *gorm.DB, email string) bool {
	var count int64
	result := db.Model(&models.Student{Email: email}).Count(&count)
	return result.Error == nil && count > 0
}

func getStudentsByEmail(db *gorm.DB, emails []string) ([]models.Student, error) {
	students := []models.Student{}
	err := db.Model(&models.Student{}).Where("email IN (?)", emails).Find(&students).Error
	return students, err
}