-- name: CreateUser :one
INSERT INTO users (username, email, hashed_password, role)
VALUES ($1, $2, $3, $4)
RETURNING id, username, email, role, created_at;
