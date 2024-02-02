-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, 
  to_account_id, 
  amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransferByID :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: GetTransfers :many
SELECT * FROM transfers
ORDER BY id;

-- name: UpdateTransferByID :one
UPDATE transfers
  set amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTransferByID :exec
DELETE FROM transfers
WHERE id = $1;