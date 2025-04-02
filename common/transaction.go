package common

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type TransactionManager interface {
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type ContextKey string

const (
	TransactionKey ContextKey = "transaction"
)

func GetTxFromContext(ctx context.Context) *sqlx.Tx {
	if tx, ok := ctx.Value(TransactionKey).(*sqlx.Tx); ok {
		return tx
	}
	return nil
}
