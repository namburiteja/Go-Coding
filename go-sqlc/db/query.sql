-- name: CreateEmployee :execresult
INSERT INTO employee (
    name,
    age,
    salary,
    team
)
VALUES (?, ?, ?, ?);


-- name: GetEmployees :many
SELECT *
FROM employee;


-- name: GetEmployee :one
SELECT *
FROM employee
WHERE id = ?;


-- name: UpdateEmployee :exec
UPDATE employee
SET salary = ?, team = ?
WHERE id = ?;


-- name: DeleteEmployee :exec
DELETE FROM employee
WHERE id = ?;