package builder

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/jmoiron/sqlx"
)

type sqlxTransactionManager struct {
	db *sqlx.DB
}

func NewSQLxTransactionManager(db *sqlx.DB) common.TransactionManager {
	return &sqlxTransactionManager{db: db}
}

func (m *sqlxTransactionManager) Begin(ctx context.Context) (context.Context, error) {
	tx, err := m.db.Beginx()
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, common.TransactionKey, tx), nil
}

func (m *sqlxTransactionManager) Commit(ctx context.Context) error {
	tx := common.GetTxFromContext(ctx)
	if tx == nil {
		return common.NewInternalServerError().WithReason("no transaction found in context")
	}
	return tx.Commit()
}

func (m *sqlxTransactionManager) Rollback(ctx context.Context) error {
	tx := common.GetTxFromContext(ctx)
	if tx == nil {
		return nil // no transaction to rollback
	}
	return tx.Rollback()
}
