package services

import (
	"Zxun2/OneCV-Govtech/db"
	"Zxun2/OneCV-Govtech/errors"
	"Zxun2/OneCV-Govtech/models"
)

// SuspendStudent suspends a student given the email
func SuspendStudent(payload models.SuspendStudentPayload) (response models.SuspendStudentResponse) {
	_, err := getStudentByEmailAndUpdateStatus(payload.Student)
	if err != nil {
		return models.SuspendStudentResponse{
			Response: errors.MakeResponseErr(err),
		}
	}
	return models.SuspendStudentResponse{}
}

func getStudentByEmailAndUpdateStatus(email string) (*models.Student, error) {
	student := models.Student{}
	result :=  db.Store.Model(&models.Student{Email: email}).Find(&student).Update("status", models.SUSPENDED)
	if result.RowsAffected == 0 {
		return nil, models.ErrStudentNotFound
	}
	return nil, result.Error
}

func checkStudentExists(email string) bool {
	var count int64
	result := db.Store.Model(&models.Student{Email: email}).Count(&count)
	return result.Error == nil && count > 0
}

func getStudentsByEmail(emails []string) ([]models.Student, error) {
	students := []models.Student{}
	err := db.Store.Model(&models.Student{}).Where("email IN (?)", emails).Find(&students).Error
	return students, err
}