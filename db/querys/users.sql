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

-- name: UserNameTaken :one
SELECT COUNT(*) > 0 AS is_taken
FROM users
WHERE username = $1 AND is_deleted = FALSE;

-- name: CreateNewUser :exec
INSERT INTO users (username, password, primary_device, sex, age)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: GetUserWithPassword :one
SELECT *
FROM users
WHERE username = $1
  AND password = $2
  AND is_deleted = FALSE;

-- name: GetUserByUserName :one
SELECT *
FROM users
WHERE username = $1 AND is_deleted = FALSE;
