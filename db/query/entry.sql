-- name: CreateEntry :one
INSERT INTO entries (
  account_id, 
  amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntryByID :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: GetEntries :many
SELECT * FROM entries
ORDER BY id;

-- name: UpdateEntryByID :one
UPDATE entries
  set amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteEntryByID :exec
DELETE FROM entries
WHERE id = $1;