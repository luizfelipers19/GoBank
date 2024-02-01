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

// executeTransaction function
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
