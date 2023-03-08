package seed

import (
	"Zxun2/OneCV-Govtech/db"
	"Zxun2/OneCV-Govtech/models"
	"Zxun2/OneCV-Govtech/utils"

	"gorm.io/gorm/clause"
)

// SeedData generates sample seed data
func SeedData() (err error) {

	var teachers []models.Teacher
	var students []models.Student
	var suspendedStudents []models.Student
	var commonStudents []models.Student
	var teachersWithCommonStudents []models.Teacher
	var commonTeachers []models.Teacher
	var studentsWithCommonTeachers []models.Student

	for i := 1; i <= 5; i++  {
		commonStudents = append(commonStudents, 
			models.Student{
			Email:  utils.RandomEmail(),
		})
	}

	for i := 1; i <= 5; i++  {
		suspendedStudents = append(suspendedStudents, 
			models.Student{
			Email:  utils.RandomEmail(),
			Status: models.SUSPENDED,
		})
	}

	for i := 1; i <= 5; i++  {
		students = append(students,
			models.Student{
			Email:  utils.RandomEmail(),
		})
	}

	// Create sample teachers with common students
	for i := 1; i <= 5; i++ {
		teachersWithCommonStudents = append(teachersWithCommonStudents, models.Teacher{
			Email: utils.RandomEmail(),
			Students: commonStudents,
		})
	}

	// Create sample teachers
	for i := 1; i <= 5; i++ {
		teachers = append(teachers, models.Teacher{
			Email: utils.RandomEmail(),
		})	
	}

	for i := 1; i <= 5; i++  {
		commonTeachers = append(commonTeachers, 
			models.Teacher{
			Email:  utils.RandomEmail(),
		})
	}

	for i := 1; i <= 5; i++  {
		studentsWithCommonTeachers = append(studentsWithCommonTeachers,
			models.Student{
			Email:  utils.RandomEmail(),
			Teachers: commonTeachers,
		})
	}

	db.Store.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&teachers)
		
	db.Store.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&students)
		
	db.Store.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&suspendedStudents)
		
	db.Store.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&studentsWithCommonTeachers)
		
	db.Store.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&teachersWithCommonStudents)

	return
}


