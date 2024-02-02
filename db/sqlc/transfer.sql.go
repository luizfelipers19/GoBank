// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, 
  to_account_id, 
  amount
) VALUES (
  $1, $2, $3
)
RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransferByID = `-- name: DeleteTransferByID :exec
DELETE FROM transfers
WHERE id = $1
`

func (q *Queries) DeleteTransferByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransferByID, id)
	return err
}

const getTransferByID = `-- name: GetTransferByID :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransferByID(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransferByID, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfers = `-- name: GetTransfers :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY id
`

func (q *Queries) GetTransfers(ctx context.Context) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, getTransfers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
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

const updateTransferByID = `-- name: UpdateTransferByID :one
UPDATE transfers
  set amount = $2
WHERE id = $1
RETURNING id, from_account_id, to_account_id, amount, created_at
`

type UpdateTransferByIDParams struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

func (q *Queries) UpdateTransferByID(ctx context.Context, arg UpdateTransferByIDParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, updateTransferByID, arg.ID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
