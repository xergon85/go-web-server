-- name: GetUser :one
SELECT id, username, password, created_at 
FROM users WHERE id = ? LIMIT 1;

-- name: ListUsers :many
Select id, username, password, created_at 
FROM users;

-- name: CreateUser :one
INSERT INTO users (
  username, password, created_at
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
set username = ?,
password = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = ?;