package db

import (
	"context"
	"database/sql"
	"fmt"
)

// This store entity will provide all methods to perform queries and db operations
type Store struct {
	*Queries
	db *sql.DB
}

// Creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------------------

// executeTransaction function (generic function to be used as a helper to concrete implementations)
func (store *Store) execTransaction(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollBack error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// ------------------------------------------------------------------------------------------------------------------------------------------------

// contains the input that will be sent to the TransferTransaction operation
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// contains the Transfer result that will be the output of TransferTransaction operation
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry`
}

// TransferTransaction
// performs all needed transactions in order to perform a whole transfer (create transfer, entries for leaving and receiving money, and updating both account balances).
// Will Use the execTransactions func as a helper
func (store *Store) TransferTransaction(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTransaction(ctx, func(q *Queries) error {
		var err error

		// First, we create a Transfer record
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// Then, we create a negative Entry record for the account that is sending money
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		// Then, we create a positive Entry record for the account that is receiving money
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		// if the FromAccountID is smaller than the ToAccountID, we remove money from the FromAccount first, and then we add money to the destiny account
		if arg.FromAccountID < arg.ToAccountID {

			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)

		} else { // if the FromAccountID is greater than the ToAccountID, we remove money from the FromAccount after adding money to the destiny account

			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)

		}

		return nil

	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,

) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAmountAccountBalanceByID(ctx, AddAmountAccountBalanceByIDParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}
	account2, err = q.AddAmountAccountBalanceByID(ctx, AddAmountAccountBalanceByIDParams{
		ID:     accountID2,
		Amount: amount2,
	})

	return
}
