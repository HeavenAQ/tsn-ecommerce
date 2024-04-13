-- name: CreateProductTranslation :one
INSERT INTO product_translations (product_pk, language_pk, name, description, category)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    *;

-- name: GetProductTranslation :one
SELECT
    *
FROM
    product_translations
WHERE
    product_pk = $1
    AND language_pk = $2
LIMIT 1;

-- name: UpdateProductTranslation :one
UPDATE
    product_translations
SET
    name = $3,
    description = $4,
    category = $5
WHERE
    product_pk = $1
    AND language_pk = $2
RETURNING
    *;

-- name: DeleteProductTranslation :exec
DELETE FROM product_translations
WHERE product_pk = $1
    AND language_pk = $2;
