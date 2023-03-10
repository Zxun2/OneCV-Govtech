package tests

import (
	"Zxun2/OneCV-Govtech/models"
	"Zxun2/OneCV-Govtech/services"
	"Zxun2/OneCV-Govtech/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomStudentEntry(t *testing.T, email string) (models.Student) {
	student := models.Student{
		Email: email, 
	}
	db, err := services.CreateStudent(testDb, student.Email)
	require.NoError(t, err)
	require.NotEmpty(t, db)
	require.Equal(t, int64(1), db.RowsAffected)

	return student
}

func CreateListOfRandomStudentEntry(t *testing.T, num int) (students []models.Student) {
	for i := 0; i < num; i++ {
		students = append(students, CreateRandomStudentEntry(t, utils.RandomEmail()))
	}
	require.Equal(t, num, len(students))
	return students
}

func GetStudent(t *testing.T, email string) (models.Student) {
	student, err := services.GetStudent(testDb, email)
	require.NoError(t, err)
	require.NotEmpty(t, student)
	require.Equal(t, email, student.Email)
	return *student
}

func DeleteStudent(t *testing.T, email string) {
	db, err := services.DeleteStudent(testDb, email)
	require.NoError(t, err)
	require.NotEmpty(t, db)
	require.Equal(t, int64(1), db.RowsAffected)
}

func TestCRUDStudent(t *testing.T) {
	newStudent := CreateRandomStudentEntry(t, utils.RandomEmail())

	Student, err := services.GetStudent(testDb, newStudent.Email)
	require.NoError(t, err)
	require.NotEmpty(t, Student)
	require.Equal(t, newStudent.Email, Student.Email)

	db, err := services.DeleteStudent(testDb, newStudent.Email)
	require.NoError(t, err)
	require.NotEmpty(t, db)
	require.Equal(t, int64(1), db.RowsAffected)
}

func TestSuspendStudent(t *testing.T) {
	newStudent := CreateRandomStudentEntry(t, utils.RandomEmail())

	err := services.SuspendStudent(testDb, newStudent.Email)
	require.NoError(t, err)

	student, err := services.GetStudent(testDb, newStudent.Email)
	require.NoError(t, err)
	require.NotEmpty(t, student)
	
	db, err := services.DeleteStudent(testDb, newStudent.Email)
	require.NoError(t, err)
	require.NotEmpty(t, db)
	require.Equal(t, int64(1), db.RowsAffected)

	require.Equal(t, models.SUSPENDED, student.Status)
}