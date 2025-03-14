-- name: Get :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: Fetch :many
SELECT * FROM users
ORDER BY name;

-- name: Create :execresult
INSERT INTO users (
  name
) VALUES (
  $1
);
