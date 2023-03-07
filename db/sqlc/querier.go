// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateStudent(ctx context.Context, email string) (sql.Result, error)
	CreateTeacher(ctx context.Context, email string) (sql.Result, error)
	DeleteStudent(ctx context.Context, id int64) error
	DeleteTeacher(ctx context.Context, id int64) error
	GetStudent(ctx context.Context, id int64) (Student, error)
	GetTeacher(ctx context.Context, id int64) (Teacher, error)
	ListStudents(ctx context.Context) ([]Student, error)
	ListTeachers(ctx context.Context) ([]Teacher, error)
}

var _ Querier = (*Queries)(nil)
