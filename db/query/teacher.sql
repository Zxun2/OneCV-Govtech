-- name: GetTeacher :one
SELECT * FROM teachers
WHERE id = ? LIMIT 1;

-- name: ListTeachers :many
SELECT * FROM teachers
ORDER BY email;

-- name: CreateTeacher :execresult
INSERT INTO teachers (
  email
) VALUES (
  ?
);

-- name: DeleteTeacher :exec
DELETE FROM teachers
WHERE id = ?;