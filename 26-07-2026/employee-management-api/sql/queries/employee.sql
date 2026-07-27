-- name: CreateEmployee :execresult
INSERT INTO employees (
    name,
    email,
    age,
    department
)
VALUES (?, ?, ?, ?);

-- name: GetEmployees :many
SELECT *
FROM employees;

-- name: GetEmployee :one
SELECT *
FROM employees
WHERE id = ?;

-- name: UpdateEmployee :exec
UPDATE employees
SET
    name = ?,
    email = ?,
    age = ?,
    department = ?
WHERE id = ?;

-- name: DeleteEmployee :exec
DELETE FROM employees
WHERE id = ?;