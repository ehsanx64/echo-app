-- name: Get :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: Fetch :many
SELECT * FROM users
ORDER BY name;

-- name: Create :execresult
INSERT INTO users (
  name
) VALUES (
  ?
);
