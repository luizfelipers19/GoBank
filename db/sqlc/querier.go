// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	AddAmountAccountBalanceByID(ctx context.Context, arg AddAmountAccountBalanceByIDParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	DeleteAccountByID(ctx context.Context, id int64) error
	DeleteEntryByID(ctx context.Context, id int64) error
	DeleteTransferByID(ctx context.Context, id int64) error
	GetAccountByID(ctx context.Context, id int64) (Account, error)
	GetAccountByIDForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntries(ctx context.Context) ([]Entry, error)
	GetEntryByID(ctx context.Context, id int64) (Entry, error)
	GetTransferByID(ctx context.Context, id int64) (Transfer, error)
	GetTransfers(ctx context.Context) ([]Transfer, error)
	ListAccounts(ctx context.Context) ([]Account, error)
	UpdateAccountByID(ctx context.Context, arg UpdateAccountByIDParams) (Account, error)
	UpdateEntryByID(ctx context.Context, arg UpdateEntryByIDParams) (Entry, error)
	UpdateTransferByID(ctx context.Context, arg UpdateTransferByIDParams) (Transfer, error)
}

var _ Querier = (*Queries)(nil)
