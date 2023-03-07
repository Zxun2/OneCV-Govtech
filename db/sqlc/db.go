// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createStudentStmt, err = db.PrepareContext(ctx, createStudent); err != nil {
		return nil, fmt.Errorf("error preparing query CreateStudent: %w", err)
	}
	if q.createTeacherStmt, err = db.PrepareContext(ctx, createTeacher); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTeacher: %w", err)
	}
	if q.deleteStudentStmt, err = db.PrepareContext(ctx, deleteStudent); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteStudent: %w", err)
	}
	if q.deleteTeacherStmt, err = db.PrepareContext(ctx, deleteTeacher); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTeacher: %w", err)
	}
	if q.getStudentStmt, err = db.PrepareContext(ctx, getStudent); err != nil {
		return nil, fmt.Errorf("error preparing query GetStudent: %w", err)
	}
	if q.getTeacherStmt, err = db.PrepareContext(ctx, getTeacher); err != nil {
		return nil, fmt.Errorf("error preparing query GetTeacher: %w", err)
	}
	if q.listStudentsStmt, err = db.PrepareContext(ctx, listStudents); err != nil {
		return nil, fmt.Errorf("error preparing query ListStudents: %w", err)
	}
	if q.listTeachersStmt, err = db.PrepareContext(ctx, listTeachers); err != nil {
		return nil, fmt.Errorf("error preparing query ListTeachers: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createStudentStmt != nil {
		if cerr := q.createStudentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createStudentStmt: %w", cerr)
		}
	}
	if q.createTeacherStmt != nil {
		if cerr := q.createTeacherStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTeacherStmt: %w", cerr)
		}
	}
	if q.deleteStudentStmt != nil {
		if cerr := q.deleteStudentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteStudentStmt: %w", cerr)
		}
	}
	if q.deleteTeacherStmt != nil {
		if cerr := q.deleteTeacherStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTeacherStmt: %w", cerr)
		}
	}
	if q.getStudentStmt != nil {
		if cerr := q.getStudentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getStudentStmt: %w", cerr)
		}
	}
	if q.getTeacherStmt != nil {
		if cerr := q.getTeacherStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTeacherStmt: %w", cerr)
		}
	}
	if q.listStudentsStmt != nil {
		if cerr := q.listStudentsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listStudentsStmt: %w", cerr)
		}
	}
	if q.listTeachersStmt != nil {
		if cerr := q.listTeachersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTeachersStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                DBTX
	tx                *sql.Tx
	createStudentStmt *sql.Stmt
	createTeacherStmt *sql.Stmt
	deleteStudentStmt *sql.Stmt
	deleteTeacherStmt *sql.Stmt
	getStudentStmt    *sql.Stmt
	getTeacherStmt    *sql.Stmt
	listStudentsStmt  *sql.Stmt
	listTeachersStmt  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                tx,
		tx:                tx,
		createStudentStmt: q.createStudentStmt,
		createTeacherStmt: q.createTeacherStmt,
		deleteStudentStmt: q.deleteStudentStmt,
		deleteTeacherStmt: q.deleteTeacherStmt,
		getStudentStmt:    q.getStudentStmt,
		getTeacherStmt:    q.getTeacherStmt,
		listStudentsStmt:  q.listStudentsStmt,
		listTeachersStmt:  q.listTeachersStmt,
	}
}