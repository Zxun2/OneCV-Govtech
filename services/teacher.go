package services

import (
	"Zxun2/OneCV-Govtech/models"
	"Zxun2/OneCV-Govtech/utils"
	"errors"

	"gorm.io/gorm"
)

// checkTeacherExists checks if a teacher exists in the database
func checkTeacherExists(db *gorm.DB, email string) bool {
	var count int64
	db.Model(&models.Teacher{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// CreateTeacher create a teacher
func CreateTeacher(db *gorm.DB, email string) (*gorm.DB, error) {
	isExist := checkTeacherExists(db, email)
	if isExist {
		return nil, models.ErrTeacherAlreadyExists
	}

	teacher := models.Teacher{
		Email: email,
	}
	result := db.Create(&teacher)
	return result, result.Error
}

// GetTeacher gets a teacher by email
func GetTeacher(db *gorm.DB, email string) (*models.Teacher, error) {
	teacher := models.Teacher{}
	result := db.Preload("Students").Where("email = ?", email).First(&teacher)
	return &teacher, result.Error
}

// DeleteTeacher deletes a teacher
func DeleteTeacher(db *gorm.DB, email string) (*gorm.DB, error) {
	teacher := models.Teacher{}
	result := db.Where("email = ?", email).Delete(&teacher)
	return result, result.Error
}

// RegisterStudentsToTeacher registers multiple students to a teacher
func RegisterStudentsToTeacher(db *gorm.DB, teacherEmail string, studentEmails []string) (error) {
	studentEmails = utils.RemoveDuplicates(studentEmails)

	teacher, err := GetTeacher(db, teacherEmail)
	if err != nil {
		return err
	}

	students, err := getStudentsByEmail(db, studentEmails)
	if err != nil {
		return err
	}

	if len(students) != len(studentEmails)  {
		return models.ErrStudentNotFound
	}

	if teacher.Email == "" {
		return models.ErrTeacherNotFound
	}

	err = db.Model(&teacher).Association("Students").Append(students)
	if err != nil {
		return err
	}

	return nil
}

// GetCommonStudents retrieves a list of students common to a given list of teachers
func GetCommonStudents(db *gorm.DB, teacherEmails []string) ([]string, error) {
	teacherEmails = utils.RemoveDuplicates(teacherEmails)

	if len(teacherEmails) == 0 {
		return []string{}, errors.New("No teacher emails provided")
	}

	var students []models.Student
	err := db.Table("students").
		Select("students.email").
		Joins("JOIN teacher_students ON students.id = teacher_students.student_id").
		Joins("JOIN teachers ON teachers.id = teacher_students.teacher_id").
		Where("teachers.email IN (?) AND students.status = ?", teacherEmails, "active").
		Group("students.id, students.email").
		Having("COUNT(DISTINCT teachers.id) = ?", len(teacherEmails)).
		Find(&students).Error

		if err != nil {
			return []string{}, err
		}

		var studentEmails []string 

		for i := range students {
			studentEmails = append(studentEmails, students[i].Email)
		}

    return studentEmails, nil
}

// ListStudentsReceiveNotifications retrieves a list of students who can receive a given notification.
func ListStudentsReceiveNotifications(db *gorm.DB, teacherEmail string, notification string) ([]string, error) {
	studentEmails := utils.ParseEmails(notification)

	var teacher models.Teacher
	var students []string = studentEmails

	db.Model(models.Teacher{}).Preload("Students").Where("email = ?", teacherEmail).
			First(&teacher)

	if teacher.Email == "" {
			// Handle error: no teacher found for the given email
			return nil, models.ErrTeacherNotFound
	}

	for i := range teacher.Students {
			if teacher.Students[i].Status == "active" {
					students = append(students, teacher.Students[i].Email)
			}
	}

	return students, nil
}