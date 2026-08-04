-- name: CreateMerchant :execresult
INSERT INTO merchants (
    name,
    email,
    password,
    phone
)
VALUES (
    ?,
    ?,
    ?,
    ?
);

-- name: GetMerchantByID :one
SELECT *
FROM merchants
WHERE id = ?;

-- name: GetAllMerchants :many
SELECT *
FROM merchants;

-- name: UpdateMerchantCommission :exec
UPDATE merchants
SET commission_percentage = ?
WHERE id = ?;

-- name: UpdateMerchant :exec
UPDATE merchants
SET name = ?,email = ?
WHERE id = ?;

-- name: DeleteMerchantById :exec
DELETE
FROM merchants
WHERE id = ?;

-- name: GetMerchantByEmail :one
SELECT *
FROM merchants
WHERE email = ?;