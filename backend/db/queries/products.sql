-- name: CreateProduct :one
INSERT INTO products (price, image, status, quantity)
    VALUES ($1, $2, $3, $4)
RETURNING
    *;

-- name: GetProduct :one
SELECT
    *
FROM
    products
WHERE
    pk = $1;

-- name: ListProducts :many
SELECT
    *
FROM
    products
LIMIT $1 offset $2;

-- name: UpdateProduct :one
UPDATE
    products
SET
    price = $2,
    image = $3,
    status = $4,
    quantity = $5
WHERE
    pk = $1
RETURNING
    *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE pk = $1;
