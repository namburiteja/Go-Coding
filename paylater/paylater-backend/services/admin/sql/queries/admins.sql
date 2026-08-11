-- name: CreateAdmin :execresult
INSERT INTO admins (
    name,
    email,
    password
)
VALUES (
    ?,
    ?,
    ?
);

-- name: GetAdminByID :one
SELECT *
FROM admins
WHERE id = ?;

-- name: GetAdminByEmail :one
SELECT *
FROM admins
WHERE email = ?;

-- name: GetAllAdmins :many
SELECT *
FROM admins;

-- name: UpdateAdmin :exec
UPDATE admins
SET
    name = ?,
    email = ?
WHERE id = ?;

-- name: DeleteAdminByID :exec
DELETE
FROM admins
WHERE id = ?;
