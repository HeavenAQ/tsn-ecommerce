-- name: GetLanguageIdByCode :one
SELECT
    pk
FROM
    languages
WHERE
    code = $1;
