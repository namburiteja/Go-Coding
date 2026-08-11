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

-- name: GetMerchantFeeTotals :many
SELECT
    merchant_id,
    CAST(COALESCE(SUM(commission_amount), 0) AS CHAR) AS total_fee_collected
FROM transactions
WHERE merchant_id IS NOT NULL
GROUP BY merchant_id;