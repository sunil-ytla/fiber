-- name: GetUserByID :one
SELECT * FROM users where id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username, password) VALUES ($1, $2) RETURNING *;

-- name: GetBookById :one
SELECT * FROM books where id = $1 LIMIT 1;

-- name: AddBook :one
INSERT INTO books (title, status, author, year, userid) 
VALUES ($1, $2, $3, $4, $5) RETURNING *;