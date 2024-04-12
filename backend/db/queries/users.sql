-- name: CreateUser :one
INSERT INTO users (email, phone, PASSWORD, first_name, last_name, language_pk, address, last_login)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING
    *;
