-- name: GetUsersAtCreditLimit :many
SELECT *
FROM customers
WHERE total_due = credit_limit;

-- name: GetCustomersWithDue :many
SELECT *
FROM customers
WHERE total_due > 0
ORDER BY total_due DESC;

-- name: GetCustomerDueByName :many
SELECT total_due
FROM customers
WHERE name = ?;


-- name: GetAllMerchantsFeeCollected :many
SELECT
    m.id,
    m.name,
    COALESCE(SUM(t.commission_amount), 0) AS total_fee_collected
FROM merchants m
LEFT JOIN transactions t
    ON m.id = t.merchant_id
GROUP BY m.id, m.name
ORDER BY total_fee_collected DESC;