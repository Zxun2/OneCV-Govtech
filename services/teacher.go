package services

import (
	"Zxun2/OneCV-Govtech/models"
	"Zxun2/OneCV-Govtech/utils"
	"log"

	"gorm.io/gorm"
)

// RegisterStudentsToTeacher registers multiple students to a teacher
func RegisterStudentsToTeacher(db *gorm.DB, teacherEmail string, studentEmails []string) (error) {
	teacher, err := getTeacherByEmail(db, teacherEmail)
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

func getTeacherByEmail(db *gorm.DB, email string) (*models.Teacher, error) {
	teacher := models.Teacher{}
	err :=  db.Model(&models.Teacher{}).Where("email = ?", email).Find(&teacher).Error
	log.Println(teacher)
	return &teacher, err
}

// GetCommonStudents retrieves a list of students common to a given list of teachers
func GetCommonStudents(db *gorm.DB, teacherEmails []string) ([]string, error) {
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
			return nil, err
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