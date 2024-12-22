-- name: GetUserById :one
SELECT id, username, primary_device, sex, age 
FROM users 
WHERE id = $1 AND is_deleted = FALSE;

-- name: GetUsers :many
SELECT id, username, primary_device, sex, age 
FROM users 
WHERE is_deleted = FALSE
ORDER BY id ASC
LIMIT $1 OFFSET $2;


-- name: UpdateUserPassword :exec
UPDATE users 
SET password = $1 
WHERE id = $2 AND is_deleted = FALSE;


-- name: UpdateUser :exec
UPDATE users 
SET username = $1, age = $2 
WHERE id = $3 AND is_deleted = FALSE;


-- name: SoftDelete :exec
UPDATE users 
SET is_deleted = TRUE 
WHERE id = $1;
