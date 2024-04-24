// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyRUB Currency = "RUB"
)

func (e *Currency) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Currency(s)
	case string:
		*e = Currency(s)
	default:
		return fmt.Errorf("unsupported scan type for Currency: %T", src)
	}
	return nil
}

type NullCurrency struct {
	Currency Currency `json:"currency"`
	Valid    bool     `json:"valid"` // Valid is true if Currency is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCurrency) Scan(value interface{}) error {
	if value == nil {
		ns.Currency, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Currency.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCurrency) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Currency), nil
}

type Accounts struct {
	ID       int64            `json:"id"`
	Owner    string           `json:"owner"`
	Balance  pgtype.Numeric   `json:"balance"`
	Currency Currency         `json:"currency"`
	Created  pgtype.Timestamp `json:"created"`
}

type Entries struct {
	ID        int64            `json:"id"`
	AccountID int64            `json:"account_id"`
	Amount    pgtype.Numeric   `json:"amount"`
	Created   pgtype.Timestamp `json:"created"`
}

type Transfers struct {
	ID            int64            `json:"id"`
	FromAccountID int64            `json:"from_account_id"`
	ToAccountID   int64            `json:"to_account_id"`
	Amount        pgtype.Numeric   `json:"amount"`
	Created       pgtype.Timestamp `json:"created"`
}
