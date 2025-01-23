// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package database

import (
	"context"
	"encoding/json"
)

const add = `-- name: Add :one
INSERT INTO messages (content) VALUES ($1) RETURNING id
`

func (q *Queries) Add(ctx context.Context, content json.RawMessage) (int64, error) {
	row := q.db.QueryRow(ctx, add, content)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const find = `-- name: Find :one
SELECT id, content FROM messages WHERE content = ($1::jsonb)
`

func (q *Queries) Find(ctx context.Context, content json.RawMessage) (Message, error) {
	row := q.db.QueryRow(ctx, find, content)
	var i Message
	err := row.Scan(&i.ID, &i.Content)
	return i, err
}

const list = `-- name: List :many
SELECT id, content FROM messages
`

func (q *Queries) List(ctx context.Context) ([]Message, error) {
	rows, err := q.db.Query(ctx, list)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(&i.ID, &i.Content); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
