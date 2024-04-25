// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: account.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO
    accounts (owner, balance, currency)
VALUES
    ($1, $2, $3) 
RETURNING id, owner, balance, currency, created
`

type CreateAccountParams struct {
	Owner    string         `json:"owner"`
	Balance  pgtype.Numeric `json:"balance"`
	Currency Currency       `json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error) {
	row := q.db.QueryRow(ctx, createAccount, arg.Owner, arg.Balance, arg.Currency)
	var i Accounts
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.Created,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, owner, balance, currency, created FROM accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Accounts, error) {
	row := q.db.QueryRow(ctx, getAccount, id)
	var i Accounts
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.Created,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, owner, balance, currency, created FROM accounts
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListAccountsParams struct {
	Owner  string `json:"owner"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Accounts, error) {
	rows, err := q.db.Query(ctx, listAccounts, arg.Owner, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Accounts
	for rows.Next() {
		var i Accounts
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.Created,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccountBalance = `-- name: UpdateAccountBalance :one
UPDATE accounts
SET balance = $2
WHERE id = $1
RETURNING id, owner, balance, currency, created
`

type UpdateAccountBalanceParams struct {
	ID      int64          `json:"id"`
	Balance pgtype.Numeric `json:"balance"`
}

func (q *Queries) UpdateAccountBalance(ctx context.Context, arg UpdateAccountBalanceParams) (Accounts, error) {
	row := q.db.QueryRow(ctx, updateAccountBalance, arg.ID, arg.Balance)
	var i Accounts
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.Created,
	)
	return i, err
}

const updateAccountOwner = `-- name: UpdateAccountOwner :exec
UPDATE accounts
SET owner = $2
WHERE id = $1
`

type UpdateAccountOwnerParams struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
}

func (q *Queries) UpdateAccountOwner(ctx context.Context, arg UpdateAccountOwnerParams) error {
	_, err := q.db.Exec(ctx, updateAccountOwner, arg.ID, arg.Owner)
	return err
}