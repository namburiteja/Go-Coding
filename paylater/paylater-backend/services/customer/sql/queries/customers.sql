-- name: CreateCustomer :execresult
INSERT INTO customers (
    name,
    email,
    password,
    payment_due_date
)
VALUES (
    ?,
    ?,
    ?,
    ?
);

-- name: GetCustomerByID :one
SELECT *
FROM customers
WHERE id = ?;

-- name: GetAllCustomers :many
SELECT * 
FROM customers;

-- name: UpdateCustomer :exec
UPDATE customers
SET
    name = ?,
    email = ?
WHERE id = ?;

-- name: UpdateCustomerDue :exec
UPDATE customers
SET total_due = ?
WHERE id = ?;

-- name: UpdateCustomerCreditState :exec
UPDATE customers
SET
    total_due = ?,
    payment_due_date = ?,
    status = ?
WHERE id = ?;

-- name: DeleteCustomerById :exec
DELETE
FROM customers
WHERE id = ?;

-- name: GetCustomerByEmail :one
SELECT *
FROM customers
WHERE email = ?;


-- name: IncreaseCustomerDue :exec
UPDATE customers
SET total_due = total_due + ?
WHERE id = ?;

-- name: DecreaseCustomerDue :exec
UPDATE customers
SET total_due = total_due - ?
WHERE id = ?;

-- name: UpdateCustomerStatus :exec
UPDATE customers
SET status = ?
WHERE id = ?;

-- name: GetCustomerByIDForUpdate :one
SELECT *
FROM customers
WHERE id = ?
LIMIT 1
FOR UPDATE;

-- name: GetUsersAtCreditLimit :many
SELECT *
FROM customers
WHERE total_due >= credit_limit;

-- name: GetCustomersWithDue :many
SELECT *
FROM customers
WHERE total_due > 0
ORDER BY total_due DESC;

-- name: GetCustomerDueByName :one
SELECT *
FROM customers
WHERE name = ?
LIMIT 1;