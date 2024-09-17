-- name: CreateUser :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING id, name, email;

-- name: GetUserByID :one
SELECT id, name, email
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, email
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;
