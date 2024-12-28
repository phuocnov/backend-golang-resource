// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accounts.sql

package tutorial

import (
	"context"
)

const createAccount = `-- name: createAccount :one
INSERT INTO accounts (
  owner,
  balance,
  currency
) VALUES ( $1, $2, $3 )

RETURNING id, owner, balance, currency, created_at
`

type createAccountParams struct {
	Owner    string
	Balance  int64
	Currency string
}

func (q *Queries) createAccount(ctx context.Context, arg createAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccount, arg.Owner, arg.Balance, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const getAccount = `-- name: getAccount :one
SELECT id, owner, balance, currency, created_at 
FROM accounts
WHERE id = $1 
LIMIT 1
`

func (q *Queries) getAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRow(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}