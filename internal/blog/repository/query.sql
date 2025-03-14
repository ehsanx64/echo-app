-- name: Get :one
SELECT * FROM posts
WHERE id = ? LIMIT 1;

-- name: Fetch :many
SELECT * FROM posts
ORDER BY title;

-- name: Create :execresult
INSERT INTO posts (
  title
) VALUES (
  ?
);
