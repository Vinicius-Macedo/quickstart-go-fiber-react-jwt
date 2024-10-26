-- name: CreateUser :one
INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUserByEmail :one
SELECT u.*, i.image_url
FROM users u
LEFT JOIN images i ON u.email = i.user_email
WHERE u.email = $1;

-- name: UpdateUserPasswordByEmail :exec
UPDATE users SET password = $2 WHERE email = $1;

-- name: DeleteUserByUsername :exec
DELETE FROM users WHERE username = $1;

-- name: DeleteUserByEmail :exec
DELETE FROM users WHERE email = $1;

-- name: InsertImageOnUserByEmail :one
INSERT INTO images (user_email, image_url) VALUES ($1, $2) RETURNING *;

-- name: CheckUserExistsByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CheckUserExistsByUsername :one
SELECT * FROM users WHERE username = $1;