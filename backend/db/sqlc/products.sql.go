// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: products.sql

package db

import (
	"context"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (price, "imageURLs", status, quantity)
    VALUES ($1, $2, $3, $4)
RETURNING
    pk, id, price, discount, "imageURLs", status, quantity, created_at, updated_at
`

type CreateProductParams struct {
	Price     int32
	ImageURLs []string
	Status    ProductStatus
	Quantity  int32
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.Price,
		arg.ImageURLs,
		arg.Status,
		arg.Quantity,
	)
	var i Product
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.Price,
		&i.Discount,
		&i.ImageURLs,
		&i.Status,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE pk = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, pk int64) error {
	_, err := q.db.Exec(ctx, deleteProduct, pk)
	return err
}

const getProdcutByStatus = `-- name: GetProdcutByStatus :many
SELECT
    pk, id, price, discount, "imageURLs", status, quantity, created_at, updated_at
FROM
    products
WHERE
    status = $1
`

func (q *Queries) GetProdcutByStatus(ctx context.Context, status ProductStatus) ([]Product, error) {
	rows, err := q.db.Query(ctx, getProdcutByStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.Pk,
			&i.ID,
			&i.Price,
			&i.Discount,
			&i.ImageURLs,
			&i.Status,
			&i.Quantity,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProduct = `-- name: GetProduct :one
SELECT
    pk, id, price, discount, "imageURLs", status, quantity, created_at, updated_at
FROM
    products
WHERE
    pk = $1
`

func (q *Queries) GetProduct(ctx context.Context, pk int64) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, pk)
	var i Product
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.Price,
		&i.Discount,
		&i.ImageURLs,
		&i.Status,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProductByPrice = `-- name: GetProductByPrice :many
SELECT
    pk, id, price, discount, "imageURLs", status, quantity, created_at, updated_at
FROM
    products
WHERE
    price >= $1
    AND price <= $2
`

type GetProductByPriceParams struct {
	Price   int32
	Price_2 int32
}

func (q *Queries) GetProductByPrice(ctx context.Context, arg GetProductByPriceParams) ([]Product, error) {
	rows, err := q.db.Query(ctx, getProductByPrice, arg.Price, arg.Price_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.Pk,
			&i.ID,
			&i.Price,
			&i.Discount,
			&i.ImageURLs,
			&i.Status,
			&i.Quantity,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductByQuantity = `-- name: GetProductByQuantity :many
SELECT
    pk, id, price, discount, "imageURLs", status, quantity, created_at, updated_at
FROM
    products
WHERE
    quantity >= $1
    AND quantity <= $2
`

type GetProductByQuantityParams struct {
	Quantity   int32
	Quantity_2 int32
}

func (q *Queries) GetProductByQuantity(ctx context.Context, arg GetProductByQuantityParams) ([]Product, error) {
	rows, err := q.db.Query(ctx, getProductByQuantity, arg.Quantity, arg.Quantity_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.Pk,
			&i.ID,
			&i.Price,
			&i.Discount,
			&i.ImageURLs,
			&i.Status,
			&i.Quantity,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProducts = `-- name: ListProducts :many
SELECT
    pk, id, price, discount, "imageURLs", status, quantity, created_at, updated_at
FROM
    products
LIMIT $1 offset $2
`

type ListProductsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.Pk,
			&i.ID,
			&i.Price,
			&i.Discount,
			&i.ImageURLs,
			&i.Status,
			&i.Quantity,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE
    products
SET
    price = $2,
    "imageURLs" = $3,
    status = $4,
    quantity = $5
WHERE
    pk = $1
RETURNING
    pk, id, price, discount, "imageURLs", status, quantity, created_at, updated_at
`

type UpdateProductParams struct {
	Pk        int64
	Price     int32
	ImageURLs []string
	Status    ProductStatus
	Quantity  int32
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.Pk,
		arg.Price,
		arg.ImageURLs,
		arg.Status,
		arg.Quantity,
	)
	var i Product
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.Price,
		&i.Discount,
		&i.ImageURLs,
		&i.Status,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProductIdxImageURL = `-- name: UpdateProductIdxImageURL :one
UPDATE
    products
SET
    "imageURLs"[$2] = $3
WHERE
    pk = $1
RETURNING
    pk, id, price, discount, "imageURLs", status, quantity, created_at, updated_at
`

type UpdateProductIdxImageURLParams struct {
	Pk          int64
	ImageURLs   []string
	ImageURLs_2 []string
}

func (q *Queries) UpdateProductIdxImageURL(ctx context.Context, arg UpdateProductIdxImageURLParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProductIdxImageURL, arg.Pk, arg.ImageURLs, arg.ImageURLs_2)
	var i Product
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.Price,
		&i.Discount,
		&i.ImageURLs,
		&i.Status,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
