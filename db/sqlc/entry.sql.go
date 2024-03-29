// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: entry.sql

package db

import (
	"context"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (
  account_id, 
  amount
) VALUES (
  $1, $2
)
RETURNING id, account_id, amount, created_at
`

type CreateEntryParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.AccountID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntryByID = `-- name: DeleteEntryByID :exec
DELETE FROM entries
WHERE id = $1
`

func (q *Queries) DeleteEntryByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEntryByID, id)
	return err
}

const getEntries = `-- name: GetEntries :many
SELECT id, account_id, amount, created_at FROM entries
ORDER BY id
`

func (q *Queries) GetEntries(ctx context.Context) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, getEntries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entry{}
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
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

const getEntryByID = `-- name: GetEntryByID :one
SELECT id, account_id, amount, created_at FROM entries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntryByID(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntryByID, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const updateEntryByID = `-- name: UpdateEntryByID :one
UPDATE entries
  set amount = $2
WHERE id = $1
RETURNING id, account_id, amount, created_at
`

type UpdateEntryByIDParams struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

func (q *Queries) UpdateEntryByID(ctx context.Context, arg UpdateEntryByIDParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, updateEntryByID, arg.ID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
