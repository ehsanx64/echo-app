// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package repository

import (
	"context"
	"database/sql"
)

const create = `-- name: Create :execresult
INSERT INTO posts (
  title
) VALUES (
  ?
)
`

func (q *Queries) Create(ctx context.Context, title string) (sql.Result, error) {
	return q.db.ExecContext(ctx, create, title)
}

const fetch = `-- name: Fetch :many
SELECT id, title FROM posts
ORDER BY title
`

func (q *Queries) Fetch(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, fetch)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(&i.ID, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const get = `-- name: Get :one
SELECT id, title FROM posts
WHERE id = ? LIMIT 1
`

func (q *Queries) Get(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, get, id)
	var i Post
	err := row.Scan(&i.ID, &i.Title)
	return i, err
}
