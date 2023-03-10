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

func TestRegisterStudentsToTeacher(t *testing.T) {
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

	for i := range students {
		DeleteStudent(t, students[i])
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
	sampleStudents := []models.Student {
			{
				Email: "emily@gmail.com",
			},
			{
				Email: "zongxun@gmail.com",
			},
		}

	sampleTeacherOne := models.Teacher {
		Email: "sampleOne@gmail.com",
		Students: sampleStudents,
	}

	sampleTeacherTwo := models.Teacher {
		Email: "sampleTwo@gmail.com",
		Students: sampleStudents,
	}

	sampleTeacherThree := models.Teacher {
		Email: "sampleThree@gmail.com",
		Students: sampleStudents,
	}

	testDb.Create(&sampleTeacherOne)
	testDb.Create(&sampleTeacherTwo)
	testDb.Create(&sampleTeacherThree)

	students, err := services.GetCommonStudents(testDb, []string{
		"sampleOne@gmail.com",
		"sampleTwo@gmail.com",
		"sampleThree@gmail.com",
	})

	require.NoError(t, err)
	require.Equal(t, 2, len(students))
	require.Equal(t, "emily@gmail.com", students[0])
	require.Equal(t, "zongxun@gmail.com", students[1])
	
	_, err = services.DeleteTeacher(testDb, sampleTeacherOne.Email)
	_, err = services.DeleteTeacher(testDb, sampleTeacherTwo.Email)
	_, err = services.DeleteTeacher(testDb, sampleTeacherThree.Email)
	_, err = services.DeleteTeacher(testDb, sampleTeacherThree.Email)
	_, err = services.DeleteStudent(testDb, sampleTeacherThree.Students[0].Email)
	_, err = services.DeleteStudent(testDb, sampleTeacherThree.Students[1].Email)
}