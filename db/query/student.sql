-- name: GetStudent :one
SELECT * FROM students
WHERE id = ? LIMIT 1;

-- name: ListStudents :many
SELECT * FROM students
WHERE status = 'active'
ORDER BY email;

-- name: CreateStudent :execresult
INSERT INTO students (
  email
) VALUES (
  ?
);

-- name: DeleteStudent :exec
DELETE FROM students
WHERE id = ?;