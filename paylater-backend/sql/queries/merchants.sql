-- name: CreateMerchant :execresult
INSERT INTO merchants (
    name,
    email,
    phone,
    commission_percentage
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
