package tests

import (
	"Zxun2/OneCV-Govtech/models"
	"Zxun2/OneCV-Govtech/services"
	"Zxun2/OneCV-Govtech/utils"
	"testing"

	"github.com/stretchr/testify/require"
)


func CreateRandomTeacherEntry(t *testing.T) (models.Teacher) {
	teacher := models.Teacher{
		Email: utils.RandomEmail(),
	}
	db, err := services.CreateTeacher(testDb, teacher.Email)
	require.NoError(t, err)
	require.NotEmpty(t, db)
	require.Equal(t, int64(1), db.RowsAffected)
	return teacher
}

func CreateTeacherWithRandomStudents(t *testing.T) (models.Teacher) {
	students := utils.CreateListOfRandomEmails(3)
	for i := range students {
		CreateRandomStudentEntry(t, students[i])
	}
	teacher := CreateRandomTeacherEntry(t)
	err :=	services.RegisterStudentsToTeacher(testDb, teacher.Email, students)
	require.NoError(t, err)
	teacher = GetTeacher(t, teacher.Email)
	require.Equal(t, len(students), len(teacher.Students))
	require.Equal(t, students[0], teacher.Students[0].Email)
	require.Equal(t, students[1], teacher.Students[1].Email)
	require.Equal(t, students[2], teacher.Students[2].Email)
	return teacher
}

func CreateTeacherWithSpecifiedStudents(t *testing.T, students []string) (models.Teacher) {
	for i := range students {
		// Check if student is in database
		student, err := services.GetStudent(testDb, students[i])
		if student.Email == "" {
			// If not, create a new student
			CreateRandomStudentEntry(t, students[i])
		} else {
			require.NoError(t, err)
		}
	}

	teacher := CreateRandomTeacherEntry(t)
	err :=	services.RegisterStudentsToTeacher(testDb, teacher.Email, students)
	require.NoError(t, err)
	teacher = GetTeacher(t, teacher.Email)
	require.Equal(t, len(students), len(teacher.Students))
	require.Equal(t, students[0], teacher.Students[0].Email)
	require.Equal(t, students[1], teacher.Students[1].Email)
	require.Equal(t, students[2], teacher.Students[2].Email)
	return teacher
}

func TestRegisterStudentsToTeacher(t *testing.T) {
	teacher := CreateTeacherWithRandomStudents(t)

	// Delete students
	for i := range teacher.Students {
		DeleteStudent(t, teacher.Students[i].Email)
	}
	DeleteTeacher(t, teacher.Email)
}

func GetTeacher(t *testing.T, email string) (models.Teacher) {
	teacher, err := services.GetTeacher(testDb, email)
	require.NoError(t, err)
	require.NotEmpty(t, teacher)
	require.Equal(t, email, teacher.Email)
	return *teacher
}

func DeleteTeacher(t *testing.T, email string) {
	db, err := services.DeleteTeacher(testDb, email)
	require.NoError(t, err)
	require.NotEmpty(t, db)
	require.Equal(t, int64(1), db.RowsAffected)
}

func TestCRUDTeacher(t *testing.T) {
	newTeacher := CreateRandomTeacherEntry(t)
	GetTeacher(t, newTeacher.Email)
	DeleteTeacher(t, newTeacher.Email)
}

func TestCommonStudents(t *testing.T) {
	emails := utils.CreateListOfRandomEmails(3)
	sampleTeacherOne := CreateTeacherWithSpecifiedStudents(t, emails)
	sampleTeacherTwo := CreateTeacherWithSpecifiedStudents(t, emails)
	sampleTeacherThree := CreateTeacherWithSpecifiedStudents(t, emails)

	// Test common students
	commonStudents, err := services.GetCommonStudents(testDb, []string{
		sampleTeacherOne.Email, sampleTeacherTwo.Email, sampleTeacherThree.Email})
	require.NoError(t, err)
	require.Equal(t, len(emails), len(commonStudents))

	// Test common students with non-existing teacher
	commonStudents, err = services.GetCommonStudents(testDb, []string{sampleTeacherOne.Email, sampleTeacherTwo.Email, "non-existing-teacher"})
	require.NoError(t, err)
	require.Equal(t, 0, len(commonStudents))

	// Test common students with empty teacher list
	commonStudents, err = services.GetCommonStudents(testDb, []string{})
	require.ErrorContains(t, err, "No teacher emails provided")
	require.Equal(t, 0, len(commonStudents))

	// Test common students with one teacher
	commonStudents, err = services.GetCommonStudents(testDb, []string{sampleTeacherOne.Email})
	require.NoError(t, err)
	require.Equal(t, len(emails), len(commonStudents))

	// Clean up
	DeleteTeacher(t, sampleTeacherOne.Email)
	DeleteTeacher(t, sampleTeacherTwo.Email)
	DeleteTeacher(t, sampleTeacherThree.Email)

	DeleteStudent(t, emails[0])
	DeleteStudent(t, emails[1])
	DeleteStudent(t, emails[2])
}

func TestListStudentsWhoCanReceiveNotifications(t *testing.T) {
	emails := utils.CreateListOfRandomEmails(3)
	randomEmails := utils.CreateListOfRandomEmails(3) 
	for i := range randomEmails{
		CreateRandomStudentEntry(t, randomEmails[i])
	}

	sampleTeacherOne := CreateTeacherWithSpecifiedStudents(t, emails)
	sampleNotification := "Hello students!"

	// Only registered students under teacher
	commonStudents, err := services.ListStudentsReceiveNotifications(testDb, sampleTeacherOne.Email, sampleNotification)
	require.NoError(t, err)
	require.Equal(t, len(emails), len(commonStudents))

	// Registered students under teacher and students not under teacher
	newNotification := sampleNotification + " @" + randomEmails[0] + " @" + randomEmails[1]
	commonStudents, err = services.ListStudentsReceiveNotifications(testDb, sampleTeacherOne.Email, newNotification)
	require.NoError(t, err)
	require.Equal(t, len(emails) + 2, len(commonStudents))

	DeleteTeacher(t, sampleTeacherOne.Email)
	DeleteStudent(t, emails[0])
	DeleteStudent(t, emails[1])
	DeleteStudent(t, emails[2])
	DeleteStudent(t, randomEmails[0])
	DeleteStudent(t, randomEmails[1])
	DeleteStudent(t, randomEmails[2])
}