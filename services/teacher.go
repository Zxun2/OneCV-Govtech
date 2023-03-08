package services

import (
	"Zxun2/OneCV-Govtech/db"
	"Zxun2/OneCV-Govtech/errors"
	"Zxun2/OneCV-Govtech/models"
	"Zxun2/OneCV-Govtech/utils"
	"log"
)

// RegisterStudentsToTeacher registers multiple students to a teacher
func RegisterStudentsToTeacher(payload models.RegistersStudentsPayload) (response models.SuspendStudentResponse) {
	teacher, err := getTeacherByEmail(payload.Teacher)
	if err != nil {
		return models.SuspendStudentResponse{
			Response: errors.MakeResponseErr(err),
		}
	}

	students, err := getStudentsByEmail(payload.Students)
	if err != nil {
		return models.SuspendStudentResponse{
			Response: errors.MakeResponseErr(err),
		}
	}

	if len(students) != len(payload.Students) || teacher.Email == "" {
		return models.SuspendStudentResponse{
			Response: errors.MakeResponseErr(models.ErrStudentNotFound),
		}
	}

	err = db.Store.Model(&teacher).Association("Students").Append(students)
	if err != nil {
		return models.SuspendStudentResponse{
			Response: errors.MakeResponseErr(err),
		}
	}

	return models.SuspendStudentResponse{}
}

func getTeacherByEmail(email string) (*models.Teacher, error) {
	teacher := models.Teacher{}
	err :=  db.Store.Model(&models.Teacher{}).Where("email = ?", email).Find(&teacher).Error
	log.Println(teacher)
	return &teacher, err
}

// GetCommonStudents retrieves a list of students common to a given list of teachers
func GetCommonStudents(teacherEmails []string) ([]string, error) {
	var students []models.Student
	err := db.Store.Table("students").
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
func ListStudentsReceiveNotifications(payload models.ListStudentReceivingNotificationPayload) (models.ListStudentReceivingNotificationResponse) {
	teacherEmail := payload.Teacher
	notification := payload.Notification

	studentEmails := utils.ParseEmails(notification)

	log.Println(studentEmails)

	var teacher models.Teacher
	var students []string = studentEmails

	db.Store.Model(models.Teacher{}).Preload("Students").Where("email = ?", teacherEmail).
			First(&teacher)

	if teacher.Email == "" {
			// Handle error: no teacher found for the given email
			return models.ListStudentReceivingNotificationResponse{
				Response: errors.MakeResponseErr(models.ErrTeacherNotFound),
			}
	}

	for i := range teacher.Students {
			if teacher.Students[i].Status == "active" {
					students = append(students, teacher.Students[i].Email)
			}
	}

	return models.ListStudentReceivingNotificationResponse{
		Recipients: students,
	}
}