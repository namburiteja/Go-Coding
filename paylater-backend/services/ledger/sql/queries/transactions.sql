-- name: CreateTransaction :execresult
INSERT INTO transactions (
    customer_id,
    merchant_id,
    transaction_type,
    amount,
    commission_percentage,
    commission_amount,
    transaction_date
)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    NOW()
);

-- name: GetTransactionsByCustomerID :many
SELECT *
FROM transactions
WHERE customer_id = ?
ORDER BY transaction_date DESC;

-- name: GetTransactionsByMerchantID :many
SELECT *
FROM transactions
WHERE merchant_id = ?
ORDER BY transaction_date DESC;

-- name: GetAllTransactions :many
SELECT *
FROM transactions;