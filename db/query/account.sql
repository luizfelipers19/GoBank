-- name: CreateAccount :one
INSERT INTO accounts (
  owner, 
  balance, 
  currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccountByID :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountByIDForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;


-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id;

-- name: UpdateAccountByID :one
UPDATE accounts
  set balance = $2
WHERE id = $1
RETURNING *;

-- name: AddAmountAccountBalanceByID :one
UPDATE accounts
  set balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccountByID :exec
DELETE FROM accounts
WHERE id = $1;